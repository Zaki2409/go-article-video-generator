package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/Zaki2409/go-article-video-generator/internal/models"
)

type VideoPromptRequest struct {
	Style     string `json:"style"`
	Frames    int    `json:"frames"`
	Language  string `json:"language"`
	NoQueue   int    `json:"noqueue"`
	ScenePrompts struct {
		MainScene    string   `json:"mainScene"`
		Description  string   `json:"description"`
		VisualStyle  string   `json:"visualStyle"`
		KeyMoments   []string `json:"keyMoments"`
	} `json:"scenePrompts"`
}

type VideoPromptResponse struct {
	Storyboard []struct {
		Frame       int    `json:"frame"`
		Prompt      string `json:"prompt"`
		Description string `json:"description"`
	} `json:"storyboard"`
	VideoURL string `json:"videoUrl,omitempty"`
}

type VideoGenerator struct {
	apiKey  string
	apiHost string
}

func NewVideoGenerator(apiKey, apiHost string) *VideoGenerator {
	return &VideoGenerator{
		apiKey:  apiKey,
		apiHost: apiHost,
	}
}

func (v *VideoGenerator) GenerateVideo(videoReq *models.VideoRequest) (*models.VideoResponse, error) {
	requestBody := VideoPromptRequest{
		Style:    "detailed",
		Frames:   8,
		Language: "en",
		NoQueue:  1,
		ScenePrompts: struct {
			MainScene    string   `json:"mainScene"`
			Description  string   `json:"description"`
			VisualStyle  string   `json:"visualStyle"`
			KeyMoments   []string `json:"keyMoments"`
		}{
			MainScene:   videoReq.Title,
			Description: videoReq.Text,
			VisualStyle: "cinematic",
			KeyMoments:  []string{
				"Opening scene",
				"Main action",
				"Climax",
				"Conclusion",
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST", 
		"https://openai-sora-ai-video-prompt-generator-cinematic-api.p.rapidapi.com/generateStoryboard", 
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-rapidapi-key", v.apiKey)
	req.Header.Set("x-rapidapi-host", v.apiHost)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("video api returned status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse VideoPromptResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, err
	}

	videoResp := &models.VideoResponse{
		ID:       videoReq.SummaryID,
		VideoURL: apiResponse.VideoURL,
	}

	return videoResp, nil
}