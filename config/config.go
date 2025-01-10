package config

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/v2ray"
	"github.com/sentinel-official/sentinel-go-sdk/wireguard"
)

//go:embed *.tmpl
var fs embed.FS

type Config struct {
	HandshakeDNS HandshakeDNSConfig     `mapstructure:"handshake_dns"`
	Keyring      KeyringConfig          `mapstructure:"keyring"`
	Log          LogConfig              `mapstructure:"log"`
	Node         NodeConfig             `mapstructure:"node"`
	QOS          QOSConfig              `mapstructure:"qos"`
	Query        QueryConfig            `mapstructure:"query"`
	RPC          RPCConfig              `mapstructure:"rpc"`
	Tx           TxConfig               `mapstructure:"tx"`
	V2Ray        v2ray.ServerConfig     `mapstructure:"v2ray"`
	WireGuard    wireguard.ServerConfig `mapstructure:"wireguard"`
}

// Validate validates the entire configuration.
func (c *Config) Validate() error {
	if err := c.HandshakeDNS.Validate(); err != nil {
		return fmt.Errorf("invalid handshake_dns: %w", err)
	}
	if err := c.Keyring.Validate(); err != nil {
		return fmt.Errorf("invalid keyring: %w", err)
	}
	if err := c.Log.Validate(); err != nil {
		return fmt.Errorf("invalid log: %w", err)
	}
	if err := c.Node.Validate(); err != nil {
		return fmt.Errorf("invalid node: %w", err)
	}
	if err := c.QOS.Validate(); err != nil {
		return fmt.Errorf("invalid qos: %w", err)
	}
	if err := c.Query.Validate(); err != nil {
		return fmt.Errorf("invalid query: %w", err)
	}
	if err := c.RPC.Validate(); err != nil {
		return fmt.Errorf("invalid rpc: %w", err)
	}
	if err := c.Tx.Validate(); err != nil {
		return fmt.Errorf("invalid tx: %w", err)
	}

	if c.Node.GetType() == types.ServiceTypeV2Ray {
		if err := c.V2Ray.Validate(); err != nil {
			return fmt.Errorf("invalid v2ray: %w", err)
		}
	}
	if c.Node.GetType() == types.ServiceTypeWireGuard {
		if err := c.WireGuard.Validate(); err != nil {
			return fmt.Errorf("invalid wireguard: %w", err)
		}
	}

	return nil
}

func DefaultConfig() Config {
	return Config{
		HandshakeDNS: DefaultHandshakeDNSConfig(),
		Keyring:      DefaultKeyringConfig(),
		Log:          DefaultLogConfig(),
		Node:         DefaultNodeConfig(),
		QOS:          DefaultQOSConfig(),
		Query:        DefaultQueryConfig(),
		RPC:          DefaultRPCConfig(),
		Tx:           DefaultTxConfig(),
		V2Ray:        v2ray.DefaultServerConfig(),
		WireGuard:    wireguard.DefaultServerConfig(),
	}
}

func (c *Config) WriteToFile(name string) error {
	text, err := fs.ReadFile("config.toml.tmpl")
	if err != nil {
		return err
	}

	tmpl, err := template.New("config").Parse(string(text))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, c); err != nil {
		return err
	}

	return os.WriteFile(name, buf.Bytes(), 0644)
}
