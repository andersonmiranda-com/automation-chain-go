package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
)

// OpenAIService handles OpenAI operations
type OpenAIService struct {
	client *openai.Client
	apiKey string
	model  string
	ready  bool
}

// NewOpenAI creates a new OpenAI service
func NewOpenAI() *OpenAIService {
	return &OpenAIService{
		model: "gpt-3.5-turbo", // default
	}
}

// LoadConfig loads configuration from map
func (s *OpenAIService) LoadConfig(config map[string]interface{}) error {
	// Extract API key
	if apiKey, exists := config["api_key"]; exists {
		if apiKeyStr, ok := apiKey.(string); ok {
			s.apiKey = apiKeyStr
		} else {
			return fmt.Errorf("invalid API key format")
		}
	} else {
		return fmt.Errorf("API key is required")
	}

	// Extract model (optional)
	if model, exists := config["model"]; exists {
		if modelStr, ok := model.(string); ok {
			s.model = modelStr
		}
	}

	// Validate
	if err := s.validate(); err != nil {
		return err
	}

	// Create client
	s.client = openai.NewClient(s.apiKey)
	s.ready = true

	return nil
}

// validate checks if configuration is valid
func (s *OpenAIService) validate() error {
	if s.apiKey == "" {
		return fmt.Errorf("API key is required")
	}

	if !strings.HasPrefix(s.apiKey, "sk-") {
		return fmt.Errorf("invalid API key format")
	}

	if len(s.apiKey) < 20 {
		return fmt.Errorf("API key too short")
	}

	return nil
}

// GenerateText generates text using OpenAI
func (s *OpenAIService) GenerateText(ctx context.Context, prompt string) (string, error) {
	if !s.ready {
		return "", fmt.Errorf("OpenAI service not initialized")
	}

	resp, err := s.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: s.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})

	if err != nil {
		return "", fmt.Errorf("failed to generate text: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}

// IsReady returns if service is ready
func (s *OpenAIService) IsReady() bool {
	return s.ready
}

// GetModel returns current model
func (s *OpenAIService) GetModel() string {
	return s.model
} 