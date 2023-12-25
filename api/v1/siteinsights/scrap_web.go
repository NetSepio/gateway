package siteinsights

import (
	"context"
	"fmt"

	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/chromedp/chromedp"
)

func ScrapeWebsiteContent(siteURL string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var websiteContent string
	if err := chromedp.Run(ctx, chromedp.Navigate(siteURL), chromedp.Text(`body`, &websiteContent, chromedp.ByQuery)); err != nil {
		logwrapper.Errorf("failed to scrape website content, error: %v", err)
		return "", fmt.Errorf("failed to scrape website content: %w", err)
	}

	return websiteContent, nil
}
