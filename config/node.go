package config

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net"
	"strings"
	"time"

	"github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/utils"
	"github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/spf13/pflag"
)

const MaxRemoteAddrLen = (1 << 6) - 1 // Maximum allowable length for a remote address.

type NodeConfig struct {
	APIPort                                string   `mapstructure:"api_port"`                                    // APIPort is the port for API access.
	GigabytePrices                         string   `mapstructure:"gigabyte_prices"`                             // GigabytePrices is the pricing information for gigabytes.
	HourlyPrices                           string   `mapstructure:"hourly_prices"`                               // HourlyPrices is the pricing information for hourly usage.
	IntervalBestRPCAddr                    string   `mapstructure:"interval_best_rpc_addr"`                      // IntervalBestRPCAddr is the duration between checking the best RPC address.
	IntervalGeoIPLocation                  string   `mapstructure:"interval_geo_ip_location"`                    // IntervalGeoIPLocation is the duration between checking the GeoIP location.
	IntervalSessionUsageSyncWithBlockchain string   `mapstructure:"interval_session_usage_sync_with_blockchain"` // IntervalSessionUsageSyncWithBlockchain is the duration between syncing session usage with the blockchain.
	IntervalSessionUsageSyncWithDatabase   string   `mapstructure:"interval_session_usage_sync_with_database"`   // IntervalSessionUsageSyncWithDatabase is the duration between syncing session usage with the database.
	IntervalSessionUsageValidate           string   `mapstructure:"interval_session_usage_validate"`             // IntervalSessionUsageValidate is the duration between validating session usage.
	IntervalSessionValidate                string   `mapstructure:"interval_session_validate"`                   // IntervalSessionValidate is the duration between validating sessions.
	IntervalSpeedtest                      string   `mapstructure:"interval_speedtest"`                          // IntervalSpeedtest is the duration between performing speed tests.
	IntervalStatusUpdate                   string   `mapstructure:"interval_status_update"`                      // IntervalStatusUpdate is the duration between updating the status of the node.
	Moniker                                string   `mapstructure:"moniker"`                                     // Moniker is the name or identifier for the node.
	RemoteAddrs                            []string `mapstructure:"remote_addrs"`                                // RemoteAddrs is a list of remote addresses for operations.
	ServiceType                            string   `mapstructure:"service_type"`                                // ServiceType is the type of the service.
	TLSCertPath                            string   `mapstructure:"tls_cert_path"`                               // TLSCertPath is the file path to the TLS certificate.
	TLSKeyPath                             string   `mapstructure:"tls_key_path"`                                // TLSKeyPath is the file path to the TLS private key.
}

// APIAddrs generates the API addresses for the node.
func (c *NodeConfig) APIAddrs() []string {
	addrs := make([]string, len(c.RemoteAddrs))
	port := c.GetAPIPort().OutFrom

	for i, addr := range c.RemoteAddrs {
		addrs[i] = fmt.Sprintf("https://%s:%d", addr, port)
	}

	return addrs
}

// APIListenAddr returns the API listen address.
func (c *NodeConfig) APIListenAddr() string {
	return fmt.Sprintf("0.0.0.0:%d", c.APIListenPort())
}

// APIListenPort returns the API listen port.
func (c *NodeConfig) APIListenPort() uint16 {
	return c.GetAPIPort().InFrom
}

// GetAPIPort returns the APIPort field.
func (c *NodeConfig) GetAPIPort() types.Port {
	v, err := types.NewPortFromString(c.APIPort)
	if err != nil {
		panic(err)
	}

	return v
}

// GetGigabytePrices returns the GigabytePrices field.
func (c *NodeConfig) GetGigabytePrices() v1.Prices {
	v, err := v1.NewPricesFromString(c.GigabytePrices)
	if err != nil {
		panic(err)
	}

	return v
}

