package services

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/telebot.v3"
)

// TelegramService handles Telegram operations
type TelegramService struct {
	bot       *telebot.Bot
	token     string
	channelID string
	ready     bool
}

// NewTelegram creates a new Telegram service
func NewTelegram() *TelegramService {
	return &TelegramService{}
}

// LoadConfig loads configuration from map
func (s *TelegramService) LoadConfig(config map[string]interface{}) error {
	// Extract token
	if token, exists := config["token"]; exists {
		if tokenStr, ok := token.(string); ok {
			s.token = tokenStr
		} else {
			return fmt.Errorf("invalid token format")
		}
	} else {
		return fmt.Errorf("token is required")
	}

	// Extract channel ID
	if channelID, exists := config["channel_id"]; exists {
		if channelIDStr, ok := channelID.(string); ok {
			s.channelID = channelIDStr
		} else {
			return fmt.Errorf("invalid channel ID format")
		}
	} else {
		return fmt.Errorf("channel ID is required")
	}

	// Validate
	if err := s.validate(); err != nil {
		return err
	}

	// Create bot
	bot, err := telebot.NewBot(telebot.Settings{
		Token: s.token,
	})
	if err != nil {
		return fmt.Errorf("failed to create bot: %w", err)
	}

	s.bot = bot
	s.ready = true

	return nil
}

// validate checks if configuration is valid
func (s *TelegramService) validate() error {
	if s.token == "" {
		return fmt.Errorf("token is required")
	}

	if len(s.token) < 30 {
		return fmt.Errorf("token too short")
	}

	if s.channelID == "" {
		return fmt.Errorf("channel ID is required")
	}

	if len(s.channelID) < 3 {
		return fmt.Errorf("channel ID too short")
	}

	return nil
}

// SendMessage sends a message to the configured channel
func (s *TelegramService) SendMessage(ctx context.Context, text string) error {
	if !s.ready {
		return fmt.Errorf("Telegram service not initialized")
	}

	// Parse channel ID
	var chat telebot.Chat
	if strings.HasPrefix(s.channelID, "@") {
		// Public channel
		chat = telebot.Chat{Username: strings.TrimPrefix(s.channelID, "@")}
	} else {
		// Private channel (numeric ID)
		if numericID, err := strconv.ParseInt(s.channelID, 10, 64); err == nil {
			chat = telebot.Chat{ID: numericID}
		} else {
			return fmt.Errorf("invalid channel ID format: %s", s.channelID)
		}
	}

	// Send message
	_, err := s.bot.Send(&chat, text)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

// IsReady returns if service is ready
func (s *TelegramService) IsReady() bool {
	return s.ready
}

// GetChannelID returns the channel ID
func (s *TelegramService) GetChannelID() string {
	return s.channelID
} 