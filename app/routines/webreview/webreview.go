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

	"github.com/NetSepio/gateway/config/envconfig"
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
		return nil, err
	}

	// Copy the contents of the file to the form field
	if _, err := io.Copy(fw, osFile); err != nil {
		return nil, err
	}

	// Close the multipart writer to finalize the request
	w.Close()

	// Send the request
	req, err := http.NewRequest("POST", "https://api.nft.storage/upload", &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", envconfig.EnvVars.NFT_STORAGE_KEY))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	nftRes, err := UnmarshalNFTStorageRes(bodyBytes)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &nftRes, nil
}

func Publish(siteUrl string) {
	uid_str := uuid.NewString()
	os.Mkdir("storage", os.ModePerm)

	dirName := path.Join("storage", uid_str)
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		logwrapper.Warnf("failed to create folder %v, error: %v", dirName, err.Error())
	}

	websiteURL := siteUrl
	domain, err := url.Parse(websiteURL)
	if err != nil {
		logwrapper.Warnf("failed to parse url, error: %v", err.Error())
	}
	htmlFile := domain.Host

	err = ws.CheckDomain(dirName, websiteURL)
	if err != nil {
		logwrapper.Warnf("failed to checkDomain for websiteURL : %v, error: %v", websiteURL, err.Error())
	}
	filePath := path.Join(dirName, htmlFile)
	indexFile, err := os.Open(filePath)
	if err != nil {
		logwrapper.Warnf("failed to add file: %v to ipfs for websiteURL : %v, error: %v", filePath, websiteURL, err.Error())
	}
	indexFileUploadRes, err := UploadToIpfs(indexFile, "index.html")
	if err != nil {
		logwrapper.Warnf("failed to add file: %v to ipfs for websiteURL : %v, error: %v", filePath, websiteURL, err.Error())
	}

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())

	// capture screenshot of an element
	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, ws.FullScreenshot(websiteURL, 90, &buf)); err != nil {
		logwrapper.Warnf("failed to run chromedp, error: %v", err.Error())
	}
	fileName := fmt.Sprintf("%v/fullScreenshot.png", dirName)
	if err := os.WriteFile(fileName, buf, 0644); err != nil {
		logwrapper.Warnf("failed to write file: %v, error:%v", fileName, err)
	}

	screenShotFile, err := os.Open(fileName)
	if err != nil {
		logwrapper.Warnf("failed to add file: %v to ipfs for websiteURL : %v, error: %v", filePath, websiteURL, err.Error())
	}
	screenShotFileUploadRes, err := UploadToIpfs(screenShotFile, "screenshot.png")
	if err != nil {
		logwrapper.Warnf("failed to add file: %v to ipfs for websiteURL : %v, error: %v", filePath, websiteURL, err.Error())
	}

	metaData := MetaData{
		WebsiteScreenShot: screenShotFileUploadRes.Value.Cid,
		IndexFile:         indexFileUploadRes.Value.Cid,
	}
	data, err := json.Marshal(&metaData)
	if err != nil {
		logwrapper.Warnf("failed to marshal JSON for uid_str %v : %s", uid_str, err)
	}

	metaDataRes, err := UploadToIpfs(bytes.NewReader(data), "metadata.json")
	if err != nil {
		logwrapper.Warnf("failed to add file to ipfs for fullScreenShot : %v", err.Error())
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
	err = os.RemoveAll(dirName)
	if err != nil {
		logwrapper.Warnf("failed to remove dir %v, error:%v", dirName, err.Error())
		return
	}
	cancel()

}
