package context

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/sentinel-official/sentinel-go-sdk/client"
	"github.com/sentinel-official/sentinel-go-sdk/libs/geoip"
	"github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/v2ray"
	"github.com/sentinel-official/sentinel-go-sdk/wireguard"

	"github.com/sentinel-official/dvpn-node/config"
	"github.com/sentinel-official/dvpn-node/database"
)

// SetupClient initializes the SDK client with the given configuration and assigns it to the context.
func (c *Context) SetupClient(cfg *config.Config) error {
	log.Info("Setting up client")

	// Create a codec for encoding/decoding protocol buffer messages.
	protoCodec := types.NewProtoCodec()
	txConfig := tx.NewTxConfig(protoCodec, tx.DefaultSignModes)

	// Create a keyring to manage cryptographic keys.
	kr, err := keyring.New(cfg.Keyring.GetName(), cfg.Keyring.GetBackend(), c.HomeDir(), c.Input(), protoCodec)
	if err != nil {
		return fmt.Errorf("failed to create keyring: %w", err)
	}

	// Initialize the client with the provided configurations.
	v := client.New().
		WithChainID(cfg.Tx.GetChainID()).
		WithKeyring(kr).
		WithProtoCodec(protoCodec).
		WithQueryProve(cfg.Query.GetProve()).
		WithQueryRetries(cfg.Query.GetRetries()).
		WithQueryRetryDelay(cfg.Query.GetRetryDelay()).
		WithRPCAddr(c.RPCAddr()).
		WithRPCTimeout(cfg.RPC.GetTimeout()).
		WithTxConfig(txConfig).
		WithTxFeeGranterAddr(cfg.Tx.GetFeeGranterAddr()).
		WithTxFees(nil).
		WithTxFromName(cfg.Tx.GetFromName()).
		WithTxGas(cfg.Tx.GetGas()).
		WithTxGasAdjustment(cfg.Tx.GetGasAdjustment()).
		WithTxGasPrices(cfg.Tx.GetGasPrices()).
		WithTxMemo("").
		WithTxSimulateAndExecute(cfg.Tx.GetSimulateAndExecute()).
		WithTxTimeoutHeight(0)

	// Assign the initialized client to the context.
	c.WithClient(v)
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

// setupV2RayService configures the V2Ray service and assigns it to the context.
func (c *Context) setupV2RayService(cfg *v2ray.ServerConfig) error {
	// Initialize the peer manager and V2Ray service.
	pm := v2ray.NewPeerManager()
	service := v2ray.NewServer().
		WithHomeDir(c.HomeDir()).
		WithName("v2ray").
		WithPeerManager(pm)

	ok, err := service.IsUp(context.TODO())
	if err != nil {
		return fmt.Errorf("failed to check if v2ray service is up: %w", err)
	}
	if ok {
		return fmt.Errorf("v2ray service is already up")
	}

	// Perform pre-start setup for the V2Ray service.
	if err := service.PreUp(cfg); err != nil {
		return fmt.Errorf("failed to run v2ray pre-up task: %w", err)
	}

	// Assign the service to the context.
	c.WithService(service)
	return nil
}

// setupWireGuardService configures the WireGuard service and assigns it to the context.
func (c *Context) setupWireGuardService(cfg *wireguard.ServerConfig) error {
	pools, err := cfg.IPPools()
	if err != nil {
		return fmt.Errorf("failed to get ip pools: %w", err)
	}

	// Initialize the peer manager and WireGuard service.
	pm := wireguard.NewPeerManager(pools...)
	service := wireguard.NewServer().
		WithHomeDir(c.HomeDir()).
		WithName(cfg.InInterface).
		WithPeerManager(pm)

	ok, err := service.IsUp(context.TODO())
	if err != nil {
		return fmt.Errorf("failed to check if wireguard service is up: %w", err)
	}
	if ok {
		return fmt.Errorf("wireguard service is already up")
	}

	// Perform pre-start setup tasks for the WireGuard service.
	if err := service.PreUp(cfg); err != nil {
		return fmt.Errorf("failed to run wireguard pre-up task: %w", err)
	}

	// Assign the service to the context.
	c.WithService(service)
	return nil
}

// SetupService determines the service type and configures it accordingly.
func (c *Context) SetupService(cfg *config.Config) error {
	log.Info("Setting up service")

	// Determine the type of service to set up.
	t := cfg.Node.GetType()
	switch t {
	case types.ServiceTypeV2Ray:
		// Setup the V2Ray service.
		if err := c.setupV2RayService(&cfg.V2Ray); err != nil {
			return fmt.Errorf("failed to setup v2ray service: %w", err)
		}
	case types.ServiceTypeWireGuard:
		// Setup the WireGuard service.
		if err := c.setupWireGuardService(&cfg.WireGuard); err != nil {
			return fmt.Errorf("failed to setup wireguard service: %w", err)
		}
	default:
		// Return an error for unsupported service types.
		return fmt.Errorf("invalid service type %s", t)
	}
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
	c.WithRemoteAddrs(cfg.Node.RemoteAddrs)
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
