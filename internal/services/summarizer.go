package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/Zaki2409/go-article-video-generator/internal/models"
)

type RapidAPISummarizeRequest struct {
	Lang string `json:"lang"`
	Text string `json:"text"`
}

type RapidAPISummarizeResponse struct {
	Summary string `json:"summary"`
}

type Summarizer struct {
	apiKey string
	apiHost string
}

func NewSummarizer(apiKey, apiHost string) *Summarizer {
	return &Summarizer{
		apiKey: apiKey,
		apiHost: apiHost,
	}
}

func (s *Summarizer) Summarize(article *models.Article) (*models.Summary, error) {
	requestBody := RapidAPISummarizeRequest{
		Lang: "en", // or make this configurable
		Text: fmt.Sprintf("%s\n\n%s", article.Title, article.Content),
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST", 
		"https://article-extractor-and-summarizer.p.rapidapi.com/summarize-text", 
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-rapidapi-key", s.apiKey)
	req.Header.Set("x-rapidapi-host", s.apiHost)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("rapidapi returned status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse RapidAPISummarizeResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, err
	}

	summary := &models.Summary{
		ArticleID: article.ID,
		Text:      apiResponse.Summary,
	}

	return summary, nil
}