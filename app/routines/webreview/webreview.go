package webreview

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/aptos"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	ws "github.com/NetSepio/gateway/util/pkg/webscrape"
	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
)

func UploadToIpfs(osFile io.Reader, fileName string) (*NFTStorageRes, error) {
	// Create a buffer to store the request body
	var buf bytes.Buffer

	// Create a new multipart writer with the buffer
	w := multipart.NewWriter(&buf)

	// Create a new form field
	fw, err := w.CreateFormFile("file", fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to create file %s: %w", fileName, err)
	}

	// Copy the contents of the file to the form field
	if _, err := io.Copy(fw, osFile); err != nil {
		return nil, fmt.Errorf("failed to copy contents to %s: %w", fileName, err)
	}

	// Close the multipart writer to finalize the request
	w.Close()

	// Send the request
	req, err := http.NewRequest("POST", "https://api.nft.storage/upload", &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request for nft.storage: %w", err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", envconfig.EnvVars.NFT_STORAGE_KEY))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request for nft.storage: %w", err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body of request for nft.storage: %w", err)
	}
	nftRes, err := UnmarshalNFTStorageRes(bodyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body of request for nft.storage: %w", err)
	}
	defer resp.Body.Close()
	return &nftRes, nil
}

func Publish(metadatahash string, siteUrl string) {
	db := dbconfig.GetDb()
	uid_str := uuid.NewString()
	os.Mkdir("storage", os.ModePerm)

	dirName := path.Join("storage", uid_str)
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		logwrapper.Errorf("failed to create folder %v, error: %v", dirName, err.Error())
	}

	websiteURL := siteUrl
	domain, err := url.Parse(websiteURL)
	if err != nil {
		logwrapper.Errorf("failed to parse url, error: %v", err.Error())
		return
	}
	htmlFile := domain.Host

	err = ws.CheckDomain(dirName, websiteURL)
	if err != nil {
		logwrapper.Errorf("failed to checkDomain for websiteURL : %v, error: %v", websiteURL, err.Error())
		return
	}
	filePath := path.Join(dirName, htmlFile)
	indexFile, err := os.Open(filePath)
	if err != nil {
		logwrapper.Errorf("failed to add file: %v to ipfs for websiteURL : %v, error: %v", filePath, websiteURL, err.Error())
		return
	}
	indexFileUploadRes, err := UploadToIpfs(indexFile, "index.html")
	if err != nil {
		logwrapper.Errorf("failed to add file: %v to ipfs for websiteURL : %v, error: %v", filePath, websiteURL, err.Error())
		return
	}

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())

	// capture screenshot of an element
	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, ws.FullScreenshot(websiteURL, 90, &buf)); err != nil {
		logwrapper.Errorf("failed to run chromedp, error: %v", err.Error())
		return
	}
	fileName := fmt.Sprintf("%v/fullScreenshot.png", dirName)
	if err := os.WriteFile(fileName, buf, 0644); err != nil {
		logwrapper.Errorf("failed to write file: %v, error:%v", fileName, err)
		return
	}

	screenShotFile, err := os.Open(fileName)
	if err != nil {
		logwrapper.Errorf("failed to add file: %v to ipfs for websiteURL : %v, error: %v", filePath, websiteURL, err.Error())
		return
	}
	screenShotFileUploadRes, err := UploadToIpfs(screenShotFile, "screenshot.png")
	if err != nil {
		logwrapper.Errorf("failed to add file: %v to ipfs for websiteURL : %v, error: %v", filePath, websiteURL, err.Error())
		return
	}

	metaData := MetaData{
		WebsiteScreenShot: screenShotFileUploadRes.Value.Cid,
		IndexFile:         indexFileUploadRes.Value.Cid,
	}
	data, err := json.Marshal(&metaData)
	if err != nil {
		logwrapper.Errorf("failed to marshal JSON for uid_str %v : %s", uid_str, err)
		return
	}

	metaDataRes, err := UploadToIpfs(bytes.NewReader(data), "metadata.json")
	if err != nil {
		logwrapper.Errorf("failed to add file to ipfs for fullScreenShot : %v", err.Error())
		return
	}

	//TODO send update tx
	println(metaDataRes.Value.Cid)

	siteUrlWithoutHttps := strings.TrimPrefix(siteUrl, "https://")
	siteUrlWithoutHttps = strings.TrimPrefix(siteUrlWithoutHttps, "http://")
	_, err = aptos.UploadArchive(siteUrlWithoutHttps, metaDataRes.Value.Cid)
	if err != nil {
		logwrapper.Errorf("failed to upload metadatahash, error: %v", err.Error())
		return
	}

	err = db.Model(&models.Review{}).Where("meta_data_uri = ?", metadatahash).Update("site_ipfs_hash", metaDataRes.Value.Cid).Error
	if err != nil {
		logwrapper.Warnf("failed to update site ipfs hash, error: %v", err.Error())
	}
	err = os.RemoveAll(dirName)
	if err != nil {
		logwrapper.Errorf("failed to remove dir %v, error:%v", dirName, err.Error())
		return
	}
	cancel()

}
