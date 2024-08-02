package options

import (
	"time"

	"github.com/spf13/cobra"
)

// NodeOptions represents options for node configuration.
type NodeOptions struct {
	IntervalSetSessions    time.Duration `json:"interval_set_sessions" toml:"interval_set_sessions"`       // IntervalSetSessions is the duration between setting sessions.
	IntervalUpdateSessions time.Duration `json:"interval_update_sessions" toml:"interval_update_sessions"` // IntervalUpdateSessions is the duration between updating sessions.
	IntervalUpdateStatus   time.Duration `json:"interval_update_status" toml:"interval_update_status"`     // IntervalUpdateStatus is the duration between updating status.
	IPv4Address            string        `json:"ipv4_address" toml:"ipv4_address"`                         // IPv4Address is the IPv4 address of the node.
	ListenOn               string        `json:"listen_on" toml:"listen_on"`                               // ListenOn is the address on which the node listens.
	Moniker                string        `json:"moniker" toml:"moniker"`                                   // Moniker is the name or identifier for the node.
	GigabytePrices         string        `json:"gigabyte_prices" toml:"gigabyte_prices"`                   // GigabytePrices is the pricing information for gigabytes.
	HourlyPrices           string        `json:"hourly_prices" toml:"hourly_prices"`                       // HourlyPrices is the pricing information for hourly usage.
	RemoteURL              string        `json:"remote_url" toml:"remote_url"`                             // RemoteURL is the URL for remote operations.
	Type                   string        `json:"type" toml:"type"`                                         // Type indicates the type of node.
}

// NewDefaultNodeOptions creates a new NodeOptions instance with default values.
func NewDefaultNodeOptions() *NodeOptions {
	return &NodeOptions{}
}

// WithIntervalSetSessions sets the IntervalSetSessions field and returns the modified NodeOptions instance.
func (n *NodeOptions) WithIntervalSetSessions(v time.Duration) *NodeOptions {
	n.IntervalSetSessions = v
	return n
}

// WithIntervalUpdateSessions sets the IntervalUpdateSessions field and returns the modified NodeOptions instance.
func (n *NodeOptions) WithIntervalUpdateSessions(v time.Duration) *NodeOptions {
	n.IntervalUpdateSessions = v
	return n
}

// WithIntervalUpdateStatus sets the IntervalUpdateStatus field and returns the modified NodeOptions instance.
func (n *NodeOptions) WithIntervalUpdateStatus(v time.Duration) *NodeOptions {
	n.IntervalUpdateStatus = v
	return n
}

// WithIPv4Address sets the IPv4Address field and returns the modified NodeOptions instance.
func (n *NodeOptions) WithIPv4Address(v string) *NodeOptions {
	n.IPv4Address = v
	return n
}

// WithListenOn sets the ListenOn field and returns the modified NodeOptions instance.
func (n *NodeOptions) WithListenOn(v string) *NodeOptions {
	n.ListenOn = v
	return n
}

// WithMoniker sets the Moniker field and returns the modified NodeOptions instance.
func (n *NodeOptions) WithMoniker(v string) *NodeOptions {
	n.Moniker = v
	return n
}

// WithGigabytePrices sets the GigabytePrices field and returns the modified NodeOptions instance.
func (n *NodeOptions) WithGigabytePrices(v string) *NodeOptions {
	n.GigabytePrices = v
	return n
}

// WithHourlyPrices sets the HourlyPrices field and returns the modified NodeOptions instance.
func (n *NodeOptions) WithHourlyPrices(v string) *NodeOptions {
	n.HourlyPrices = v
	return n
}

// WithRemoteURL sets the RemoteURL field and returns the modified NodeOptions instance.
func (n *NodeOptions) WithRemoteURL(v string) *NodeOptions {
	n.RemoteURL = v
	return n
}

// WithType sets the Type field and returns the modified NodeOptions instance.
func (n *NodeOptions) WithType(v string) *NodeOptions {
	n.Type = v
	return n
}

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

// AddNodeFlagsToCmd adds all node-related flags to the given Cobra command.
func AddNodeFlagsToCmd(cmd *cobra.Command) {
	SetFlagIntervalSetSessions(cmd)
	SetFlagIntervalUpdateSessions(cmd)
	SetFlagIntervalUpdateStatus(cmd)
	SetFlagIPv4Address(cmd)
	SetFlagListenOn(cmd)
	SetFlagMoniker(cmd)
	SetFlagGigabytePrices(cmd)
	SetFlagHourlyPrices(cmd)
	SetFlagRemoteURL(cmd)
	SetFlagType(cmd)
}

// NewNodeOptionsFromCmd creates and returns NodeOptions from the given Cobra command's flags.
func NewNodeOptionsFromCmd(cmd *cobra.Command) (*NodeOptions, error) {
	// Retrieve the value of the "node.interval-set-sessions" flag.
	intervalSetSessions, err := GetNodeIntervalSetSessionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "node.interval-update-sessions" flag.
	intervalUpdateSessions, err := GetNodeIntervalUpdateSessionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "node.interval-update-status" flag.
	intervalUpdateStatus, err := GetNodeIntervalUpdateStatusFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "node.ipv4-address" flag.
	ipv4Address, err := GetNodeIPv4AddressFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "node.listen-on" flag.
	listenOn, err := GetNodeListenOnFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "node.moniker" flag.
	moniker, err := GetNodeMonikerFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "node.gigabyte-prices" flag.
	gigabytePrices, err := GetNodeGigabytePricesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "node.hourly-prices" flag.
	hourlyPrices, err := GetNodeHourlyPricesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "node.remote-url" flag.
	remoteURL, err := GetNodeRemoteURLFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "node.type" flag.
	nodeType, err := GetNodeTypeFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new NodeOptions instance populated with the retrieved flag values.
	return &NodeOptions{
		IntervalSetSessions:    intervalSetSessions,
		IntervalUpdateSessions: intervalUpdateSessions,
		IntervalUpdateStatus:   intervalUpdateStatus,
		IPv4Address:            ipv4Address,
		ListenOn:               listenOn,
		Moniker:                moniker,
		GigabytePrices:         gigabytePrices,
		HourlyPrices:           hourlyPrices,
		RemoteURL:              remoteURL,
		Type:                   nodeType,
	}, nil
}
