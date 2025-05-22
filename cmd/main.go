package main

import (
	"fmt"
	"log"
	"net/http"
	 "github.com/Zaki2409/go-article-video-generator/api/v1"
    "github.com/Zaki2409/go-article-video-generator/internal/handlers"
    "github.com/Zaki2409/go-article-video-generator/internal/services"
    "github.com/Zaki2409/go-article-video-generator/pkg/config"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize services
	summarizer := services.NewSummarizer(cfg.RapidAPIKey, cfg.RapidAPIHost)
	videoGenerator := services.NewVideoGenerator(cfg.VideoRapidAPIKey, cfg.VideoRapidAPIHost)

	// Initialize handlers
	apiHandler := handlers.NewAPIHandler(summarizer, videoGenerator)

	// Set up routes
	mux := http.NewServeMux()
	v1.SetupRoutes(mux, apiHandler)

	// Start server
	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}