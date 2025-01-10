package config

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type RPCConfig struct {
	Addrs   []string `mapstructure:"addrs"`
	Timeout string   `mapstructure:"timeout"`
}

// GetAddrs returns the address of the RPC server.
func (c *RPCConfig) GetAddrs() []string {
	return c.Addrs
}

// GetTimeout returns the maximum duration for the query.
func (c *RPCConfig) GetTimeout() time.Duration {
	v, err := time.ParseDuration(c.Timeout)
	if err != nil {
		panic(err)
	}

	return v
}

// Validate validates the RPC configuration.
func (c *RPCConfig) Validate() error {
	// Validate that Addrs is not empty.
	if len(c.Addrs) == 0 {
		return errors.New("addrs cannot be empty")
	}

	// Validate each address in Addrs.
	for _, addr := range c.Addrs {
		if err := validateURL(addr); err != nil {
			return fmt.Errorf("invalid addr: %w", err)
		}
	}

	// Validate that Timeout is a valid time.Duration.
	if _, err := time.ParseDuration(c.Timeout); err != nil {
		return fmt.Errorf("invalid timeout: %w", err)
	}

	return nil
}

func DefaultRPCConfig() RPCConfig {
	return RPCConfig{
		Addrs: []string{
			"https://rpc.sentinel.co:443",
		},
		Timeout: "5s",
	}
}

func validateURL(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return fmt.Errorf("invalid url: %w", err)
	}
	if u.Scheme == "" {
		return errors.New("url must have a valid scheme")
	}
	if u.Host == "" {
		return errors.New("url must have a valid host")
	}

	// Check if the port is a valid number
	port, err := strconv.Atoi(u.Port())
	if err != nil {
		return fmt.Errorf("invalid port: %w", err)
	}
	if port < 1 || port > 65535 {
		return errors.New("url must have a valid port")
	}

	return nil
}
