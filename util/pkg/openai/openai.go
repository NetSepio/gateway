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
	Model     string        `json:"model"`
	Messages  []ChatMessage `json:"messages"`
	MaxTokens int           `json:"max_tokens"`
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
			{Role: "system",
				Content: `Your task is to provide an sumarry for site in max 340 characters including things like space including things like space. Tell what it is about. And should user trust this site as a genuine one. Don't include text which decribes uncertainty in review like 
			"Trustworthiness is unclear without user reviews or external validation." this has uncertain and you should not make user feel uncertain about it. Note, making it max 340 characters including things like space is very very important, please make this priority.`},
			{Role: "user", Content: fmt.Sprintf("Tell summary only in max 340 characters including things like space. Here is the content from the website '%s':\n%s\n\n, IMPORTANT - please make 340 characters including things like space max priority.", siteURL, websiteContent)},
		},
		MaxTokens: 150,
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

func IsReviewSpam(review string) (bool, error) {
	requestBody, err := json.Marshal(ChatRequest{
		Model: "gpt-4-1106-preview",
		Messages: []ChatMessage{
			{Role: "system", Content: `Your task is to tell if the review is spam or real one, like this 'Best website' looks spam because it doesn't have any detail.
			For more examples of spam,
			- Fadshfhdasfk
			- 1
			- 11112312134214543252q4332
			- Aaaaa
			- Good
			- Bad
			- It’s good
			- .
			- Awesome
			- goood
			- nice
			- Bullish
			- Solid
			- Scam
			- Airdrop
			- Follow for airdrop
			- N/A
			- This is good app
			- Easy to use
			- Wezsdxcfgbhjnkibhytcrxeswe4srtyuinnkjh gctxreza4rdtyuhjni
			- Good good good good good good good good good good good good good good good good good 
			- Gm	
			
			Just answer with YES and NO, nothing else
			`},
			{Role: "user", Content: fmt.Sprintf("Review %s", review)},
		},
		MaxTokens: 1,
	})
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+envconfig.EnvVars.OPENAI_API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		return false, err
	}
	if len(chatResponse.Choices) > 0 && len(chatResponse.Choices[0].Message.Content) > 0 {
		return chatResponse.Choices[0].Message.Content == "YES", nil
	}

	return false, fmt.Errorf("no insight generated: %s", body)
}
