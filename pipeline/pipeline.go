package pipeline

import (
	"context"
	"fmt"
	"log"
)

// Node represents a processing node in the pipeline
type Node interface {
	Execute(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error)
	Name() string
}

// Pipeline orchestrates the execution of nodes
type Pipeline struct {
	nodes []Node
}

// NewPipeline creates a new pipeline instance
func NewPipeline() *Pipeline {
	return &Pipeline{
		nodes: make([]Node, 0),
	}
}

// AddNode adds a node to the pipeline
func (p *Pipeline) AddNode(node Node) {
	p.nodes = append(p.nodes, node)
	log.Printf("Added node: %s", node.Name())
}

// Execute runs all nodes in the pipeline sequentially
func (p *Pipeline) Execute(ctx context.Context) error {
	log.Println("Starting pipeline execution...")
	
	input := make(map[string]interface{})
	
	for i, node := range p.nodes {
		log.Printf("Executing node %d/%d: %s", i+1, len(p.nodes), node.Name())
		
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
	
	log.Println("Pipeline execution completed successfully!")
	return nil
}

// GetNodeCount returns the number of nodes in the pipeline
func (p *Pipeline) GetNodeCount() int {
	return len(p.nodes)
} 