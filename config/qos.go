package config

import (
	"errors"
	"fmt"
)

const MaxQOSMaxPeers = (1 << 8) - (1 << 2)

type QOSConfig struct {
	MaxPeers uint `mapstructure:"max_peers"`
}

// WithMaxPeers sets the MaxPeers field and returns the updated QOSConfig.
func (c *QOSConfig) WithMaxPeers(maxPeers uint) *QOSConfig {
	c.MaxPeers = maxPeers
	return c
}

// GetMaxPeers returns the MaxPeers field.
func (c *QOSConfig) GetMaxPeers() uint {
	return c.MaxPeers
}

// Validate checks the validity of the QOS configuration.
func (c *QOSConfig) Validate() error {
	if c.MaxPeers == 0 {
		return errors.New("max_peers cannot be zero")
	}
	if c.MaxPeers > MaxQOSMaxPeers {
		return fmt.Errorf("max_peers cannot be greater than %d", MaxQOSMaxPeers)
	}

	return nil
}

// DefaultQOSConfig returns a QOSConfig instance with default values.
func DefaultQOSConfig() QOSConfig {
	return QOSConfig{
		MaxPeers: MaxQOSMaxPeers,
	}
}
