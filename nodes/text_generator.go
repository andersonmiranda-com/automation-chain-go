package nodes

import (
	"context"
	"fmt"
	"log"

	"github.com/sashabaranov/go-openai"
)

// TextGeneratorNode generates motivational text using OpenAI
type TextGeneratorNode struct {
	client *openai.Client
}

// NewTextGeneratorNode creates a new text generator node
func NewTextGeneratorNode(apiKey string) *TextGeneratorNode {
	client := openai.NewClient(apiKey)
	return &TextGeneratorNode{
		client: client,
	}
}

// Name returns the node name
func (n *TextGeneratorNode) Name() string {
	return "TextGenerator"
}

// Execute generates motivational text using OpenAI
func (n *TextGeneratorNode) Execute(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
	log.Println("Generating motivational text...")
	
	// Create the prompt for motivational text
	prompt := `Genera un texto motivacional corto y poderoso en español. 
	El texto debe ser inspirador, positivo y que motive a las personas a alcanzar sus metas. 
	Debe tener entre 100-150 palabras y ser apropiado para compartir en redes sociales. 
	No incluyas hashtags ni emojis, solo el texto puro.`

	resp, err := n.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
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
			MaxTokens: 300,
			Temperature: 0.8,
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
	return map[string]any{
		"motivational_text": generatedText,
	}, nil
} 