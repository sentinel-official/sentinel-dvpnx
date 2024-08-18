package flags

import (
	"github.com/spf13/cobra"
)

// Default values for node flags.
const (
	DefaultNodeGigabytePrices                         = ""     // Default value for gigabyte prices
	DefaultNodeHourlyPrices                           = ""     // Default value for hourly prices
	DefaultNodeIntervalBestRPCAddr                    = "15s"  // Default duration for checking the best RPC address
	DefaultNodeIntervalGeoIPLocation                  = "24h"  // Default duration for checking GeoIP location
	DefaultNodeIntervalSessionUsageSyncWithBlockchain = "1h"   // Default duration for syncing session usage with the blockchain
	DefaultNodeIntervalSessionUsageSyncWithDatabase   = "5s"   // Default duration for syncing session usage with the database
	DefaultNodeIntervalSessionUsageValidate           = "5s"   // Default duration for validating session usage
	DefaultNodeIntervalSessionValidate                = "1m"   // Default duration for validating sessions
	DefaultNodeIntervalSpeedtest                      = "168h" // Default duration for performing speed tests
	DefaultNodeIntervalStatusUpdate                   = "30m"  // Default duration for updating the node status
	DefaultNodeListenOn                               = ""     // Default listen address
	DefaultNodeMoniker                                = ""     // Default node moniker
	DefaultNodePublicIPv4Addr                         = ""     // Default public IPv4 address
	DefaultNodeRemoteURL                              = ""     // Default remote URL
	DefaultNodeType                                   = ""     // Default node type
)

// GetNodeGigabytePrices retrieves the "node.gigabyte-prices" flag value from the command.
func GetNodeGigabytePrices(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.gigabyte-prices")
}

// GetNodeHourlyPrices retrieves the "node.hourly-prices" flag value from the command.
func GetNodeHourlyPrices(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.hourly-prices")
}

// GetNodeIntervalBestRPCAddr retrieves the "node.interval-best-rpc-addr" flag value from the command.
func GetNodeIntervalBestRPCAddr(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.interval-best-rpc-addr")
}

// GetNodeIntervalGeoIPLocation retrieves the "node.interval-geoip-location" flag value from the command.
func GetNodeIntervalGeoIPLocation(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.interval-geoip-location")
}

// GetNodeIntervalSessionUsageSyncWithBlockchain retrieves the "node.interval-session-usage-sync-with-blockchain" flag value from the command.
func GetNodeIntervalSessionUsageSyncWithBlockchain(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.interval-session-usage-sync-with-blockchain")
}

// GetNodeIntervalSessionUsageSyncWithDatabase retrieves the "node.interval-session-usage-sync-with-database" flag value from the command.
func GetNodeIntervalSessionUsageSyncWithDatabase(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.interval-session-usage-sync-with-database")
}

// GetNodeIntervalSessionUsageValidate retrieves the "node.interval-session-usage-validate" flag value from the command.
func GetNodeIntervalSessionUsageValidate(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.interval-session-usage-validate")
}

// GetNodeIntervalSessionValidate retrieves the "node.interval-session-validate" flag value from the command.
func GetNodeIntervalSessionValidate(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.interval-session-validate")
}

// GetNodeIntervalSpeedtest retrieves the "node.interval-speedtest" flag value from the command.
func GetNodeIntervalSpeedtest(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.interval-speedtest")
}

// GetNodeIntervalStatusUpdate retrieves the "node.interval-update-status" flag value from the command.
func GetNodeIntervalStatusUpdate(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.interval-update-status")
}

// GetNodeListenOn retrieves the "node.listen-on" flag value from the command.
func GetNodeListenOn(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.listen-on")
}

// GetNodeMoniker retrieves the "node.moniker" flag value from the command.
func GetNodeMoniker(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.moniker")
}

// GetNodePublicIPv4Addr retrieves the "node.public-ipv4-addr" flag value from the command.
func GetNodePublicIPv4Addr(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.public-ipv4-addr")
}

// GetNodeRemoteURL retrieves the "node.remote-url" flag value from the command.
func GetNodeRemoteURL(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.remote-url")
}

// GetNodeType retrieves the "node.type" flag value from the command.
func GetNodeType(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.type")
}

// SetFlagNodeGigabytePrices adds the "node.gigabyte-prices" flag to the command.
func SetFlagNodeGigabytePrices(cmd *cobra.Command) {
	cmd.Flags().String("node.gigabyte-prices", DefaultNodeGigabytePrices, "Gigabyte pricing information for the node.")
}

// SetFlagNodeHourlyPrices adds the "node.hourly-prices" flag to the command.
func SetFlagNodeHourlyPrices(cmd *cobra.Command) {
	cmd.Flags().String("node.hourly-prices", DefaultNodeHourlyPrices, "Hourly pricing information for the node.")
}

// SetFlagNodeIntervalBestRPCAddr adds the "node.interval-best-rpc-addr" flag to the command.
func SetFlagNodeIntervalBestRPCAddr(cmd *cobra.Command) {
	cmd.Flags().String("node.interval-best-rpc-addr", DefaultNodeIntervalBestRPCAddr, "Duration between checking the best RPC address.")
}

