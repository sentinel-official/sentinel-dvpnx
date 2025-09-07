package config

import (
	"embed"
	"fmt"
	"os"

	"github.com/sentinel-official/sentinel-go-sdk/config"
	"github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/utils"
	"github.com/spf13/pflag"
)

// Embed the template files for configuration.
//
//go:embed *.tmpl
var fs embed.FS

// Config represents the overall configuration structure.
type Config struct {
	*config.Config `mapstructure:",squash"`

	HandshakeDNS *HandshakeDNSConfig `mapstructure:"handshake_dns"` // HandshakeDNS contains Handshake DNS configuration.
	Node         *NodeConfig         `mapstructure:"node"`          // Node contains node-specific configuration.
	QoS          *QoSConfig          `mapstructure:"qos"`           // QoS contains Quality of Service configuration.

	Services map[types.ServiceType]types.ServiceConfig `mapstructure:"-"`
}

// Validate validates the entire configuration.
func (c *Config) Validate() error {
	if err := c.Config.Validate(); err != nil {
		return fmt.Errorf("validating base config: %w", err)
	}

	if err := c.HandshakeDNS.Validate(); err != nil {
		return fmt.Errorf("validating handshake_dns config: %w", err)
	}

	if err := c.Node.Validate(); err != nil {
		return fmt.Errorf("validating node config: %w", err)
	}

	if err := c.QoS.Validate(); err != nil {
		return fmt.Errorf("validating QoS config: %w", err)
	}

	return nil
}

// SetForFlags adds configuration flags to the specified FlagSet.
func (c *Config) SetForFlags(f *pflag.FlagSet) {
	c.Config.SetForFlags(f)
	c.HandshakeDNS.SetForFlags(f)
	c.Node.SetForFlags(f)
	c.QoS.SetForFlags(f)
}

// DefaultConfig returns a configuration instance with default values.
func DefaultConfig() *Config {
	return &Config{
		Config:       config.DefaultConfig(),
		HandshakeDNS: DefaultHandshakeDNSConfig(),
		Node:         DefaultNodeConfig(),
		QoS:          DefaultQoSConfig(),
	}
}

// WriteAppConfig generates the application-level configuration file using the main config template.
func (c *Config) WriteAppConfig(file string) error {
	// Load the application template from the embedded filesystem.
	text, err := fs.ReadFile("config.toml.tmpl")
	if err != nil {
		return fmt.Errorf("reading config template: %w", err)
	}

	// Render the template with Config data and write the result to the specified file.
	if err := utils.ExecTemplateToFile(string(text), c, file); err != nil {
		return fmt.Errorf("writing rendered config file %q: %w", file, err)
	}

	// Restrict file permissions to owner read/write only.
	if err := os.Chmod(file, 0600); err != nil {
		return fmt.Errorf("setting file permissions: %w", err)
	}

	return nil
}
