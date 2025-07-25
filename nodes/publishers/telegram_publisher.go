package publishers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	telebot "gopkg.in/telebot.v3"
	"test-chain-go-cursor/nodes/base"
)

// TelegramPublisherNode publishes messages to a Telegram channel
type TelegramPublisherNode struct {
	bot         *telebot.Bot
	channelID   string
	config      base.NodeConfig
}

// NewTelegramPublisherNode creates a new Telegram publisher node
func NewTelegramPublisherNode(botToken, channelID string, config base.NodeConfig) (*TelegramPublisherNode, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	return &TelegramPublisherNode{
		bot:       bot,
		channelID: channelID,
		config:    config,
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
	if n.bot == nil {
		return fmt.Errorf("Telegram bot is not initialized")
	}
	
	if n.channelID == "" {
		return fmt.Errorf("channel ID is required")
	}
	
	return nil
}

// Execute publishes the motivational text to Telegram
func (n *TelegramPublisherNode) Execute(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
	log.Println("Publishing to Telegram...")
	
	// Get the motivational text from the previous node
	motivationalText, ok := input["motivational_text"].(string)
	if !ok {
		return nil, fmt.Errorf("motivational_text not found in input or not a string")
	}

	// Get message template from config or use default
	messageTemplate := "💪 *Mensaje Motivacional del Día*\n\n%s\n\n✨ ¡Que tengas un día increíble!"
	if val, exists := n.config.Parameters["message_template"]; exists {
		if template, ok := val.(string); ok {
			messageTemplate = template
		}
	}

	// Create the message with formatting
	message := fmt.Sprintf(messageTemplate, motivationalText)

	// Try different methods to send the message
	var err error
	
	// Method 1: Try as username (for public channels)
	if strings.HasPrefix(n.channelID, "@") {
		log.Printf("Trying to send to channel: %s", n.channelID)
		_, err = n.bot.Send(&telebot.Chat{Username: n.channelID}, message, &telebot.SendOptions{
			ParseMode: telebot.ModeMarkdown,
		})
		if err == nil {
			log.Printf("Message published successfully to channel: %s", n.channelID)
			return map[string]interface{}{
				"published": true,
				"channel_id": n.channelID,
				"platform": "telegram",
			}, nil
		}
		log.Printf("Failed with username method: %v", err)
	}
	
	// Method 2: Try as numeric ID (for private channels)
	if !strings.HasPrefix(n.channelID, "@") {
		// Remove @ if present and try to parse as number
		cleanID := strings.TrimPrefix(n.channelID, "@")
		if numericID, parseErr := strconv.ParseInt(cleanID, 10, 64); parseErr == nil {
			log.Printf("Trying to send to numeric channel ID: %d", numericID)
			_, err = n.bot.Send(&telebot.Chat{ID: numericID}, message, &telebot.SendOptions{
				ParseMode: telebot.ModeMarkdown,
			})
			if err == nil {
				log.Printf("Message published successfully to channel ID: %d", numericID)
				return map[string]interface{}{
					"published": true,
					"channel_id": n.channelID,
					"platform": "telegram",
				}, nil
			}
			log.Printf("Failed with numeric ID method: %v", err)
		}
	}
	
	// Method 3: Try as string ID (for channels like -1001234567890)
	if strings.HasPrefix(n.channelID, "-100") {
		if numericID, parseErr := strconv.ParseInt(n.channelID, 10, 64); parseErr == nil {
			log.Printf("Trying to send to supergroup ID: %d", numericID)
			_, err = n.bot.Send(&telebot.Chat{ID: numericID}, message, &telebot.SendOptions{
				ParseMode: telebot.ModeMarkdown,
			})
			if err == nil {
				log.Printf("Message published successfully to supergroup: %d", numericID)
				return map[string]interface{}{
					"published": true,
					"channel_id": n.channelID,
					"platform": "telegram",
				}, nil
			}
			log.Printf("Failed with supergroup ID method: %v", err)
		}
	}

	// If all methods failed, return the last error
	return nil, fmt.Errorf("failed to send message to Telegram after trying all methods. Last error: %w", err)
} 