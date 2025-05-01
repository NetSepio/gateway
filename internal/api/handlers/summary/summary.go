// package summary

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"net/http"

// 	"github.com/NetSepio/gateway/config/envconfig"
// 	"github.com/gin-gonic/gin"
// 	openai "github.com/sashabaranov/go-openai"
// )

// type RequestData struct {
// 	Terms   string `json:"terms" binding:"required"`
// 	Privacy string `json:"privacy" binding:"required"`
// }

// func ApplyRoutes(r *gin.RouterGroup) {
// 	g := r.Group("/summary")
// 	{
// 		g.POST("", summary)
// 	}
// }

// func summary(c *gin.Context) {
// 	var requestData RequestData
// 	if err := c.ShouldBindJSON(&requestData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	writer := c.Writer
// 	writer.WriteHeader(http.StatusOK)

// 	_, err := generateAndStreamSummary(requestData.Terms, writer, "Terms Summary")
// 	if err != nil {
// 		fmt.Fprintf(writer, "\nerror: %v\n", err)
// 		return
// 	}

// 	_, err = generateAndStreamSummary(requestData.Privacy, writer, "Privacy Summary")
// 	if err != nil {
// 		fmt.Fprintf(writer, "\nerror: %v\n", err)
// 		return
// 	}
// }

// func generateAndStreamSummary(content string, writer io.Writer, title string) (string, error) {
// 	open_ai_key := envconfig.EnvVars.OPENAI_API_KEY

// 	client := openai.NewClient(open_ai_key)

// 	req := openai.CompletionRequest{
// 		Model:     "gpt-3.5-turbo-instruct",
// 		Prompt:    "Summarize the following text in key points:\n\n" + content,
// 		MaxTokens: 150,
// 	}

// 	resp, err := client.CreateCompletion(context.Background(), req)
// 	if err != nil {
// 		return "", err
// 	}

// 	if len(resp.Choices) == 0 {
// 		return "", nil
// 	}

// 	summary := resp.Choices[0].Text

// 	_, err = writer.Write([]byte(fmt.Sprintf("\n--- %s ---\n", title)))
// 	if err != nil {
// 		return "", err
// 	}

// 	words := splitIntoWords(summary)
// 	for _, word := range words {
// 		_, err := writer.Write([]byte(word + " "))
// 		if err != nil {
// 			return "", err
// 		}
// 		writer.(http.Flusher).Flush()
// 	}

// 	return summary, nil
// }

// func splitIntoWords(text string) []string {
// 	words := make([]string, 0)
// 	currentWord := ""

// 	for _, char := range text {
// 		if char == ' ' || char == '\n' || char == '\t' {
// 			if currentWord != "" {
// 				words = append(words, currentWord)
// 				currentWord = ""
// 			}
// 			if char != ' ' {
// 				words = append(words, string(char))
// 			}
// 		} else {
// 			currentWord += string(char)
// 		}
// 	}

// 	if currentWord != "" {
// 		words = append(words, currentWord)
// 	}

// 	return words
// }

// package main

// import (
// 	"context"
// 	"fmt"
// 	"strings"
// 	"time"

// 	"github.com/chromedp/cdproto/cdp"
// 	"github.com/chromedp/chromedp"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// Initialize Gin router
// 	router := gin.Default()

// 	// Define a GET endpoint for the API
// 	router.GET("/summary", func(c *gin.Context) {
// 		// Get the website URL from the query parameter
// 		url := c.Query("url")
// 		if url == "" {
// 			c.JSON(400, gin.H{"error": "Website URL is required"})
// 			return
// 		}

// 		// Extract links from the provided website URL
// 		links, err := extractLinksFromURL(url)
// 		if err != nil {
// 			c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to extract links: %v", err)})
// 			return
// 		}

// 		// Check each link for terms of use and privacy policy pages
// 		var termsLinks []string
// 		var privacyLinks []string
// 		for _, link := range links {
// 			// Check if the link contains any of the keywords related to terms of use
// 			if containsTermsKeywords(link) {
// 				termsLinks = append(termsLinks, link)
// 			}
// 			// Check if the link contains any of the keywords related to privacy policy
// 			if containsPrivacyKeywords(link) {
// 				privacyLinks = append(privacyLinks, link)
// 			}
// 		}

// 		// Return the extracted links
// 		c.JSON(200, gin.H{
// 			"terms_of_use":   termsLinks,
// 			"privacy_policy": privacyLinks,
// 		})
// 	})

// 	// Run the server
// 	router.Run(":8080")
// }

// func extractLinksFromURL(url string) ([]string, error) {
// 	// Run a headless Chrome browser
// 	opts := append(
// 		chromedp.DefaultExecAllocatorOptions[3:],
// 		chromedp.NoFirstRun,
// 		chromedp.NoDefaultBrowserCheck,
// 	)

// 	parentCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
// 	defer cancel()
// 	ctx, cancel := chromedp.NewContext(parentCtx)
// 	defer cancel()
// 	var links []string

// 	// Navigate to the provided website URL
// 	if err := chromedp.Run(ctx, navigateToWebsite(url)); err != nil {
// 		return nil, fmt.Errorf("failed to navigate to website: %v", err)
// 	}

// 	// Wait for a brief period before extracting links
// 	if err := chromedp.Run(ctx, chromedp.Sleep(9*time.Second)); err != nil {
// 		return nil, fmt.Errorf("failed to wait: %v", err)
// 	}

// 	// Extract all anchor links from the page
// 	if err := chromedp.Run(ctx, extractLinks(&links)); err != nil {
// 		return nil, fmt.Errorf("failed to extract links: %v", err)
// 	}

// 	return links, nil
// }

// func navigateToWebsite(url string) chromedp.Action {
// 	return chromedp.Navigate(url)
// }

// func extractLinks(links *[]string) chromedp.Action {
// 	return chromedp.ActionFunc(func(ctx context.Context) error {
// 		// Get all anchor elements on the page
// 		var nodes []*cdp.Node
// 		if err := chromedp.Run(ctx, chromedp.Nodes("a", &nodes)); err != nil {
// 			return err
// 		}

// 		// Extract href attribute from each anchor element
// 		for _, node := range nodes {
// 			href := node.AttributeValue("href")
// 			if href != "" {
// 				*links = append(*links, href)
// 			}
// 		}

// 		return nil
// 	})
// }

// func containsTermsKeywords(link string) bool {
// 	// Common keywords related to terms of use
// 	termsKeywords := []string{"terms", "tos", "terms-of-service", "eula", "end-user-license-agreement", "termsofuse"}
// 	link = strings.ToLower(link)
// 	for _, keyword := range termsKeywords {
// 		if strings.Contains(link, keyword) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func containsPrivacyKeywords(link string) bool {
// 	// Common keywords related to privacy policy
// 	privacyKeywords := []string{"privacy", "privacy-policy", "policy", "policies"}
// 	link = strings.ToLower(link)
// 	for _, keyword := range privacyKeywords {
// 		if strings.Contains(link, keyword) {
// 			return true
// 		}
// 	}
// 	return false
// }

package summary

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/summary")
	{
		g.GET("", summary)
	}
}

func summary(c *gin.Context) {

	url := c.Query("url")
	if url == "" {
		c.JSON(400, gin.H{"error": "Website URL is required"})
		return
	}

	links, err := extractLinksFromURL(url)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to extract links: %v", err)})
		return
	}

	termsSummary, privacySummary := summarizeLinksContent(links)

	c.JSON(200, gin.H{
		"termsSummary":   termsSummary,
		"privacySummary": privacySummary,
	})

}
