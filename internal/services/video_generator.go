package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"internal/models"
)

type VideoGenerator struct {
	webhookURL string
}

func NewVideoGenerator(webhookURL string) *VideoGenerator {
	return &VideoGenerator{webhookURL: webhookURL}
}

func (v *VideoGenerator) GenerateVideo(videoReq *models.VideoRequest) (*models.VideoResponse, error) {
	requestBody := map[string]string{
		"title": videoReq.Title,
		"text":  videoReq.Text,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(v.webhookURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("zapier webhook returned status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var videoResp models.VideoResponse
	if err := json.Unmarshal(body, &videoResp); err != nil {
		return nil, err
	}

	return &videoResp, nil
}