package webreview

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/url"
	"os"
	"path"

	"github.com/TheLazarusNetwork/netsepio-engine/config/netsepio"
	"github.com/TheLazarusNetwork/netsepio-engine/config/smartcontract"
	"github.com/TheLazarusNetwork/netsepio-engine/config/smartcontract/rawtrasaction"
	"github.com/TheLazarusNetwork/netsepio-engine/generated/smartcontract/gennetsepio"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/logwrapper"
	ws "github.com/TheLazarusNetwork/netsepio-engine/util/pkg/webscrape"
	"github.com/chromedp/chromedp"
	"github.com/ethereum/go-ethereum/common"
)

func Init() {

	os.Mkdir("storage", os.ModePerm)
	client, err := smartcontract.GetClient()
	if err != nil {
		logwrapper.Fatalf("failed to get eth client, error: %v", err.Error())
	}
	netsepioInstance, err := netsepio.GetInstance(client)
	if err != nil {
		logwrapper.Fatalf("failed to get Contract instance, error: %v", err.Error())
	}

	reviewCreatedChannel := make(chan *gennetsepio.GennetsepioReviewCreated)
	_, err = netsepioInstance.WatchReviewCreated(nil, reviewCreatedChannel, []common.Address{}, []*big.Int{})
	for e := range reviewCreatedChannel {
		dirName := path.Join("storage", e.TokenId.String())
		err := os.Mkdir(dirName, os.ModePerm)
		if err != nil {
			logwrapper.Warnf("failed to create folder %v, error: %v", dirName, err.Error())
			continue
		}

		websiteURL := e.SiteURL
		domain, err := url.Parse(websiteURL)
		if err != nil {
			logwrapper.Warnf("failed to parse url, error: %v", err.Error())
			continue
		}
		htmlFile := domain.Host

		err = ws.CheckDomain(dirName, websiteURL)
		if err != nil {
			logwrapper.Warnf("failed to checkDomain for websiteURL : %v, error: %v", websiteURL, err.Error())
			continue
		}
		filePath := path.Join(dirName, htmlFile)
		hash, err := ws.AddFileToIpfs(filePath)
		if err != nil {
			logwrapper.Warnf("failed to add file: %v to ipfs for websiteURL : %v, error: %v", filePath, websiteURL, err.Error())
			continue
		}

		err = ws.GetObjectFromIpfs(hash, path.Join(dirName, "output.txt"))
		if err != nil {
			logwrapper.Warnf("failed to run GetObjectFromIpfs for hash: %v, error: %v", hash, err.Error())
			continue
		}
		// create context
		ctx, cancel := chromedp.NewContext(context.Background())
		defer cancel()

		// capture screenshot of an element
		var buf []byte
		// capture entire browser viewport, returning png with quality=90
		if err := chromedp.Run(ctx, ws.FullScreenshot(websiteURL, 90, &buf)); err != nil {
			logwrapper.Warnf("failed to run chromedp, error: %v", err.Error())
			continue
		}
		fileName := fmt.Sprintf("%v/fullScreenshot.png", dirName)
		if err := ioutil.WriteFile(fileName, buf, 0644); err != nil {
			logwrapper.Warnf("failed to write file: %v, error:%v", fileName, err)
			continue
		}

		hash, err = ws.AddFileToIpfs(fileName)
		if err != nil {
			logwrapper.Warnf("failed to add file to ipfs for fullScreenShot : %v", err.Error())
			continue
		}
		// netsepioInstance.UpdateReview(nil, e.TokenId, hash)
		_, err = rawtrasaction.SendRawTrasac(gennetsepio.GennetsepioABI, "updateReview", e.TokenId, hash)
		if err != nil {
			logwrapper.Warnf("failed to updateReview for tokenId %v : %v", e.TokenId, err.Error())
			continue
		}

		err = os.RemoveAll(dirName)
		if err != nil {
			logwrapper.Warnf("failed to remove dir %v, error:%v", dirName, err.Error())
		}
	}

}