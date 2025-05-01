package summary

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func extractLinksFromURL(url string) ([]string, error) {
	opts := chromedp.DefaultExecAllocatorOptions[:]

	parentCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(parentCtx)
	defer cancel()

	var links []string

	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("a", chromedp.ByQuery),
		extractLinks(url, &links),
	}); err != nil {
		return nil, fmt.Errorf("failed to extract links: %v", err)
	}

	return links, nil
}

func navigateToWebsite(url string) chromedp.Action {
	return chromedp.Navigate(url)
}

func extractLinks(baseURL string, links *[]string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		var nodes []*cdp.Node
		if err := chromedp.Run(ctx, chromedp.Nodes("a", &nodes)); err != nil {
			return err
		}

		base, err := url.Parse(baseURL)
		if err != nil {
			return fmt.Errorf("invalid base URL: %v", err)
		}

		for _, node := range nodes {
			href := node.AttributeValue("href")
			if href != "" {
				link, err := url.Parse(href)
				if err != nil {
					continue
				}
				if !link.IsAbs() {
					link = base.ResolveReference(link)
				}
				*links = append(*links, link.String())
			}
		}
		// fmt.Println("Extracted links: ", *links)
		return nil
	})
}

func extractContentFromLink(link string) (string, error) {
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],

		// chromedp.DefaultExecAllocatorOptions[3:],
		// chromedp.NoFirstRun,
		// chromedp.NoDefaultBrowserCheck,
	)

	parentCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(parentCtx)
	defer cancel()

	var content string

	if err := chromedp.Run(ctx, navigateToWebsite(link)); err != nil {
		return "", fmt.Errorf("failed to navigate to link: %v", err)
	}

	if err := chromedp.Run(ctx, extractText(&content)); err != nil {
		return "", fmt.Errorf("failed to extract text content: %v", err)
	}
	// fmt.Println(content)

	return content, nil
}

func extractText(content *string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		var node string
		if err := chromedp.Run(ctx, chromedp.Text("body", &node, chromedp.NodeVisible, chromedp.ByQuery)); err != nil {
			return err
		}

		*content = node
		return nil
	})
}

func containsTermsKeywords(link string) bool {
	termsKeywords := []string{"terms", "tos", "terms-of-service", "eula", "end-user-license-agreement", "termsofuse"}
	link = strings.ToLower(link)
	for _, keyword := range termsKeywords {
		if strings.Contains(link, keyword) {
			// fmt.Println("imp:  ", link)
			return true
		}
	}
	return false
}

func containsPrivacyKeywords(link string) bool {
	privacyKeywords := []string{"privacy", "privacy-policy", "policy", "policies"}
	link = strings.ToLower(link)
	for _, keyword := range privacyKeywords {
		if strings.Contains(link, keyword) {
			// fmt.Println(link)
			return true
		}
	}
	return false
}

func summarizeLinksContent(links []string) (termsSummary, privacySummary string) {
	var termsContents, privacyContents []string

	for _, link := range links {
		if containsTermsKeywords(link) {
			content, err := extractContentFromLink(link)
			if err == nil {
				termsContents = append(termsContents, content)
				// fmt.Println(termsContents)
			} else {
				fmt.Printf("Error extracting content from terms link %s: %v\n", link, err)
			}
		}
		if containsPrivacyKeywords(link) {
			content, err := extractContentFromLink(link)
			if err == nil {
				privacyContents = append(privacyContents, content)
				// fmt.Println(privacyContents)
			} else {
				fmt.Printf("Error extracting content from privacy link %s: %v\n", link, err)
			}
		}
	}

	termsSummary = summarizeContent(termsContents)
	privacySummary = summarizeContent(privacyContents)

	return termsSummary, privacySummary
}
