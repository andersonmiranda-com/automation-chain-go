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
	Name         string                    `json:"name"`
	Description  string                    `json:"description"`
	Schedule     string                    `json:"schedule"`
	Credentials  map[string]string         `json:"credentials"`
	Nodes        []nodesbase.NodeDefinition `json:"nodes"`
}

func main() {
	log.Println("🚀 Multiple Credentials Example")
	
	// Load credentials
	credentialsManager := config.NewCredentialsManager()
	if err := credentialsManager.LoadCredentials("config/credentials.json"); err != nil {
		log.Fatalf("Failed to load credentials: %v", err)
	}
	log.Println("✅ Credentials loaded successfully")
	
	// Create pipeline builder
	builder := pipelinebase.NewPipelineBuilder(credentialsManager)
	
	// Example 1: Pipeline con credenciales por defecto
	log.Println("\n📋 Example 1: Default credentials pipeline")
	pipeline1, err := loadAndBuildPipeline(builder, "telegram")
	if err != nil {
		log.Printf("❌ Failed to build pipeline 1: %v", err)
	} else {
		log.Printf("✅ Pipeline 1 built: %s with %d nodes", pipeline1.GetName(), pipeline1.GetNodeCount())
	}
	
	// Example 2: Pipeline con credenciales premium
	log.Println("\n📋 Example 2: Premium credentials pipeline")
	pipeline2, err := loadAndBuildPipeline(builder, "telegram_news")
	if err != nil {
		log.Printf("❌ Failed to build pipeline 2: %v", err)
	} else {
		log.Printf("✅ Pipeline 2 built: %s with %d nodes", pipeline2.GetName(), pipeline2.GetNodeCount())
	}
	
	// Example 3: Mostrar credenciales diferentes
	log.Println("\n🔑 Credential Comparison:")
	
	// OpenAI credentials
	defaultKey, _ := credentialsManager.GetOpenAICredential("default")
	premiumKey, _ := credentialsManager.GetOpenAICredential("premium")
	log.Printf("OpenAI Default: %s...", defaultKey[:10])
	log.Printf("OpenAI Premium: %s...", premiumKey[:10])
	
	// Telegram credentials
	motivationalToken, motivationalChannel, _ := credentialsManager.GetTelegramCredential("motivational_bot")
	newsToken, newsChannel, _ := credentialsManager.GetTelegramCredential("news_bot")
	log.Printf("Telegram Motivational: %s -> %s", motivationalToken[:10], motivationalChannel)
	log.Printf("Telegram News: %s -> %s", newsToken[:10], newsChannel)
	
	log.Println("\n🎉 Multiple credentials example completed!")
}

// loadAndBuildPipeline loads and builds a pipeline
func loadAndBuildPipeline(builder *pipelinebase.PipelineBuilder, pipelineName string) (*pipelinebase.Pipeline, error) {
	// Load pipeline configuration
	pipelineConfig, err := loadPipelineConfig(pipelineName)
	if err != nil {
		return nil, err
	}
	
	// Build pipeline
	pipeline, err := builder.BuildPipeline(pipelineConfig.Name, pipelineConfig.Nodes, pipelineConfig.Credentials)
	if err != nil {
		return nil, err
	}
	
	return pipeline, nil
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