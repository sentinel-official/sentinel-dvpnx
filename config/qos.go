package config

import (
	"errors"
	"fmt"

	"github.com/spf13/pflag"
)

const MaxQOSMaxPeers = 250 // Maximum allowed value for MaxPeers.

// QOSConfig represents the Quality of Service (QoS) configuration.
type QOSConfig struct {
	MaxPeers uint `mapstructure:"max_peers"` // MaxPeers specifies the maximum number of peers.
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
	// Ensure MaxPeers is not zero.
	if c.MaxPeers == 0 {
		return errors.New("max_peers cannot be zero")
	}

	// Ensure MaxPeers does not exceed the maximum allowed value.
	if c.MaxPeers > MaxQOSMaxPeers {
		return fmt.Errorf("max_peers cannot be greater than %d", MaxQOSMaxPeers)
	}

	return nil
}

// SetForFlags adds qos configuration flags to the specified FlagSet.
func (c *QOSConfig) SetForFlags(f *pflag.FlagSet) {
	f.UintVar(&c.MaxPeers, "qos.max-peers", c.MaxPeers, "maximum number of peers for QoS")
}

// DefaultQOSConfig returns a QOSConfig instance with default values.
func DefaultQOSConfig() *QOSConfig {
	return &QOSConfig{
		MaxPeers: MaxQOSMaxPeers,
	}
}
