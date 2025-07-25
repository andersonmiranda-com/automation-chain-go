package base

import (
	"context"
)

// Node represents a processing node in the pipeline
type Node interface {
	Execute(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error)
	Name() string
	Config() NodeConfig
	Validate() error
}

// NodeConfig holds configuration for a node
type NodeConfig struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Name       string                 `json:"name"`
	Parameters map[string]interface{} `json:"parameters"`
}

// NodeDefinition represents a node in pipeline configuration
type NodeDefinition struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Name        string                 `json:"name"`
	Credentials string                 `json:"credentials,omitempty"`
	Config      map[string]interface{} `json:"config"`
} 