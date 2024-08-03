package flags

import (
	"time"

	"github.com/spf13/cobra"
)

// GetNodeIntervalSetSessionsFromCmd retrieves the "node.interval-set-sessions" flag value from the command.
func GetNodeIntervalSetSessionsFromCmd(cmd *cobra.Command) (time.Duration, error) {
	return cmd.Flags().GetDuration("node.interval-set-sessions")
}

// GetNodeIntervalUpdateSessionsFromCmd retrieves the "node.interval-update-sessions" flag value from the command.
func GetNodeIntervalUpdateSessionsFromCmd(cmd *cobra.Command) (time.Duration, error) {
	return cmd.Flags().GetDuration("node.interval-update-sessions")
}

// GetNodeIntervalUpdateStatusFromCmd retrieves the "node.interval-update-status" flag value from the command.
func GetNodeIntervalUpdateStatusFromCmd(cmd *cobra.Command) (time.Duration, error) {
	return cmd.Flags().GetDuration("node.interval-update-status")
}

// GetNodeIPv4AddressFromCmd retrieves the "node.ipv4-address" flag value from the command.
func GetNodeIPv4AddressFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.ipv4-address")
}

// GetNodeListenOnFromCmd retrieves the "node.listen-on" flag value from the command.
func GetNodeListenOnFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.listen-on")
}

// GetNodeMonikerFromCmd retrieves the "node.moniker" flag value from the command.
func GetNodeMonikerFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.moniker")
}

// GetNodeGigabytePricesFromCmd retrieves the "node.gigabyte-prices" flag value from the command.
func GetNodeGigabytePricesFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.gigabyte-prices")
}

// GetNodeHourlyPricesFromCmd retrieves the "node.hourly-prices" flag value from the command.
func GetNodeHourlyPricesFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.hourly-prices")
}

// GetNodeRemoteURLFromCmd retrieves the "node.remote-url" flag value from the command.
func GetNodeRemoteURLFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.remote-url")
}

// GetNodeTypeFromCmd retrieves the "node.type" flag value from the command.
func GetNodeTypeFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("node.type")
}

// SetFlagIntervalSetSessions adds the "node.interval-set-sessions" flag to the command.
func SetFlagIntervalSetSessions(cmd *cobra.Command) {
	cmd.Flags().Duration("node.interval-set-sessions", 0, "Duration between setting sessions for the node.")
}

// SetFlagIntervalUpdateSessions adds the "node.interval-update-sessions" flag to the command.
func SetFlagIntervalUpdateSessions(cmd *cobra.Command) {
	cmd.Flags().Duration("node.interval-update-sessions", 0, "Duration between updating sessions for the node.")
}

// SetFlagIntervalUpdateStatus adds the "node.interval-update-status" flag to the command.
func SetFlagIntervalUpdateStatus(cmd *cobra.Command) {
	cmd.Flags().Duration("node.interval-update-status", 0, "Duration between updating the status of the node.")
}

// SetFlagIPv4Address adds the "node.ipv4-address" flag to the command.
func SetFlagIPv4Address(cmd *cobra.Command) {
	cmd.Flags().String("node.ipv4-address", "", "Public IPv4 address of the node.")
}

// SetFlagListenOn adds the "node.listen-on" flag to the command.
func SetFlagListenOn(cmd *cobra.Command) {
	cmd.Flags().String("node.listen-on", "", "Address on which the node should listen for connections.")
}

// SetFlagMoniker adds the "node.moniker" flag to the command.
func SetFlagMoniker(cmd *cobra.Command) {
	cmd.Flags().String("node.moniker", "", "Name or identifier for the node.")
}

// SetFlagGigabytePrices adds the "node.gigabyte-prices" flag to the command.
func SetFlagGigabytePrices(cmd *cobra.Command) {
	cmd.Flags().String("node.gigabyte-prices", "", "Gigabyte pricing information for the node.")
}

// SetFlagHourlyPrices adds the "node.hourly-prices" flag to the command.
func SetFlagHourlyPrices(cmd *cobra.Command) {
	cmd.Flags().String("node.hourly-prices", "", "Hourly pricing information for the node.")
}

// SetFlagRemoteURL adds the "node.remote-url" flag to the command.
func SetFlagRemoteURL(cmd *cobra.Command) {
	cmd.Flags().String("node.remote-url", "", "URL for remote operations related to the node.")
}

// SetFlagType adds the "node.type" flag to the command.
func SetFlagType(cmd *cobra.Command) {
	cmd.Flags().String("node.type", "", "Type of the node.")
}
