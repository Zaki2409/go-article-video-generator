package config 

import (
	"log"
	"os"
)

type Config struct {
	ChatGPTAPIKey string
	ZapierWebhookURL string 
}

func LoadConfig () *Config {
	ChatGPTKEY := os.Getenv("CHATGPT_API_KEY") 
	if ChatGPTKEY == "" {
		log.Fatal("Chat gpt api key environment variable is not set")
	} 
	zapierURL :-= os.Getenv("ZAPEIR_WEBHOOK_URL")
	if zapierURL == "" {
		log.Fatal("ZAPEIR_WEBHOOK_URL environment variable is not set")
	}

	return &Config {
		ChatGPTAPIKey : ChatGPTKEY
		ZapierWebhookURL : zapierURL
	}
} 