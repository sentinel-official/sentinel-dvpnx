package config

import (
	"errors"
	"fmt"
)

const MaxHandshakeDNSPeers = 1 << 3

// HandshakeDNSConfig represents the Handshake DNS configuration.
type HandshakeDNSConfig struct {
	Enable bool `mapstructure:"enable"`
	Peers  uint `mapstructure:"peers"`
}

// WithEnable sets the Enable field and returns the updated HandshakeDNSConfig.
func (c *HandshakeDNSConfig) WithEnable(enable bool) *HandshakeDNSConfig {
	c.Enable = enable
	return c
}

// WithPeers sets the Peers field and returns the updated HandshakeDNSConfig.
func (c *HandshakeDNSConfig) WithPeers(peers uint) *HandshakeDNSConfig {
	c.Peers = peers
	return c
}

// GetEnable returns the Enable field.
func (c *HandshakeDNSConfig) GetEnable() bool {
	return c.Enable
}

// GetPeers returns the Peers field.
func (c *HandshakeDNSConfig) GetPeers() uint {
	return c.Peers
}

// Validate checks the validity of the HandshakeDNSConfig configuration.
func (c *HandshakeDNSConfig) Validate() error {
	if !c.Enable {
		return nil
	}
	if c.Peers == 0 {
		return errors.New("peers cannot be zero")
	}
	if c.Peers > MaxHandshakeDNSPeers {
		return fmt.Errorf("peers cannot be greater than %d", MaxHandshakeDNSPeers)
	}

	return nil
}

// DefaultHandshakeDNSConfig returns a HandshakeDNSConfig instance with default values.
func DefaultHandshakeDNSConfig() HandshakeDNSConfig {
	return HandshakeDNSConfig{
		Enable: false,
		Peers:  MaxHandshakeDNSPeers,
	}
}
