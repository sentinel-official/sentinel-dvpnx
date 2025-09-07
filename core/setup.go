package core

import (
	"context"
	"fmt"

	"github.com/sentinel-official/sentinel-go-sdk/core"
	"github.com/sentinel-official/sentinel-go-sdk/libs/geoip"
	"github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/sentinel-official/sentinel-go-sdk/openvpn"
	"github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/v2ray"
	"github.com/sentinel-official/sentinel-go-sdk/wireguard"

	"github.com/sentinel-official/sentinel-dvpnx/config"
	"github.com/sentinel-official/sentinel-dvpnx/database"
)

// SetupClient initializes the SDK client with the given configuration and assigns it to the context.
func (c *Context) SetupClient(cfg *config.Config) error {
	log.Info("Initializing blockchain client",
		"keyring.backend", cfg.Keyring.GetBackend(),
		"keyring.name", cfg.Keyring.GetName(),
		"rpc.addr", cfg.RPC.GetAddr(),
		"rpc.chain_id", cfg.RPC.GetChainID(),
		"tx.from_name", cfg.Tx.GetFromName(),
	)

	v, err := core.NewClientFromConfig(cfg.Config)
	if err != nil {
		return fmt.Errorf("creating client from config: %w", err)
	}

	// Seal the client.
	v.Seal()

	// Assign the initialized client to the context.
	c.WithClient(v)
	return nil
}

// SetupDatabase creates and configures the database, then assigns it to the context.
func (c *Context) SetupDatabase(_ *config.Config) error {
	log.Info("Initializing database", "file", c.DatabaseFile())

	db, err := database.NewDefault(c.DatabaseFile())
	if err != nil {
		return fmt.Errorf("initializing database %q: %w", c.DatabaseFile(), err)
	}

	// Assign the database instance to the context.
	c.WithDatabase(db)
	return nil
}

// SetupGeoIPClient initializes the GeoIP client and assigns it to the context.
func (c *Context) SetupGeoIPClient(_ *config.Config) error {
	log.Info("Initializing GeoIP client")
	v := geoip.NewDefaultClient()

	// Assign the GeoIP client to the context.
	c.WithGeoIPClient(v)
	return nil
}

// SetupAccAddr retrieves the account address for transactions and assigns it to the context.
func (c *Context) SetupAccAddr(cfg *config.Config) error {
	log.Info("Retrieving addr for key", "name", cfg.Tx.GetFromName())

	addr, err := c.Client().KeyAddr(cfg.Tx.GetFromName())
	if err != nil {
		return fmt.Errorf("getting addr for key %q: %w", cfg.Tx.GetFromName(), err)
	}

	log.Info("Querying account information", "addr", addr)

	acc, err := c.Client().Account(context.TODO(), addr)
	if err != nil {
		return fmt.Errorf("querying account %q: %w", addr, err)
	}
	if acc == nil {
		return fmt.Errorf("account %s does not exist", addr)
	}

	// Assign the account address to the context.
	c.WithAccAddr(addr)
	return nil
}

// SetupService determines the service type and configures it accordingly.
func (c *Context) SetupService(ctx context.Context, cfg *config.Config) error {
	var (
		s  types.ServerService         // Interface for the node service
		st = cfg.Node.GetServiceType() // Get the service type from config
	)

	log.Info("Initializing service", "type", st)

	// Initialize the appropriate server service based on the configured type
	switch st {
	case types.ServiceTypeV2Ray:
		s = v2ray.NewServer("v2ray", c.HomeDir(), cfg.Services[types.ServiceTypeV2Ray].(*v2ray.ServerConfig))
	case types.ServiceTypeWireGuard:
		s = wireguard.NewServer("wireguard", c.HomeDir(), cfg.Services[types.ServiceTypeWireGuard].(*wireguard.ServerConfig))
	case types.ServiceTypeOpenVPN:
		s = openvpn.NewServer("openvpn", c.HomeDir(), cfg.Services[types.ServiceTypeOpenVPN].(*openvpn.ServerConfig))
	default:
		return fmt.Errorf("unsupported service type %q", st)
	}

	log.Info("Checking service status")

	ok, err := s.IsRunning()
	if err != nil {
		return fmt.Errorf("checking service %q status: %w", st, err)
	}
	if ok {
		return fmt.Errorf("service %q is already running", st)
	}

	if err := s.Setup(ctx); err != nil {
		return err
	}

	// Assign the service to the context
	c.WithService(s)
	return nil
}

// Setup initializes all components of the node context.
func (c *Context) Setup(ctx context.Context, cfg *config.Config) error {
	// Assign configuration values to the context.
	c.WithAPIAddrs(cfg.Node.APIAddrs())
	c.WithAPIListenAddr(cfg.Node.APIListenAddr())
	c.WithGigabytePrices(cfg.Node.GetGigabytePrices())
	c.WithHourlyPrices(cfg.Node.GetHourlyPrices())
	c.WithMaxPeers(cfg.QoS.GetMaxPeers())
	c.WithMoniker(cfg.Node.GetMoniker())
	c.WithRemoteAddrs(cfg.Node.GetRemoteAddrs())
	c.WithRPCAddrs(cfg.RPC.GetAddrs())

	log.Info("Setting up blockchain client")
	if err := c.SetupClient(cfg); err != nil {
		return fmt.Errorf("setting up client: %w", err)
	}

	log.Info("Setting up database")
	if err := c.SetupDatabase(cfg); err != nil {
		return fmt.Errorf("setting up database: %w", err)
	}

	log.Info("Setting up GeoIP client")
	if err := c.SetupGeoIPClient(cfg); err != nil {
		return fmt.Errorf("setting up GeoIP client: %w", err)
	}

	log.Info("Setting up service")
	if err := c.SetupService(ctx, cfg); err != nil {
		return fmt.Errorf("setting up service: %w", err)
	}

	log.Info("Setting up account addr")
	if err := c.SetupAccAddr(cfg); err != nil {
		return fmt.Errorf("setting up account addr: %w", err)
	}

	return nil
}
