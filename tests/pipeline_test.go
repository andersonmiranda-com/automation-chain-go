package tests

import (
	"testing"

	"test-chain-go-cursor/config"
	nodesbase "test-chain-go-cursor/nodes/base"
	pipelinebase "test-chain-go-cursor/pipelines/base"
)

func TestPipelineBuilder(t *testing.T) {
	// Create credentials manager with test data
	credentialsManager := config.NewCredentialsManager()
	testCredentials := map[string]interface{}{
		"openai": map[string]interface{}{
			"default": "sk-test-default-openai-key",
		},
		"telegram": map[string]interface{}{
			"motivational_bot": map[string]interface{}{
				"token":      "test-motivational-bot-token",
				"channel_id": "@test_motivational_channel",
			},
		},
	}
	
	// Manually set test credentials
	credentialsManager.SetCredentials(testCredentials)
	
	// Create pipeline builder
	builder := pipelinebase.NewPipelineBuilder(credentialsManager)
	
	// Define a simple pipeline for testing
	nodeDefs := []nodesbase.NodeDefinition{
		{
			ID:          "text_generator",
			Type:        "text_generator",
			Name:        "Test Text Generator",
			Credentials: "default",
			Config: map[string]interface{}{
				"model": "gpt-3.5-turbo",
				"prompt_template": "Generate a motivational text of 50 words.",
				"max_tokens": 100,
				"temperature": 0.7,
			},
		},
	}
	
	// Build pipeline
	pipeline, err := builder.BuildPipeline("test_pipeline", nodeDefs)
	
	if err != nil {
		t.Fatalf("Failed to build pipeline: %v", err)
	}
	
	// Verify pipeline has correct number of nodes
	expectedNodes := 1
	if pipeline.GetNodeCount() != expectedNodes {
		t.Errorf("Expected %d nodes, got %d", expectedNodes, pipeline.GetNodeCount())
	}
	
	// Verify pipeline name
	expectedName := "test_pipeline"
	if pipeline.GetName() != expectedName {
		t.Errorf("Expected pipeline name %s, got %s", expectedName, pipeline.GetName())
	}
	
	t.Logf("Pipeline built successfully with %d nodes", pipeline.GetNodeCount())
} 