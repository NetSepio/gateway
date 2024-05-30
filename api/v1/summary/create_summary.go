package summary

import (
	"context"
	"fmt"
	"strings"

	"github.com/NetSepio/gateway/config/envconfig"
	openai "github.com/sashabaranov/go-openai"
)

func summarizeContent(contents []string) string {
	if len(contents) == 0 {
		return ""
	}
	var builder strings.Builder
	for _, content := range contents {
		builder.WriteString(content)
		builder.WriteString("\n")
	}

	prompt := builder.String()
	if len(prompt) > 128000 {
		prompt = prompt[:127999]
	}

	open_ai_key := envconfig.EnvVars.OPENAI_API_KEY

	client := openai.NewClient(open_ai_key)

	req := openai.CompletionRequest{
		Model:     "gpt-4-turbo",
		Prompt:    "Summarize the following in key points under 150 words:\n\n" + prompt,
		MaxTokens: 150,
	}

	summary, err := client.CreateCompletion(context.Background(), req)
	if err != nil {
		fmt.Println("error:", err)
		return "Failed to summarize content"
	}

	return summary.Choices[0].Text
}
