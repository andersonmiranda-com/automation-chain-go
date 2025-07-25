package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"test-chain-go-cursor/config"
	nodesbase "test-chain-go-cursor/nodes/base"
	pipelinebase "test-chain-go-cursor/pipelines/base"
)

// PipelineConfig represents the configuration of a pipeline
type PipelineConfig struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Schedule    string                    `json:"schedule"`
	Nodes       []nodesbase.NodeDefinition `json:"nodes"`
}

func main() {
	log.Println("🚀 Starting LangChain Go - New Architecture")
	
	// Load credentials
	credentialsManager := config.NewCredentialsManager()
	if err := credentialsManager.LoadCredentials("config/credentials.json"); err != nil {
		log.Fatalf("Failed to load credentials: %v", err)
	}
	log.Println("✅ Credentials loaded successfully")
	
	// Load pipeline configuration
	pipelineConfig, err := loadPipelineConfig("telegram")
	if err != nil {
		log.Fatalf("Failed to load pipeline config: %v", err)
	}
	log.Println("✅ Pipeline configuration loaded successfully")
	
	// Create pipeline builder
	builder := pipelinebase.NewPipelineBuilder(credentialsManager)
	
	// Build pipeline
	pipeline, err := builder.BuildPipeline(pipelineConfig.Name, pipelineConfig.Nodes)
	if err != nil {
		log.Fatalf("Failed to build pipeline: %v", err)
	}
	
	log.Printf("📋 Pipeline configured with %d nodes", pipeline.GetNodeCount())
	
	// Execute pipeline
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := pipeline.Execute(ctx); err != nil {
		log.Fatalf("Pipeline execution failed: %v", err)
	}
	
	log.Println("🎉 Application completed successfully!")
}

// loadPipelineConfig loads pipeline configuration from JSON file
func loadPipelineConfig(name string) (*PipelineConfig, error) {
	filePath := "config/pipelines/" + name + ".json"
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	
	var config PipelineConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	
	return &config, nil
} 