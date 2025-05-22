package config

import (
	"log"
	"os"
)

type Config struct {
	RapidAPIKey         string
	RapidAPIHost        string
	VideoRapidAPIKey    string
	VideoRapidAPIHost   string
}

func LoadConfig() *Config {
	rapidAPIKey := os.Getenv("RAPIDAPI_KEY")
	if rapidAPIKey == "" {
		log.Fatal("RAPIDAPI_KEY environment variable not set")
	}

	rapidAPIHost := os.Getenv("RAPIDAPI_HOST")
	if rapidAPIHost == "" {
		log.Fatal("RAPIDAPI_HOST environment variable not set")
	}

	videoRapidAPIKey := os.Getenv("VIDEO_RAPIDAPI_KEY")
	if videoRapidAPIKey == "" {
		log.Fatal("VIDEO_RAPIDAPI_KEY environment variable not set")
	}

	videoRapidAPIHost := os.Getenv("VIDEO_RAPIDAPI_HOST")
	if videoRapidAPIHost == "" {
		log.Fatal("VIDEO_RAPIDAPI_HOST environment variable not set")
	}

	return &Config{
		RapidAPIKey:         rapidAPIKey,
		RapidAPIHost:        rapidAPIHost,
		VideoRapidAPIKey:    videoRapidAPIKey,
		VideoRapidAPIHost:   videoRapidAPIHost,
	}
}