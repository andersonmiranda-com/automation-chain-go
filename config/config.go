package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration variables
type Config struct {
	OpenAIApiKey      string
	TelegramBotToken  string
	TelegramChannelID string
}

// Load loads configuration from environment variables
func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config := &Config{
		OpenAIApiKey:      getEnv("OPENAI_API_KEY", ""),
		TelegramBotToken:  getEnv("TELEGRAM_BOT_TOKEN", ""),
		TelegramChannelID: getEnv("TELEGRAM_CHANNEL_ID", ""),
	}

	// Validate required configuration
	if config.OpenAIApiKey == "" {
		log.Fatal("OPENAI_API_KEY is required")
	}
	if config.TelegramBotToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is required")
	}
	if config.TelegramChannelID == "" {
		log.Fatal("TELEGRAM_CHANNEL_ID is required")
	}

	return config
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
} 