// GetHourlyPrices returns the HourlyPrices field.
func (c *NodeConfig) GetHourlyPrices() v1.Prices {
	v, err := v1.NewPricesFromString(c.HourlyPrices)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalBestRPCAddr returns the IntervalBestRPCAddr field.
func (c *NodeConfig) GetIntervalBestRPCAddr() time.Duration {
	v, err := time.ParseDuration(c.IntervalBestRPCAddr)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalGeoIPLocation returns the IntervalGeoIPLocation field.
func (c *NodeConfig) GetIntervalGeoIPLocation() time.Duration {
	v, err := time.ParseDuration(c.IntervalGeoIPLocation)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSessionUsageSyncWithBlockchain returns the IntervalSessionUsageSyncWithBlockchain field.
func (c *NodeConfig) GetIntervalSessionUsageSyncWithBlockchain() time.Duration {
	v, err := time.ParseDuration(c.IntervalSessionUsageSyncWithBlockchain)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSessionUsageSyncWithDatabase returns the IntervalSessionUsageSyncWithDatabase field.
func (c *NodeConfig) GetIntervalSessionUsageSyncWithDatabase() time.Duration {
	v, err := time.ParseDuration(c.IntervalSessionUsageSyncWithDatabase)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSessionUsageValidate returns the IntervalSessionUsageValidate field.
func (c *NodeConfig) GetIntervalSessionUsageValidate() time.Duration {
	v, err := time.ParseDuration(c.IntervalSessionUsageValidate)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSessionValidate returns the IntervalSessionValidate field.
func (c *NodeConfig) GetIntervalSessionValidate() time.Duration {
	v, err := time.ParseDuration(c.IntervalSessionValidate)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSpeedtest returns the IntervalSpeedtest field.
func (c *NodeConfig) GetIntervalSpeedtest() time.Duration {
	v, err := time.ParseDuration(c.IntervalSpeedtest)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalStatusUpdate returns the IntervalStatusUpdate field.
func (c *NodeConfig) GetIntervalStatusUpdate() time.Duration {
	v, err := time.ParseDuration(c.IntervalStatusUpdate)
	if err != nil {
		panic(err)
	}

	return v
}

// GetMoniker returns the Moniker field.
func (c *NodeConfig) GetMoniker() string {
	return c.Moniker
}

// GetRemoteAddrs returns the RemoteAddrs field.
func (c *NodeConfig) GetRemoteAddrs() []string {
	return c.RemoteAddrs
}

// GetServiceType returns the ServiceType field.
func (c *NodeConfig) GetServiceType() types.ServiceType {
	return types.ServiceTypeFromString(c.ServiceType)
}

// GetTLSCertPath returns the TLSCertPath field.
func (c *NodeConfig) GetTLSCertPath() string {
	return c.TLSCertPath
}

// GetTLSKeyPath returns the TLSKeyPath field.
func (c *NodeConfig) GetTLSKeyPath() string {
	return c.TLSKeyPath
}

// Validate validates the node configuration.
func (c *NodeConfig) Validate() error {
	// Ensure the API port is not empty and validate it.
	if c.APIPort == "" {
		return errors.New("api_port cannot be empty")
	}
	if _, err := types.NewPortFromString(c.APIPort); err != nil {
		return fmt.Errorf("invalid api_port: %w", err)
	}

	// Validate the GigabytePrices field.
	if _, err := v1.NewPricesFromString(c.GigabytePrices); err != nil {
		return fmt.Errorf("invalid gigabyte_prices: %w", err)
	}

	// Validate the HourlyPrices field.
	if _, err := v1.NewPricesFromString(c.HourlyPrices); err != nil {
		return fmt.Errorf("invalid hourly_prices: %w", err)
	}

	// Validate interval fields.
	if _, err := time.ParseDuration(c.IntervalBestRPCAddr); err != nil {
		return fmt.Errorf("invalid interval_best_rpc_addr: %w", err)
	}
	if _, err := time.ParseDuration(c.IntervalGeoIPLocation); err != nil {
		return fmt.Errorf("invalid interval_geoip_location: %w", err)
	}
	if _, err := time.ParseDuration(c.IntervalSessionUsageSyncWithBlockchain); err != nil {
		return fmt.Errorf("invalid interval_session_usage_sync_with_blockchain: %w", err)
	}
	if _, err := time.ParseDuration(c.IntervalSessionUsageSyncWithDatabase); err != nil {
		return fmt.Errorf("invalid interval_session_usage_sync_with_database: %w", err)
	}
	if _, err := time.ParseDuration(c.IntervalSessionUsageValidate); err != nil {
		return fmt.Errorf("invalid interval_session_usage_validate: %w", err)
	}
	if _, err := time.ParseDuration(c.IntervalSessionValidate); err != nil {
		return fmt.Errorf("invalid interval_session_validate: %w", err)
	}
	if _, err := time.ParseDuration(c.IntervalSpeedtest); err != nil {
		return fmt.Errorf("invalid interval_speedtest: %w", err)
	}
	if _, err := time.ParseDuration(c.IntervalStatusUpdate); err != nil {
		return fmt.Errorf("invalid interval_status_update: %w", err)
	}

	// Ensure the Moniker field is not empty.
	if c.Moniker == "" {
		return errors.New("moniker cannot be empty")
	}

	// Ensure the RemoteAddrs field is not empty.
	if len(c.RemoteAddrs) == 0 {
		return errors.New("remote_addrs cannot be empty")
	}

	// Validate each address in the RemoteAddrs field.
	for _, addr := range c.RemoteAddrs {
		if err := validateRemoteAddr(addr); err != nil {
			return fmt.Errorf("invalid remote_addr %s: %w", addr, err)
		}
	}

	// Validate the node type.
	validServiceTypes := map[string]bool{
		types.ServiceTypeV2Ray.String():     true,
		types.ServiceTypeWireGuard.String(): true,
		types.ServiceTypeOpenVPN.String():   true,
	}
	if !validServiceTypes[c.ServiceType] {
		return errors.New("type must be one of: v2ray, wireguard, openvpn")
	}

	// Ensure the TLSCertPath field is not empty.
	if c.TLSCertPath == "" {
		return errors.New("tls_cert_path cannot be empty")
	}

	// Ensure the TLSKeyPath field is not empty.
	if c.TLSKeyPath == "" {
		return errors.New("tls_key_path cannot be empty")
	}

	return nil
}

// SetForFlags adds node configuration flags to the specified FlagSet.
func (c *NodeConfig) SetForFlags(f *pflag.FlagSet) {
	f.StringVar(&c.APIPort, "node.api-port", c.APIPort, "port for API access")
	f.StringVar(&c.GigabytePrices, "node.gigabyte-prices", c.GigabytePrices, "pricing information for gigabytes")
	f.StringVar(&c.HourlyPrices, "node.hourly-prices", c.HourlyPrices, "pricing information for hourly usage")
	f.StringVar(&c.IntervalBestRPCAddr, "node.interval-best-rpc-addr", c.IntervalBestRPCAddr, "interval for checking the best RPC address")
	f.StringVar(&c.IntervalGeoIPLocation, "node.interval-geo-ip-location", c.IntervalGeoIPLocation, "interval for checking GeoIP location")
	f.StringVar(&c.IntervalSessionUsageSyncWithBlockchain, "node.interval-session-usage-sync-with-blockchain", c.IntervalSessionUsageSyncWithBlockchain, "interval for syncing session usage with blockchain")
	f.StringVar(&c.IntervalSessionUsageSyncWithDatabase, "node.interval-session-usage-sync-with-database", c.IntervalSessionUsageSyncWithDatabase, "interval for syncing session usage with database")
	f.StringVar(&c.IntervalSessionUsageValidate, "node.interval-session-usage-validate", c.IntervalSessionUsageValidate, "interval for validating session usage")
	f.StringVar(&c.IntervalSessionValidate, "node.interval-session-validate", c.IntervalSessionValidate, "interval for validating sessions")
	f.StringVar(&c.IntervalSpeedtest, "node.interval-speedtest", c.IntervalSpeedtest, "interval for performing speed tests")
	f.StringVar(&c.IntervalStatusUpdate, "node.interval-status-update", c.IntervalStatusUpdate, "interval for updating node status")
	f.StringVar(&c.Moniker, "node.moniker", c.Moniker, "moniker (identifier) for the node")
	f.StringSliceVar(&c.RemoteAddrs, "node.remote-addrs", c.RemoteAddrs, "list of remote addresses for the node")
	f.StringVar(&c.ServiceType, "node.service-type", c.ServiceType, "service type of the node (e.g., v2ray, wireguard, openvpn)")
}

// DefaultNodeConfig returns a NodeConfig instance with default values.
func DefaultNodeConfig() *NodeConfig {
	return &NodeConfig{
		APIPort:                                fmt.Sprintf("%d", utils.RandomPort()),
		GigabytePrices:                         "udvpn:0.01,0",
		HourlyPrices:                           "udvpn:0.02,0",
		IntervalBestRPCAddr:                    (5 * time.Minute).String(),
		IntervalGeoIPLocation:                  (6 * time.Hour).String(),
		IntervalSessionUsageSyncWithBlockchain: (2*time.Hour - 5*time.Minute).String(),
		IntervalSessionUsageSyncWithDatabase:   (3 * time.Second).String(),
		IntervalSessionUsageValidate:           (3 * time.Second).String(),
		IntervalSessionValidate:                (5 * time.Minute).String(),
		IntervalSpeedtest:                      (7 * 24 * time.Hour).String(),
		IntervalStatusUpdate:                   (1*time.Hour - 5*time.Minute).String(),
		Moniker:                                randMoniker(),
		RemoteAddrs:                            []string{},
		ServiceType:                            randServiceType().String(),
		TLSCertPath:                            "",
		TLSKeyPath:                             "",
	}
}

func randServiceType() types.ServiceType {
	services := []types.ServiceType{
		types.ServiceTypeWireGuard,
		types.ServiceTypeV2Ray,
		types.ServiceTypeOpenVPN,
	}

	return services[rand.IntN(len(services))]
}

func randMoniker() string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"

	var result strings.Builder
	result.Grow(8)

	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			result.WriteByte(letters[rand.IntN(len(letters))])
		} else {
			result.WriteByte(numbers[rand.IntN(len(numbers))])
		}
	}

	return result.String()
}

func validateRemoteAddr(addr string) error {
	// Ensure the address is not empty or too long.
	if len(addr) == 0 {
		return errors.New("addr cannot be empty")
	}
	if len(addr) > MaxRemoteAddrLen {
		return fmt.Errorf("addr cannot be longer than %d chars", MaxRemoteAddrLen)
	}

	// Validate the IP address format.
	if ip := net.ParseIP(addr); ip != nil {
		if ipv4 := ip.To4(); ipv4 != nil {
			return nil
		}
		if ipv6 := ip.To16(); ipv6 != nil {
			return nil
		}

		return errors.New("invalid ip addr")
	}

	return nil
}
