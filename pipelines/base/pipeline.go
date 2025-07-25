package base

import (
	"context"
	"fmt"
	"log"

	"test-chain-go-cursor/nodes/base"
)

// Pipeline orchestrates the execution of nodes
type Pipeline struct {
	name  string
	nodes []base.Node
}

// NewPipeline creates a new pipeline instance
func NewPipeline(name string) *Pipeline {
	return &Pipeline{
		name:  name,
		nodes: make([]base.Node, 0),
	}
}

// AddNode adds a node to the pipeline
func (p *Pipeline) AddNode(node base.Node) {
	p.nodes = append(p.nodes, node)
	log.Printf("Added node: %s to pipeline: %s", node.Name(), p.name)
}

// Execute runs all nodes in the pipeline sequentially
func (p *Pipeline) Execute(ctx context.Context) error {
	log.Printf("Starting pipeline execution: %s", p.name)
	
	input := make(map[string]interface{})
	
	for i, node := range p.nodes {
		log.Printf("Executing node %d/%d: %s", i+1, len(p.nodes), node.Name())
		
		// Validate node before execution
		if err := node.Validate(); err != nil {
			log.Printf("Node validation failed: %s - %v", node.Name(), err)
			return fmt.Errorf("node %s validation failed: %w", node.Name(), err)
		}
		
		output, err := node.Execute(ctx, input)
		if err != nil {
			log.Printf("Error in node %s: %v", node.Name(), err)
			return fmt.Errorf("node %s failed: %w", node.Name(), err)
		}
		
		// Merge output with input for next node
		for key, value := range output {
			input[key] = value
		}
		
		log.Printf("Node %s completed successfully", node.Name())
	}
	
	log.Printf("Pipeline %s execution completed successfully!", p.name)
	return nil
}

// GetNodeCount returns the number of nodes in the pipeline
func (p *Pipeline) GetNodeCount() int {
	return len(p.nodes)
}

// GetName returns the pipeline name
func (p *Pipeline) GetName() string {
	return p.name
} 