// SetFlagNodeIntervalGeoIPLocation adds the "node.interval-geoip-location" flag to the command.
func SetFlagNodeIntervalGeoIPLocation(cmd *cobra.Command) {
	cmd.Flags().String("node.interval-geoip-location", DefaultNodeIntervalGeoIPLocation, "Duration between checking the GeoIP location.")
}

// SetFlagNodeIntervalSessionUsageSyncWithBlockchain adds the "node.interval-session-usage-sync-with-blockchain" flag to the command.
func SetFlagNodeIntervalSessionUsageSyncWithBlockchain(cmd *cobra.Command) {
	cmd.Flags().String("node.interval-session-usage-sync-with-blockchain", DefaultNodeIntervalSessionUsageSyncWithBlockchain, "Duration between syncing session usage with the blockchain.")
}

// SetFlagNodeIntervalSessionUsageSyncWithDatabase adds the "node.interval-session-usage-sync-with-database" flag to the command.
func SetFlagNodeIntervalSessionUsageSyncWithDatabase(cmd *cobra.Command) {
	cmd.Flags().String("node.interval-session-usage-sync-with-database", DefaultNodeIntervalSessionUsageSyncWithDatabase, "Duration between syncing session usage with the database.")
}

// SetFlagNodeIntervalSessionUsageValidate adds the "node.interval-session-usage-validate" flag to the command.
func SetFlagNodeIntervalSessionUsageValidate(cmd *cobra.Command) {
	cmd.Flags().String("node.interval-session-usage-validate", DefaultNodeIntervalSessionUsageValidate, "Duration between validating session usage.")
}

// SetFlagNodeIntervalSessionValidate adds the "node.interval-session-validate" flag to the command.
func SetFlagNodeIntervalSessionValidate(cmd *cobra.Command) {
	cmd.Flags().String("node.interval-session-validate", DefaultNodeIntervalSessionValidate, "Duration between validating sessions.")
}

// SetFlagNodeIntervalSpeedtest adds the "node.interval-speedtest" flag to the command.
func SetFlagNodeIntervalSpeedtest(cmd *cobra.Command) {
	cmd.Flags().String("node.interval-speedtest", DefaultNodeIntervalSpeedtest, "Duration between performing speed tests.")
}

// SetFlagNodeIntervalStatusUpdate adds the "node.interval-update-status" flag to the command.
func SetFlagNodeIntervalStatusUpdate(cmd *cobra.Command) {
	cmd.Flags().String("node.interval-update-status", DefaultNodeIntervalStatusUpdate, "Duration between updating the status of the node.")
}

// SetFlagNodeListenOn adds the "node.listen-on" flag to the command.
func SetFlagNodeListenOn(cmd *cobra.Command) {
	cmd.Flags().String("node.listen-on", DefaultNodeListenOn, "Address on which the node should listen for connections.")
}

// SetFlagNodeMoniker adds the "node.moniker" flag to the command.
func SetFlagNodeMoniker(cmd *cobra.Command) {
	cmd.Flags().String("node.moniker", DefaultNodeMoniker, "Name or identifier for the node.")
}

// SetFlagNodePublicIPv4Addr adds the "node.public-ipv4-addr" flag to the command.
func SetFlagNodePublicIPv4Addr(cmd *cobra.Command) {
	cmd.Flags().String("node.public-ipv4-addr", DefaultNodePublicIPv4Addr, "Public IPv4 address of the node.")
}

// SetFlagNodeRemoteURL adds the "node.remote-url" flag to the command.
func SetFlagNodeRemoteURL(cmd *cobra.Command) {
	cmd.Flags().String("node.remote-url", DefaultNodeRemoteURL, "URL for remote operations related to the node.")
}

// SetFlagNodeType adds the "node.type" flag to the command.
func SetFlagNodeType(cmd *cobra.Command) {
	cmd.Flags().String("node.type", DefaultNodeType, "Type of the node.")
}

// AddNodeFlags adds all node-related flags to the given Cobra command.
func AddNodeFlags(cmd *cobra.Command) {
	SetFlagNodeGigabytePrices(cmd)
	SetFlagNodeHourlyPrices(cmd)
	SetFlagNodeIntervalSessionUsageSyncWithBlockchain(cmd)
	SetFlagNodeIntervalSessionUsageSyncWithDatabase(cmd)
	SetFlagNodeIntervalSessionUsageValidate(cmd)
	SetFlagNodeIntervalSessionValidate(cmd)
	SetFlagNodeIntervalStatusUpdate(cmd)
	SetFlagNodeIntervalBestRPCAddr(cmd)
	SetFlagNodeIntervalGeoIPLocation(cmd)
	SetFlagNodeIntervalSpeedtest(cmd)
	SetFlagNodeListenOn(cmd)
	SetFlagNodeMoniker(cmd)
	SetFlagNodePublicIPv4Addr(cmd)
	SetFlagNodeRemoteURL(cmd)
	SetFlagNodeType(cmd)
}
