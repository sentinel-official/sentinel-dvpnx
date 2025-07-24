package context

import (
	"context"
	"fmt"
	"path/filepath"

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
	log.Info("Setting up client")

	// Initialize the base client with the provided configurations.
	cc := core.NewClient().
		WithQueryProve(cfg.Query.GetProve()).
		WithQueryRetryAttempts(cfg.Query.GetRetryAttempts()).
		WithQueryRetryDelay(cfg.Query.GetRetryDelay()).
		WithRPCAddr(c.RPCAddr()).
		WithRPCChainID(cfg.RPC.GetChainID()).
		WithRPCTimeout(cfg.RPC.GetTimeout()).
		WithTxBroadcastRetryAttempts(cfg.Tx.GetBroadcastRetryAttempts()).
		WithTxBroadcastRetryDelay(cfg.Tx.GetBroadcastRetryDelay()).
		WithTxFeeGranterAddr(cfg.Tx.GetFeeGranterAddr()).
		WithTxFees(nil).
		WithTxFromName(cfg.Tx.GetFromName()).
		WithTxGasAdjustment(cfg.Tx.GetGasAdjustment()).
		WithTxGas(cfg.Tx.GetGas()).
		WithTxGasPrices(cfg.Tx.GetGasPrices()).
		WithTxMemo("").
		WithTxQueryRetryAttempts(cfg.Tx.GetQueryRetryAttempts()).
		WithTxQueryRetryDelay(cfg.Tx.GetQueryRetryDelay()).
		WithTxSimulateAndExecute(cfg.Tx.GetSimulateAndExecute()).
		WithTxTimeoutHeight(0)

	// Setup the keyring for the base client
	if err := cc.SetupKeyring(cfg.Keyring); err != nil {
		return fmt.Errorf("failed to setup keyring: %w", err)
	}

	// Assign the initialized client to the context.
	c.WithClient(cc)
	return nil
}

// SetupDatabase creates and configures the database, then assigns it to the context.
func (c *Context) SetupDatabase(_ *config.Config) error {
	log.Info("Setting up database")

	// Construct the database path within the home directory.
	dbPath := filepath.Join(c.HomeDir(), "data.db")

	// Initialize the database connection.
	db, err := database.NewDefault(dbPath)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	// Assign the database instance to the context.
	c.WithDatabase(db)
	return nil
}

// SetupGeoIPClient initializes the GeoIP client and assigns it to the context.
func (c *Context) SetupGeoIPClient(_ *config.Config) error {
	log.Info("Setting up geoip client")

	// Create a default GeoIP client instance.
	v := geoip.NewDefaultClient()

	// Assign the GeoIP client to the context.
	c.WithGeoIPClient(v)
	return nil
}

// SetupAccAddr retrieves the account address for transactions and assigns it to the context.
func (c *Context) SetupAccAddr(cfg *config.Config) error {
	log.Info("Setting up account address")

	// Retrieve the key associated with the configured account name.
	key, err := c.Client().Key(cfg.Tx.GetFromName())
	if err != nil {
		return fmt.Errorf("failed to retrieve key: %w", err)
	}
	if key == nil {
		return fmt.Errorf("key %s does not exist", cfg.Tx.GetFromName())
	}

	// Extract the address from the key.
	addr, err := key.GetAddress()
	if err != nil {
		return fmt.Errorf("failed to retrieve address from key: %w", err)
	}

	// Query the account to ensure it exists and is valid.
	acc, err := c.Client().Account(context.TODO(), addr)
	if err != nil {
		return fmt.Errorf("failed to query account: %w", err)
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
	log.Info("Setting up service")

	var (
		service     types.ServerService         // The service instance to configure
		serviceType = cfg.Node.GetServiceType() // Type of the service from configuration
	)

	// Determine the service type from the node configuration
	switch cfg.Node.GetServiceType() {
	case types.ServiceTypeV2Ray:
		service = v2ray.NewServer(c.HomeDir())
	case types.ServiceTypeWireGuard:
		service = wireguard.NewServer(c.HomeDir())
	case types.ServiceTypeOpenVPN:
		service = openvpn.NewServer(c.HomeDir())
	default:
		return fmt.Errorf("invalid service type %s", serviceType)
	}

	// Check if the service is already running
	ok, err := service.IsUp(context.TODO())
	if err != nil {
		return fmt.Errorf("failed to check if service is up: %w", err)
	}
	if ok {
		return fmt.Errorf("service is already up")
	}

	// Perform pre-start setup tasks for the service
	if err := service.PreUp(); err != nil {
		return fmt.Errorf("failed to run pre-up task: %w", err)
	}

	// Assign the service to the context
	c.WithService(service)
	return nil
}

// Setup initializes all components of the node context.
func (c *Context) Setup(cfg *config.Config) error {
	log.Info("Setting up node context...")

	// Assign configuration values to the context.
	c.WithAPIAddrs(cfg.Node.APIAddrs())
	c.WithGigabytePrices(cfg.Node.GetGigabytePrices())
	c.WithHourlyPrices(cfg.Node.GetHourlyPrices())
	c.WithMoniker(cfg.Node.GetMoniker())
	c.WithRemoteAddrs(cfg.Node.GetRemoteAddrs())
	c.WithRPCAddrs(cfg.RPC.GetAddrs())

	// Set up the client for blockchain communication.
	if err := c.SetupClient(cfg); err != nil {
		return fmt.Errorf("failed to setup client: %w", err)
	}

	// Set up the local database for storing data.
	if err := c.SetupDatabase(cfg); err != nil {
		return fmt.Errorf("failed to setup database: %w", err)
	}

	// Set up the GeoIP client for geographical location resolution.
	if err := c.SetupGeoIPClient(cfg); err != nil {
		return fmt.Errorf("failed to setup geoip client: %w", err)
	}

	// Set up the appropriate service (e.g., V2Ray or WireGuard).
	if err := c.SetupService(cfg); err != nil {
		return fmt.Errorf("failed to setup service: %w", err)
	}

	// Retrieve and configure the account address for transactions.
	if err := c.SetupAccAddr(cfg); err != nil {
		return fmt.Errorf("failed to setup account address: %w", err)
	}

	return nil
}
