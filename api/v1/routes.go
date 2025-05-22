package v1

import (
	"net/http"
	"github.com/Zaki2409/go-article-video-generator/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.APIHandler) {
	mux.HandleFunc("POST /api/v1/summarize", handler.SummarizeArticle)
	mux.HandleFunc("POST /api/v1/generate-video", handler.GenerateVideo)
}