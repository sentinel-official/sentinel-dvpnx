package flags

import (
	"time"

	"github.com/spf13/cobra"
)

// GetNodeGigabytePricesFromCmd retrieves the "node.gigabyte-prices" flag value from the command.
func GetNodeGigabytePricesFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.gigabyte-prices")
}

// GetNodeHourlyPricesFromCmd retrieves the "node.hourly-prices" flag value from the command.
func GetNodeHourlyPricesFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.hourly-prices")
}

// GetNodeIntervalSessionUsageSyncWithBlockchainFromCmd retrieves the "node.interval-session-usage-sync-with-blockchain" flag value from the command.
func GetNodeIntervalSessionUsageSyncWithBlockchainFromCmd(cmd *cobra.Command) (time.Duration, error) {
	return cmd.Flags().GetDuration("node.interval-session-usage-sync-with-blockchain")
}

// GetNodeIntervalSessionUsageSyncWithDatabaseFromCmd retrieves the "node.interval-session-usage-sync-with-database" flag value from the command.
func GetNodeIntervalSessionUsageSyncWithDatabaseFromCmd(cmd *cobra.Command) (time.Duration, error) {
	return cmd.Flags().GetDuration("node.interval-session-usage-sync-with-database")
}

// GetNodeIntervalSessionUsageValidateFromCmd retrieves the "node.interval-session-usage-validate" flag value from the command.
func GetNodeIntervalSessionUsageValidateFromCmd(cmd *cobra.Command) (time.Duration, error) {
	return cmd.Flags().GetDuration("node.interval-session-usage-validate")
}

// GetNodeIntervalSessionValidateFromCmd retrieves the "node.interval-session-validate" flag value from the command.
func GetNodeIntervalSessionValidateFromCmd(cmd *cobra.Command) (time.Duration, error) {
	return cmd.Flags().GetDuration("node.interval-session-validate")
}

// GetNodeIntervalUpdateStatusFromCmd retrieves the "node.interval-update-status" flag value from the command.
func GetNodeIntervalUpdateStatusFromCmd(cmd *cobra.Command) (time.Duration, error) {
	return cmd.Flags().GetDuration("node.interval-update-status")
}

// GetNodeListenOnFromCmd retrieves the "node.listen-on" flag value from the command.
func GetNodeListenOnFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.listen-on")
}

// GetNodeMonikerFromCmd retrieves the "node.moniker" flag value from the command.
func GetNodeMonikerFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.moniker")
}

// GetNodePublicIPv4AddrFromCmd retrieves the "node.public-ipv4-addr" flag value from the command.
func GetNodePublicIPv4AddrFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.public-ipv4-addr")
}

// GetNodeRemoteURLFromCmd retrieves the "node.remote-url" flag value from the command.
func GetNodeRemoteURLFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.remote-url")
}

// GetNodeTypeFromCmd retrieves the "node.type" flag value from the command.
func GetNodeTypeFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.type")
}

// SetFlagNodeGigabytePrices adds the "node.gigabyte-prices" flag to the command.
func SetFlagNodeGigabytePrices(cmd *cobra.Command) {
	cmd.Flags().String("node.gigabyte-prices", "", "Gigabyte pricing information for the node.")
}

// SetFlagNodeHourlyPrices adds the "node.hourly-prices" flag to the command.
func SetFlagNodeHourlyPrices(cmd *cobra.Command) {
	cmd.Flags().String("node.hourly-prices", "", "Hourly pricing information for the node.")
}

// SetFlagNodeIntervalSessionUsageSyncWithBlockchain adds the "node.interval-session-usage-sync-with-blockchain" flag to the command.
func SetFlagNodeIntervalSessionUsageSyncWithBlockchain(cmd *cobra.Command) {
	cmd.Flags().Duration("node.interval-session-usage-sync-with-blockchain", 0, "Duration between syncing session usage with the blockchain.")
}

// SetFlagNodeIntervalSessionUsageSyncWithDatabase adds the "node.interval-session-usage-sync-with-database" flag to the command.
func SetFlagNodeIntervalSessionUsageSyncWithDatabase(cmd *cobra.Command) {
	cmd.Flags().Duration("node.interval-session-usage-sync-with-database", 0, "Duration between syncing session usage with the database.")
}

// SetFlagNodeIntervalSessionUsageValidate adds the "node.interval-session-usage-validate" flag to the command.
func SetFlagNodeIntervalSessionUsageValidate(cmd *cobra.Command) {
	cmd.Flags().Duration("node.interval-session-usage-validate", 0, "Duration between validating session usage.")
}

// SetFlagNodeIntervalSessionValidate adds the "node.interval-session-validate" flag to the command.
func SetFlagNodeIntervalSessionValidate(cmd *cobra.Command) {
	cmd.Flags().Duration("node.interval-session-validate", 0, "Duration between validating sessions.")
}

// SetFlagNodeIntervalUpdateStatus adds the "node.interval-update-status" flag to the command.
func SetFlagNodeIntervalUpdateStatus(cmd *cobra.Command) {
	cmd.Flags().Duration("node.interval-update-status", 0, "Duration between updating the status of the node.")
}

// SetFlagNodeListenOn adds the "node.listen-on" flag to the command.
func SetFlagNodeListenOn(cmd *cobra.Command) {
	cmd.Flags().String("node.listen-on", "", "Address on which the node should listen for connections.")
}

// SetFlagNodeMoniker adds the "node.moniker" flag to the command.
func SetFlagNodeMoniker(cmd *cobra.Command) {
	cmd.Flags().String("node.moniker", "", "Name or identifier for the node.")
}

// SetFlagNodePublicIPv4Addr adds the "node.public-ipv4-addr" flag to the command.
func SetFlagNodePublicIPv4Addr(cmd *cobra.Command) {
	cmd.Flags().String("node.public-ipv4-addr", "", "Public IPv4 address of the node.")
}

// SetFlagNodeRemoteURL adds the "node.remote-url" flag to the command.
func SetFlagNodeRemoteURL(cmd *cobra.Command) {
	cmd.Flags().String("node.remote-url", "", "URL for remote operations related to the node.")
}

// SetFlagNodeType adds the "node.type" flag to the command.
func SetFlagNodeType(cmd *cobra.Command) {
	cmd.Flags().String("node.type", "", "Type of the node.")
}
