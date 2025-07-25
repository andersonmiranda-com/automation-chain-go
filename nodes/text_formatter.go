package nodes

import (
	"context"
	"fmt"
	"log"
	"strings"
)

// TextFormatterNode formats the generated text before publishing
type TextFormatterNode struct{}

// NewTextFormatterNode creates a new text formatter node
func NewTextFormatterNode() *TextFormatterNode {
	return &TextFormatterNode{}
}

// Name returns the node name
func (n *TextFormatterNode) Name() string {
	return "TextFormatter"
}

// Execute formats the motivational text
func (n *TextFormatterNode) Execute(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
	log.Println("Formatting text...")
	
	// Get the motivational text from the previous node
	motivationalText, ok := input["motivational_text"].(string)
	if !ok {
		return nil, fmt.Errorf("motivational_text not found in input or not a string")
	}

	// Apply formatting
	formattedText := n.formatText(motivationalText)
	
	log.Printf("Text formatted successfully")
	
	// Return the formatted text for the next node
	return map[string]any{
		"motivational_text": formattedText,
		"original_text":     motivationalText,
	}, nil
}

// formatText applies formatting to the text
func (n *TextFormatterNode) formatText(text string) string {
	// Remove extra whitespace
	text = strings.TrimSpace(text)
	
	// Ensure proper paragraph breaks
	text = strings.ReplaceAll(text, "\n\n", "\n")
	
	// Add a period at the end if missing
	if !strings.HasSuffix(text, ".") && !strings.HasSuffix(text, "!") && !strings.HasSuffix(text, "?") {
		text += "."
	}
	
	return text
} 