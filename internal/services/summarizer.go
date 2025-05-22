package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"internal/models"
)

type ChatGPTRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type ChatGPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type Summarizer struct {
	apiKey string
}

func NewSummarizer(apiKey string) *Summarizer {
	return &Summarizer{apiKey: apiKey}
}

func (s *Summarizer) Summarize(article *models.Article) (*models.Summary, error) {
	prompt := fmt.Sprintf("Please summarize the following article in 3-5 sentences:\n\nTitle: %s\n\nContent: %s", article.Title, article.Content)

	requestBody := ChatGPTRequest{
		Model: "gpt-3.5-turbo",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var chatResponse ChatGPTResponse
	if err := json.Unmarshal(body, &chatResponse); err != nil {
		return nil, err
	}

	if len(chatResponse.Choices) == 0 {
		return nil, fmt.Errorf("no response from ChatGPT")
	}

	summary := &models.Summary{
		ArticleID: article.ID,
		Text:      chatResponse.Choices[0].Message.Content,
	}

	return summary, nil
}