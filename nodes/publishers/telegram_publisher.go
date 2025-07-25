package publishers

import (
	"context"
	"fmt"
	"log"

	"test-chain-go-cursor/nodes/base"
	"test-chain-go-cursor/services"
)

// TelegramPublisherNode publishes messages to a Telegram channel
type TelegramPublisherNode struct {
	telegram *services.TelegramService
	config   base.NodeConfig
}

// NewTelegramPublisherNode creates a new Telegram publisher node
func NewTelegramPublisherNode(config base.NodeConfig) (*TelegramPublisherNode, error) {
	telegram := services.NewTelegram()

	// Load Telegram config from node config
	if telegramConfig, exists := config.Parameters["telegram"]; exists {
		if telegramMap, ok := telegramConfig.(map[string]interface{}); ok {
			if err := telegram.LoadConfig(telegramMap); err != nil {
				return nil, fmt.Errorf("failed to load Telegram config: %w", err)
			}
		}
	}

	return &TelegramPublisherNode{
		telegram: telegram,
		config:   config,
	}, nil
}

// Name returns the node name
func (n *TelegramPublisherNode) Name() string {
	return n.config.Name
}

// Config returns the node configuration
func (n *TelegramPublisherNode) Config() base.NodeConfig {
	return n.config
}

// Validate validates the node configuration
func (n *TelegramPublisherNode) Validate() error {
	if !n.telegram.IsReady() {
		return fmt.Errorf("Telegram service is not initialized")
	}

	return nil
}

// Execute publishes the generated text to Telegram
func (n *TelegramPublisherNode) Execute(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
	log.Println("Publishing to Telegram...")

	// Get the generated text from the previous node
	generatedText, ok := input["generated_text"].(string)
	if !ok {
		return nil, fmt.Errorf("generated_text not found in input or not a string")
	}

	// Get message template from config or use default
	messageTemplate := "💪 *Daily Motivation*\n\n%s\n\n✨ Have an amazing day!"
	if val, exists := n.config.Parameters["message_template"]; exists {
		if template, ok := val.(string); ok {
			messageTemplate = template
		}
	}

	// Create the message with formatting
	message := fmt.Sprintf(messageTemplate, generatedText)

	// Send message using Telegram service
	if err := n.telegram.SendMessage(ctx, message); err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	log.Printf("Message published successfully to channel: %s", n.telegram.GetChannelID())

	return map[string]interface{}{
		"published":  true,
		"channel_id": n.telegram.GetChannelID(),
		"platform":   "telegram",
	}, nil
} 