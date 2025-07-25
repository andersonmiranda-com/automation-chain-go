package config

import (
	"test-chain-go-cursor/config/credentials"
)

// CredentialsManager is now an alias for the new Manager
type CredentialsManager = credentials.Manager

// NewCredentialsManager creates a new credentials manager
func NewCredentialsManager() *CredentialsManager {
	return credentials.NewManager()
} 