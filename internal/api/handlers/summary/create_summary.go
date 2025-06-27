package summary

import (
	"context"
	"fmt"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	"github.com/NetSepio/gateway/utils/load"
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
	if len(prompt) > 10000 {
		prompt = prompt[:9900]
	}

	open_ai_key := load.Cfg.OPENAI_API_KEY

	client := openai.NewClient(open_ai_key)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Summarize the following in key points under 150 words:\n\n" + prompt,
			},
		}}

	summary, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println("error:", err)
		return "Failed to summarize content"
	}

	return summary.Choices[0].Message.Content
}
