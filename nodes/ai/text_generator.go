package ai

import (
	"context"
	"fmt"
	"log"

	"github.com/sashabaranov/go-openai"
	"test-chain-go-cursor/nodes/base"
)

// TextGeneratorNode generates text using OpenAI
type TextGeneratorNode struct {
	client *openai.Client
	config base.NodeConfig
}

// NewTextGeneratorNode creates a new text generator node
func NewTextGeneratorNode(apiKey string, config base.NodeConfig) *TextGeneratorNode {
	client := openai.NewClient(apiKey)
	return &TextGeneratorNode{
		client: client,
		config: config,
	}
}

// Name returns the node name
func (n *TextGeneratorNode) Name() string {
	return n.config.Name
}

// Config returns the node configuration
func (n *TextGeneratorNode) Config() base.NodeConfig {
	return n.config
}

// Validate validates the node configuration
func (n *TextGeneratorNode) Validate() error {
	if n.client == nil {
		return fmt.Errorf("OpenAI client is not initialized")
	}
	
	// Validate required parameters
	requiredParams := []string{"model", "prompt_template"}
	for _, param := range requiredParams {
		if _, exists := n.config.Parameters[param]; !exists {
			return fmt.Errorf("required parameter '%s' not found in configuration", param)
		}
	}
	
	return nil
}

// Execute generates text using OpenAI
func (n *TextGeneratorNode) Execute(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
	log.Println("Generating text with OpenAI...")
	
	// Get configuration parameters
	model := n.config.Parameters["model"].(string)
	promptTemplate := n.config.Parameters["prompt_template"].(string)
	maxTokens := 300
	temperature := 0.8
	
	// Get optional parameters
	if val, exists := n.config.Parameters["max_tokens"]; exists {
		if maxT, ok := val.(float64); ok {
			maxTokens = int(maxT)
		}
	}
	
	if val, exists := n.config.Parameters["temperature"]; exists {
		if temp, ok := val.(float64); ok {
			temperature = temp
		}
	}
	
	// Process prompt template with input data
	prompt := processTemplate(promptTemplate, input)
	
	resp, err := n.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Eres un experto en motivación y desarrollo personal. Generas textos inspiradores y motivacionales.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens:   maxTokens,
			Temperature: float32(temperature),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to generate text: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no response from OpenAI")
	}

	generatedText := resp.Choices[0].Message.Content
	log.Printf("Generated text: %s", generatedText)

	// Return the generated text for the next node
	return map[string]interface{}{
		"motivational_text": generatedText,
		"model_used":        model,
		"tokens_used":       resp.Usage.TotalTokens,
	}, nil
}

// processTemplate processes a template string with input data
func processTemplate(template string, input map[string]interface{}) string {
	// Simple template processing - replace {{key}} with values
	// In a real implementation, you might want to use a proper template engine
	result := template
	
	// For now, just return the template as-is
	// TODO: Implement proper template processing
	return result
} 