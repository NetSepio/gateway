package summary

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

type RequestData struct {
	Terms   string `json:"terms" binding:"required"`
	Privacy string `json:"privacy" binding:"required"`
}

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/summary")
	{
		g.POST("", summary)
	}
}

func summary(c *gin.Context) {
	var requestData RequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	writer := c.Writer
	writer.WriteHeader(http.StatusOK)

	_, err := generateAndStreamSummary(requestData.Terms, writer, "Terms Summary")
	if err != nil {
		fmt.Fprintf(writer, "\nerror: %v\n", err)
		return
	}

	_, err = generateAndStreamSummary(requestData.Privacy, writer, "Privacy Summary")
	if err != nil {
		fmt.Fprintf(writer, "\nerror: %v\n", err)
		return
	}
}

func generateAndStreamSummary(content string, writer io.Writer, title string) (string, error) {
	open_ai_key := envconfig.EnvVars.OPENAI_API_KEY

	client := openai.NewClient(open_ai_key)

	req := openai.CompletionRequest{
		Model:     "gpt-3.5-turbo-instruct",
		Prompt:    "Summarize the following text in key points:\n\n" + content,
		MaxTokens: 150,
	}

	resp, err := client.CreateCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", nil
	}

	summary := resp.Choices[0].Text

	_, err = writer.Write([]byte(fmt.Sprintf("\n--- %s ---\n", title)))
	if err != nil {
		return "", err
	}

	words := splitIntoWords(summary)
	for _, word := range words {
		_, err := writer.Write([]byte(word + " "))
		if err != nil {
			return "", err
		}
		writer.(http.Flusher).Flush()
	}

	return summary, nil
}

func splitIntoWords(text string) []string {
	words := make([]string, 0)
	currentWord := ""

	for _, char := range text {
		if char == ' ' || char == '\n' || char == '\t' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
			if char != ' ' {
				words = append(words, string(char))
			}
		} else {
			currentWord += string(char)
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	return words
}
