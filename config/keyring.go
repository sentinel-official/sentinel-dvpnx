package config

import (
	"errors"
)

type KeyringConfig struct {
	Name    string `mapstructure:"name"`    // Name is the name of the keyring.
	Backend string `mapstructure:"backend"` // Backend specifies the keyring backend to use.
}

// GetName returns the keyring name.
func (c *KeyringConfig) GetName() string {
	return c.Name
}

// GetBackend returns the keyring backend.
func (c *KeyringConfig) GetBackend() string {
	return c.Backend
}

// Validate validates the Keyring configuration.
func (c *KeyringConfig) Validate() error {
	if c.Name == "" {
		return errors.New("name cannot be empty")
	}

	// Check if the backend is one of the allowed values.
	validBackends := map[string]bool{
		"file":    true,
		"kwallet": true,
		"memory":  true,
		"os":      true,
		"pass":    true,
		"test":    true,
	}
	if !validBackends[c.Backend] {
		return errors.New("backend must be one of: file, kwallet, memory, os, pass, test")
	}

	return nil
}

func DefaultKeyringConfig() KeyringConfig {
	return KeyringConfig{
		Name:    "sentinel",
		Backend: "test",
	}
}
