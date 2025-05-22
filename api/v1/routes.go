package v1

import (
	"net/http"
	"internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.APIHandler) {
	mux.HandleFunc("POST /api/v1/summarize", handler.SummarizeArticle)
	mux.HandleFunc("POST /api/v1/generate-video", handler.GenerateVideo)
}