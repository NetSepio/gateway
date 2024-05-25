package summary

import (
	"context"
	"net/http"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

type RequestData struct {
	Terms   string `json:"terms" binding:"required"`
	Privacy string `json:"privacy" binding:"required"`
}

type SummaryResponse struct {
	TermsSummary   string `json:"terms_summary"`
	PrivacySummary string `json:"privacy_summary"`
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

	termsSummary, err := generateSummary(requestData.Terms)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	privacySummary, err := generateSummary(requestData.Privacy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, SummaryResponse{
		TermsSummary:   termsSummary,
		PrivacySummary: privacySummary,
	})
}

func generateSummary(content string) (string, error) {
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

	return resp.Choices[0].Text, nil
}
