package options

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/sentinel-official/dvpn-node/cmd/flags"
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

// NewDefaultNode creates a new NodeOptions instance with default values.
func NewDefaultNode() *NodeOptions {
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

// AddNodeFlagsToCmd adds all node-related flags to the given Cobra command.
func AddNodeFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagIntervalSetSessions(cmd)
	flags.SetFlagIntervalUpdateSessions(cmd)
	flags.SetFlagIntervalUpdateStatus(cmd)
	flags.SetFlagIPv4Address(cmd)
	flags.SetFlagListenOn(cmd)
	flags.SetFlagMoniker(cmd)
	flags.SetFlagGigabytePrices(cmd)
	flags.SetFlagHourlyPrices(cmd)
	flags.SetFlagRemoteURL(cmd)
	flags.SetFlagType(cmd)
}

// NewNodeOptionsFromCmd creates and returns NodeOptions from the given Cobra command's flags.
func NewNodeOptionsFromCmd(cmd *cobra.Command) (*NodeOptions, error) {
	// Retrieve the interval set sessions flag value from the command.
	intervalSetSessions, err := flags.GetNodeIntervalSetSessionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval update sessions flag value from the command.
	intervalUpdateSessions, err := flags.GetNodeIntervalUpdateSessionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval update status flag value from the command.
	intervalUpdateStatus, err := flags.GetNodeIntervalUpdateStatusFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the IPv4 address flag value from the command.
	ipv4Address, err := flags.GetNodeIPv4AddressFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the listen-on flag value from the command.
	listenOn, err := flags.GetNodeListenOnFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the moniker flag value from the command.
	moniker, err := flags.GetNodeMonikerFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gigabyte prices flag value from the command.
	gigabytePrices, err := flags.GetNodeGigabytePricesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the hourly prices flag value from the command.
	hourlyPrices, err := flags.GetNodeHourlyPricesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the remote URL flag value from the command.
	remoteURL, err := flags.GetNodeRemoteURLFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the node type flag value from the command.
	nodeType, err := flags.GetNodeTypeFromCmd(cmd)
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
