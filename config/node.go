package config

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net"
	"strings"
	"time"

	"github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/utils"
)

const MaxRemoteAddrLen = (1 << 6) - 1

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
	Type                                   string   `mapstructure:"type"`                                        // Type is the service type of the node.
}

func (c *NodeConfig) APIListenAddr() string {
	return fmt.Sprintf("0.0.0.0:%d", c.APIListenPort())
}

func (c *NodeConfig) APIListenPort() uint16 {
	return c.GetAPIPort().InFrom
}

func (c *NodeConfig) APIRemoteAddrs() []string {
	addrs := make([]string, len(c.RemoteAddrs))
	for i, addr := range c.RemoteAddrs {
		addrs[i] = fmt.Sprintf("https://%s:%d", addr, c.APIListenPort())
	}

	return addrs
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

// GetType returns the Type field.
func (c *NodeConfig) GetType() types.ServiceType {
	return types.ServiceTypeFromString(c.Type)
}

// Validate validates the node configuration.
func (c *NodeConfig) Validate() error {
	if c.APIPort == "" {
		return errors.New("api_port cannot be empty")
	}
	if _, err := types.NewPortFromString(c.APIPort); err != nil {
		return fmt.Errorf("invalid api_port: %w", err)
	}
	if _, err := v1.NewPricesFromString(c.GigabytePrices); err != nil {
		return fmt.Errorf("invalid gigabyte_prices: %w", err)
	}
	if _, err := v1.NewPricesFromString(c.HourlyPrices); err != nil {
		return fmt.Errorf("invalid hourly_prices: %w", err)
	}
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
	if c.Moniker == "" {
		return errors.New("moniker cannot be empty")
	}
	if len(c.RemoteAddrs) == 0 {
		return errors.New("remote_addrs cannot be empty")
	}
	for _, addr := range c.RemoteAddrs {
		if err := validateRemoteAddr(addr); err != nil {
			return fmt.Errorf("invalid remote_addr %s: %w", addr, err)
		}
	}

	validTypes := map[string]bool{
		types.ServiceTypeV2Ray.String():     true,
		types.ServiceTypeWireGuard.String(): true,
	}
	if !validTypes[c.Type] {
		return errors.New("type must be one of: v2ray, wireguard")
	}

	return nil
}

func DefaultNodeConfig() NodeConfig {
	return NodeConfig{
		APIPort:                                fmt.Sprintf("%d", utils.RandomPort()),
		GigabytePrices:                         "0.01;0;udvpn",
		HourlyPrices:                           "0.02;0;udvpn",
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
		Type:                                   randServiceType().String(),
	}
}

func randServiceType() types.ServiceType {
	if rand.Int()%2 == 0 {
		return types.ServiceTypeV2Ray
	}
	return types.ServiceTypeWireGuard
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
	if len(addr) == 0 {
		return errors.New("addr cannot be empty")
	}
	if len(addr) > MaxRemoteAddrLen {
		return fmt.Errorf("addr cannot be longer than %d chars", MaxRemoteAddrLen)
	}

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
