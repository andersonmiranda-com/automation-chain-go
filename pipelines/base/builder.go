package base

import (
	"fmt"
	"log"

	"test-chain-go-cursor/config"
	"test-chain-go-cursor/nodes/ai"
	"test-chain-go-cursor/nodes/base"
	"test-chain-go-cursor/nodes/publishers"
)

// PipelineBuilder builds pipelines from configuration
type PipelineBuilder struct {
	credentials *config.CredentialsManager
}

// NewPipelineBuilder creates a new pipeline builder
func NewPipelineBuilder(credentials *config.CredentialsManager) *PipelineBuilder {
	return &PipelineBuilder{
		credentials: credentials,
	}
}

// BuildPipeline builds a pipeline from node definitions
func (b *PipelineBuilder) BuildPipeline(name string, nodeDefs []base.NodeDefinition) (*Pipeline, error) {
	log.Printf("Building pipeline: %s", name)
	
	pipeline := NewPipeline(name)
	
	for _, nodeDef := range nodeDefs {
		node, err := b.createNode(nodeDef)
		if err != nil {
			return nil, fmt.Errorf("failed to create node %s: %w", nodeDef.Name, err)
		}
		
		pipeline.AddNode(node)
	}
	
	log.Printf("Pipeline %s built successfully with %d nodes", name, pipeline.GetNodeCount())
	return pipeline, nil
}

// createNode creates a node from node definition
func (b *PipelineBuilder) createNode(nodeDef base.NodeDefinition) (base.Node, error) {
	// Create node config
	nodeConfig := base.NodeConfig{
		ID:         nodeDef.ID,
		Type:       nodeDef.Type,
		Name:       nodeDef.Name,
		Parameters: nodeDef.Config,
	}
	
	switch nodeDef.Type {
	case "text_generator":
		// Get OpenAI credential from node definition
		openaiCredential := nodeDef.Credentials
		if openaiCredential == "" {
			openaiCredential = "default" // fallback to default
		}
		
		apiKey, err := b.credentials.GetOpenAICredential(openaiCredential)
		if err != nil {
			return nil, fmt.Errorf("failed to get OpenAI credential '%s': %w", openaiCredential, err)
		}
		
		return ai.NewTextGeneratorNode(apiKey, nodeConfig), nil
		
	case "telegram_publisher":
		// Get Telegram credential from node definition
		telegramCredential := nodeDef.Credentials
		if telegramCredential == "" {
			return nil, fmt.Errorf("telegram_publisher node requires 'credentials' field")
		}
		
		token, channelID, err := b.credentials.GetTelegramCredential(telegramCredential)
		if err != nil {
			return nil, fmt.Errorf("failed to get Telegram credential '%s': %w", telegramCredential, err)
		}
		
		return publishers.NewTelegramPublisherNode(token, channelID, nodeConfig)
		
	default:
		return nil, fmt.Errorf("unknown node type: %s", nodeDef.Type)
	}
} 