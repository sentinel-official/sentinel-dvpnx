package config

import (
	"errors"
	"fmt"

	"github.com/spf13/pflag"
)

const MaxQoSMaxPeers = 250 // Maximum allowed value for MaxPeers.

// QoSConfig represents the Quality of Service (QoS) configuration.
type QoSConfig struct {
	MaxPeers uint `mapstructure:"max_peers"` // MaxPeers specifies the maximum number of peers.
}

// WithMaxPeers sets the MaxPeers field and returns the updated QoSConfig.
func (c *QoSConfig) WithMaxPeers(maxPeers uint) *QoSConfig {
	c.MaxPeers = maxPeers

	return c
}

// GetMaxPeers returns the MaxPeers field.
func (c *QoSConfig) GetMaxPeers() uint {
	return c.MaxPeers
}

// Validate checks the validity of the QoS configuration.
func (c *QoSConfig) Validate() error {
	// Ensure MaxPeers is not zero.
	if c.MaxPeers == 0 {
		return errors.New("max_peers cannot be zero")
	}

	// Ensure MaxPeers does not exceed the maximum allowed value.
	if c.MaxPeers > MaxQoSMaxPeers {
		return fmt.Errorf("max_peers cannot be greater than %d", MaxQoSMaxPeers)
	}

	return nil
}

// SetForFlags adds qos configuration flags to the specified FlagSet.
func (c *QoSConfig) SetForFlags(f *pflag.FlagSet) {
	f.UintVar(&c.MaxPeers, "qos.max-peers", c.MaxPeers, "maximum number of peers for service")
}

// DefaultQoSConfig returns a QoSConfig instance with default values.
func DefaultQoSConfig() *QoSConfig {
	return &QoSConfig{
		MaxPeers: MaxQoSMaxPeers,
	}
}
