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

	cc, err := core.NewClientFromConfig(cfg.Config)
	if err != nil {
		return fmt.Errorf("creating client from config: %w", err)
	}

	// Seal the client.
	cc.Seal()

	// Assign the initialized client to the context.
	c.WithClient(cc)
	return nil
}

// SetupDatabase creates and configures the database, then assigns it to the context.
func (c *Context) SetupDatabase(_ *config.Config) error {
	log.Info("Initializing database connection", "file", c.DatabaseFile())

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
func (c *Context) SetupService(cfg *config.Config) error {
	var (
		service     types.ServerService         // The service instance to configure
		serviceType = cfg.Node.GetServiceType() // Type of the service from configuration
	)

	log.Info("Initializing service", "type", serviceType)

	switch cfg.Node.GetServiceType() {
	case types.ServiceTypeV2Ray:
		service = v2ray.NewServer(c.HomeDir())
	case types.ServiceTypeWireGuard:
		service = wireguard.NewServer(c.HomeDir())
	case types.ServiceTypeOpenVPN:
		service = openvpn.NewServer(c.HomeDir())
	default:
		return fmt.Errorf("unsupported service type %q", serviceType)
	}

	log.Info("Checking service status")

	ok, err := service.IsUp()
	if err != nil {
		return fmt.Errorf("checking service status: %w", err)
	}
	if ok {
		return fmt.Errorf("service is already up")
	}

	// Assign the service to the context
	c.WithService(service)
	return nil
}

// Setup initializes all components of the node context.
func (c *Context) Setup(cfg *config.Config) error {
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
	if err := c.SetupService(cfg); err != nil {
		return fmt.Errorf("setting up service: %w", err)
	}

	log.Info("Setting up account addr")
	if err := c.SetupAccAddr(cfg); err != nil {
		return fmt.Errorf("setting up account addr: %w", err)
	}

	return nil
}
