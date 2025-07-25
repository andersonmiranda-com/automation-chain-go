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
	pipelineName := flag.String("pipeline", "", "Pipeline to execute")
	flag.Parse()

	if *pipelineName == "" {
		log.Fatal("Pipeline name required")
	}

	log.Printf("🚀 Executing pipeline: %s", *pipelineName)

	// Ejecutar pipeline
	if err := runPipeline(*pipelineName); err != nil {
		log.Fatalf("Pipeline execution failed: %v", err)
	}

	log.Println("✅ Pipeline completed successfully")
}

func runPipeline(name string) error {
	// Tu lógica actual de ejecución
	pipelineConfig, err := loadPipelineConfig(name)
	if err != nil {
		return err
	}

	builder := pipelinebase.NewPipelineBuilder()
	pipeline, err := builder.BuildPipeline(pipelineConfig.Name, pipelineConfig.Nodes)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return pipeline.Execute(ctx)
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