package config

import (
	"errors"
	"fmt"

	"github.com/spf13/pflag"
)

const MaxHandshakeDNSPeers = 1 << 3 // Maximum number of peers for Handshake DNS.

// HandshakeDNSConfig represents the Handshake DNS configuration.
type HandshakeDNSConfig struct {
	Enable bool `mapstructure:"enable"` // Enable specifies if Handshake DNS is enabled.
	Peers  uint `mapstructure:"peers"`  // Peers specifies the number of DNS peers.
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
	// If Handshake DNS is not enabled, validation passes.
	if !c.Enable {
		return nil
	}

	// Ensure the number of peers is not zero.
	if c.Peers == 0 {
		return errors.New("peers cannot be zero")
	}

	// Ensure the number of peers does not exceed the maximum allowed value.
	if c.Peers > MaxHandshakeDNSPeers {
		return fmt.Errorf("peers cannot be greater than %d", MaxHandshakeDNSPeers)
	}

	return nil
}

// SetForFlags adds handshake-dns configuration flags to the specified FlagSet.
func (c *HandshakeDNSConfig) SetForFlags(f *pflag.FlagSet) {
	f.BoolVar(&c.Enable, "handshake-dns.enable", c.Enable, "enable or disable Handshake DNS")
	f.UintVar(&c.Peers, "handshake-dns.peers", c.Peers, "number of Handshake DNS peers")
}

// DefaultHandshakeDNSConfig returns a HandshakeDNSConfig instance with default values.
func DefaultHandshakeDNSConfig() *HandshakeDNSConfig {
	return &HandshakeDNSConfig{
		Enable: false,
		Peers:  MaxHandshakeDNSPeers,
	}
}
