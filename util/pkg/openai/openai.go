package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NetSepio/gateway/config/envconfig"
)

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

func GenerateInsight(siteURL, websiteContent string) (string, error) {
	requestBody, err := json.Marshal(ChatRequest{
		Model: "gpt-4-1106-preview",
		Messages: []ChatMessage{
			{Role: "system", Content: "Your task is to provide an summarry for site. Tell what it is about. And should user trust this site as a genuine one"},
			{Role: "user", Content: fmt.Sprintf("Here is the content from the website '%s':\n%s\nWhat is an interesting insight about this website?", siteURL, websiteContent)},
		},
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+envconfig.EnvVars.OPENAI_API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		return "", err
	}
	if len(chatResponse.Choices) > 0 && len(chatResponse.Choices[0].Message.Content) > 0 {
		return chatResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no insight generated: %s", body)
}
