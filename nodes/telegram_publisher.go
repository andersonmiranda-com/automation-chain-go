package nodes

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	telebot "gopkg.in/telebot.v3"
)

// TelegramPublisherNode publishes messages to a Telegram channel
type TelegramPublisherNode struct {
	bot         *telebot.Bot
	channelID   string
}

// NewTelegramPublisherNode creates a new Telegram publisher node
func NewTelegramPublisherNode(botToken, channelID string) (*TelegramPublisherNode, error) {
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
	}, nil
}

// Name returns the node name
func (n *TelegramPublisherNode) Name() string {
	return "TelegramPublisher"
}

// Execute publishes the motivational text to Telegram
func (n *TelegramPublisherNode) Execute(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
	log.Println("Publishing to Telegram...")
	
	// Get the motivational text from the previous node
	motivationalText, ok := input["motivational_text"].(string)
	if !ok {
		return nil, fmt.Errorf("motivational_text not found in input or not a string")
	}

	// Create the message with some formatting
	message := fmt.Sprintf("💪 *Mensaje Motivacional del Día*\n\n%s\n\n✨ ¡Que tengas un día increíble!", motivationalText)

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
				return map[string]any{
					"published": true,
					"channel_id": n.channelID,
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
				return map[string]any{
					"published": true,
					"channel_id": n.channelID,
				}, nil
			}
			log.Printf("Failed with supergroup ID method: %v", err)
		}
	}

	// If all methods failed, return the last error
	return nil, fmt.Errorf("failed to send message to Telegram after trying all methods. Last error: %w", err)
} 