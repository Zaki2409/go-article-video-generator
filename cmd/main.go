package main

import (
	"fmt"
	"log"
	"net/http"
	"internal/handlers"
	"internal/services"
	"pkg/config"
	"api/v1"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize services
	summarizer := services.NewSummarizer(cfg.ChatGPTAPIKey)
	videoGenerator := services.NewVideoGenerator(cfg.ZapierWebhookURL)

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