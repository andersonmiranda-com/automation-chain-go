package ai

import (
	"context"
	"fmt"
	"log"

	"automation-chain/nodes/base"
	"automation-chain/services"
)

// TextGeneratorNode generates text using OpenAI
type TextGeneratorNode struct {
	openai *services.OpenAIService
	config base.NodeConfig
}

// NewTextGeneratorNode creates a new text generator node
func NewTextGeneratorNode(config base.NodeConfig) (*TextGeneratorNode, error) {
	openai := services.NewOpenAI()

	// Load OpenAI config from node config
	if openaiConfig, exists := config.Parameters["openai"]; exists {
		if openaiMap, ok := openaiConfig.(map[string]interface{}); ok {
			if err := openai.LoadConfig(openaiMap); err != nil {
				return nil, fmt.Errorf("failed to load OpenAI config: %w", err)
			}
		}
	}

	return &TextGeneratorNode{
		openai: openai,
		config: config,
	}, nil
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
	if !n.openai.IsReady() {
		return fmt.Errorf("OpenAI service is not initialized")
	}

	// Validate required parameters
	requiredParams := []string{"prompt_template"}
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

	// Get prompt template
	promptTemplate := n.config.Parameters["prompt_template"].(string)

	// Process prompt template with input data
	prompt := processTemplate(promptTemplate, input)

	// Generate text using OpenAI service
	generatedText, err := n.openai.GenerateText(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate text: %w", err)
	}

	log.Printf("Generated text: %s", generatedText)

	// Return the generated text for the next node
	return map[string]interface{}{
		"generated_text": generatedText,
		"model_used":     n.openai.GetModel(),
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