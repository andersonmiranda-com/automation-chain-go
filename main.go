package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	nodesbase "automation-chain/nodes/base"
	pipelinebase "automation-chain/pipelines/base"
)

// PipelineConfig represents the configuration of a pipeline
type PipelineConfig struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Schedule    string                    `json:"schedule"`
	Nodes       []nodesbase.NodeDefinition `json:"nodes"`
}

func main() {
	// Parse command line arguments
	pipelineName := flag.String("pipeline", "telegram", "Name of the pipeline to execute")
	flag.Parse()

	log.Printf("🚀 Starting Automation Chain Go - Pipeline: %s", *pipelineName)

	// Load pipeline configuration
	pipelineConfig, err := loadPipelineConfig(*pipelineName)
	if err != nil {
		log.Fatalf("Failed to load pipeline config: %v", err)
	}
	log.Println("✅ Pipeline configuration loaded successfully")

	// Create pipeline builder (loads credentials internally)
	builder := pipelinebase.NewPipelineBuilder()

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