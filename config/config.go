package config

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/sentinel-official/sentinel-go-sdk/config"
	"github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/v2ray"
	"github.com/sentinel-official/sentinel-go-sdk/wireguard"
	"github.com/spf13/pflag"
)

// Embed the template files for configuration.
//
//go:embed *.tmpl
var fs embed.FS

// Config represents the overall configuration structure.
type Config struct {
	*config.Config `mapstructure:",squash"`
	HandshakeDNS   *HandshakeDNSConfig     `mapstructure:"handshake_dns"` // HandshakeDNS contains Handshake DNS configuration.
	Node           *NodeConfig             `mapstructure:"node"`          // Node contains node-specific configuration.
	QOS            *QOSConfig              `mapstructure:"qos"`           // QOS contains Quality of Service configuration.
	V2Ray          *v2ray.ServerConfig     `mapstructure:"v2ray"`         // V2Ray contains V2Ray server configuration.
	WireGuard      *wireguard.ServerConfig `mapstructure:"wireguard"`     // WireGuard contains WireGuard server configuration.
}

// Validate validates the entire configuration.
func (c *Config) Validate() error {
	if err := c.Config.Validate(); err != nil {
		return err
	}
	if err := c.HandshakeDNS.Validate(); err != nil {
		return fmt.Errorf("invalid handshake_dns: %w", err)
	}
	if err := c.Node.Validate(); err != nil {
		return fmt.Errorf("invalid node: %w", err)
	}
	if err := c.QOS.Validate(); err != nil {
		return fmt.Errorf("invalid qos: %w", err)
	}

	// Validate V2Ray configuration if the node type is V2Ray.
	if c.Node.GetType() == types.ServiceTypeV2Ray {
		if err := c.V2Ray.Validate(); err != nil {
			return fmt.Errorf("invalid v2ray: %w", err)
		}
	}

	// Validate WireGuard configuration if the node type is WireGuard.
	if c.Node.GetType() == types.ServiceTypeWireGuard {
		if err := c.WireGuard.Validate(); err != nil {
			return fmt.Errorf("invalid wireguard: %w", err)
		}
	}

	return nil
}

// SetForFlags adds configuration flags to the specified FlagSet.
func (c *Config) SetForFlags(f *pflag.FlagSet) {
	c.Config.SetForFlags(f)
	c.HandshakeDNS.SetForFlags(f)
	c.Node.SetForFlags(f)
	c.QOS.SetForFlags(f)
	c.V2Ray.SetForFlags(f)
	c.WireGuard.SetForFlags(f)
}

// DefaultConfig returns a configuration instance with default values.
func DefaultConfig() *Config {
	return &Config{
		Config:       config.DefaultConfig(),
		HandshakeDNS: DefaultHandshakeDNSConfig(),
		Node:         DefaultNodeConfig(),
		QOS:          DefaultQOSConfig(),
		V2Ray:        v2ray.DefaultServerConfig(),
		WireGuard:    wireguard.DefaultServerConfig(),
	}
}

// WriteToFile writes the configuration to a file using a template.
func (c *Config) WriteToFile(name string) error {
	// Read the configuration template file.
	text, err := fs.ReadFile("config.toml.tmpl")
	if err != nil {
		return err
	}

	// Parse the template file.
	tmpl, err := template.New("config").Parse(string(text))
	if err != nil {
		return err
	}

	// Execute the template and write the output to a buffer.
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, c); err != nil {
		return err
	}

	// Write the buffer contents to the specified file.
	return os.WriteFile(name, buf.Bytes(), 0644)
}
