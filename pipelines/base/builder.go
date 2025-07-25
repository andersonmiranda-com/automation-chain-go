package base

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"automation-chain/nodes/ai"
	"automation-chain/nodes/base"
	"automation-chain/nodes/publishers"
)

// PipelineBuilder builds pipelines from configuration
type PipelineBuilder struct {
	credentials map[string]interface{}
}

// NewPipelineBuilder creates a new pipeline builder
func NewPipelineBuilder() *PipelineBuilder {
	// Load credentials from config/credentials.json
	credentials, err := loadCredentials("config/credentials.json")
	if err != nil {
		log.Printf("Warning: Failed to load credentials: %v", err)
		credentials = make(map[string]interface{})
	}

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

		// Get OpenAI config from credentials
		if openaiServices, exists := b.credentials["openai"]; exists {
			if openaiMap, ok := openaiServices.(map[string]interface{}); ok {
				if openaiConfig, exists := openaiMap[openaiCredential]; exists {
					if configMap, ok := openaiConfig.(map[string]interface{}); ok {
						// Add OpenAI config to node parameters
						nodeConfig.Parameters["openai"] = configMap
					}
				}
			}
		}

		return ai.NewTextGeneratorNode(nodeConfig)

	case "telegram_publisher":
		// Get Telegram credential from node definition
		telegramCredential := nodeDef.Credentials
		if telegramCredential == "" {
			return nil, fmt.Errorf("telegram_publisher node requires 'credentials' field")
		}

		// Get Telegram config from credentials
		if telegramServices, exists := b.credentials["telegram"]; exists {
			if telegramMap, ok := telegramServices.(map[string]interface{}); ok {
				if telegramConfig, exists := telegramMap[telegramCredential]; exists {
					if configMap, ok := telegramConfig.(map[string]interface{}); ok {
						// Add Telegram config to node parameters
						nodeConfig.Parameters["telegram"] = configMap
					}
				}
			}
		}

		return publishers.NewTelegramPublisherNode(nodeConfig)

	default:
		return nil, fmt.Errorf("unknown node type: %s", nodeDef.Type)
	}
}

// loadCredentials loads credentials from JSON file
func loadCredentials(filePath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var credentials map[string]interface{}
	if err := json.Unmarshal(data, &credentials); err != nil {
		return nil, err
	}

	return credentials, nil
} 