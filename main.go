package main

import (
	"context"
	"log"
	"time"

	"test-chain-go-cursor/config"
	"test-chain-go-cursor/nodes"
	"test-chain-go-cursor/pipeline"
)

func main() {
	log.Println("🚀 Starting LangChain Go - Motivational Text Generator")
	
	// Load configuration
	cfg := config.Load()
	log.Println("✅ Configuration loaded successfully")
	
	// Create pipeline
	p := pipeline.NewPipeline()
	
	// Create and add nodes
	textGenerator := nodes.NewTextGeneratorNode(cfg.OpenAIApiKey)
	p.AddNode(textGenerator)
	
	telegramPublisher, err := nodes.NewTelegramPublisherNode(cfg.TelegramBotToken, cfg.TelegramChannelID)
	if err != nil {
		log.Fatalf("Failed to create Telegram publisher: %v", err)
	}
	p.AddNode(telegramPublisher)
	
	log.Printf("📋 Pipeline configured with %d nodes", p.GetNodeCount())
	
	// Execute pipeline
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := p.Execute(ctx); err != nil {
		log.Fatalf("Pipeline execution failed: %v", err)
	}
	
	log.Println("🎉 Application completed successfully!")
} 