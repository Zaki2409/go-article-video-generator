package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/Zaki2409/go-article-video-generator/internal/models"
    "github.com/Zaki2409/go-article-video-generator/internal/services"
)

type APIHandler struct {
	summarizer     *services.Summarizer
	videoGenerator *services.VideoGenerator
}

func NewAPIHandler(summarizer *services.Summarizer, videoGenerator *services.VideoGenerator) *APIHandler {
	return &APIHandler{
		summarizer:     summarizer,
		videoGenerator: videoGenerator,
	}
}

func (h *APIHandler) SummarizeArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	summary, err := h.summarizer.Summarize(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

func (h *APIHandler) GenerateVideo(w http.ResponseWriter, r *http.Request) {
	var videoReq models.VideoRequest
	if err := json.NewDecoder(r.Body).Decode(&videoReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	videoResp, err := h.videoGenerator.GenerateVideo(&videoReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videoResp)
}