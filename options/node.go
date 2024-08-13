package options

import (
	"time"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/dvpn-node/cmd/flags"
)

// Node represents options for node configuration.
type Node struct {
	GigabytePrices                         cosmossdk.Coins `json:"gigabyte_prices" toml:"gigabyte_prices"`                                                         // GigabytePrices is the pricing information for gigabytes.
	HourlyPrices                           cosmossdk.Coins `json:"hourly_prices" toml:"hourly_prices"`                                                             // HourlyPrices is the pricing information for hourly usage.
	IntervalBestRPCAddr                    time.Duration   `json:"interval_best_rpc_addr" toml:"interval_best_rpc_addr"`                                           // IntervalBestRPCAddr is the duration between checking the best RPC address.
	IntervalGeoIPLocation                  time.Duration   `json:"interval_geoip_location" toml:"interval_geoip_location"`                                         // IntervalGeoIPLocation is the duration between checking the GeoIP location.
	IntervalSessionUsageSyncWithBlockchain time.Duration   `json:"interval_session_usage_sync_with_blockchain" toml:"interval_session_usage_sync_with_blockchain"` // IntervalSessionUsageSyncWithBlockchain is the duration between syncing session usage with the blockchain.
	IntervalSessionUsageSyncWithDatabase   time.Duration   `json:"interval_session_usage_sync_with_database" toml:"interval_session_usage_sync_with_database"`     // IntervalSessionUsageSyncWithDatabase is the duration between syncing session usage with the database.
	IntervalSessionUsageValidate           time.Duration   `json:"interval_session_usage_validate" toml:"interval_session_usage_validate"`                         // IntervalSessionUsageValidate is the duration between validating session usage.
	IntervalSessionValidate                time.Duration   `json:"interval_session_validate" toml:"interval_session_validate"`                                     // IntervalSessionValidate is the duration between validating sessions.
	IntervalSpeedtest                      time.Duration   `json:"interval_speedtest" toml:"interval_speedtest"`                                                   // IntervalSpeedtest is the duration between performing speed tests.
	IntervalStatusUpdate                   time.Duration   `json:"interval_status_update" toml:"interval_status_update"`                                           // IntervalStatusUpdate is the duration between updating the status of the node.
	ListenOn                               string          `json:"listen_on" toml:"listen_on"`                                                                     // ListenOn is the address on which the node listens.
	Moniker                                string          `json:"moniker" toml:"moniker"`                                                                         // Moniker is the name or identifier for the node.
	PublicIPv4Addr                         string          `json:"public_ipv4_addr" toml:"public_ipv4_addr"`                                                       // PublicIPv4Addr is the IPv4 address of the node.
	RemoteURL                              string          `json:"remote_url" toml:"remote_url"`                                                                   // RemoteURL is the URL for remote operations.
	TLSCertPath                            string          `json:"tls_cert_path" toml:"tls_cert_path"`                                                             // TLSCertPath is the path to the TLS certificate file.
	TLSKeyPath                             string          `json:"tls_key_path" toml:"tls_key_path"`                                                               // TLSKeyPath is the path to the TLS key file.
	Type                                   sdk.ServiceType `json:"type" toml:"type"`                                                                               // Type indicates the type of node.
}

// NewNode creates a new Node instance with default values.
func NewNode() *Node {
	return &Node{}
}

// WithGigabytePrices sets the GigabytePrices field and returns the modified Node instance.
func (n *Node) WithGigabytePrices(v cosmossdk.Coins) *Node {
	n.GigabytePrices = v
	return n
}

// WithHourlyPrices sets the HourlyPrices field and returns the modified Node instance.
func (n *Node) WithHourlyPrices(v cosmossdk.Coins) *Node {
	n.HourlyPrices = v
	return n
}

// WithIntervalBestRPCAddr sets the IntervalBestRPCAddr field and returns the modified Node instance.
func (n *Node) WithIntervalBestRPCAddr(v time.Duration) *Node {
	n.IntervalBestRPCAddr = v
	return n
}

// WithIntervalGeoIPLocation sets the IntervalGeoIPLocation field and returns the modified Node instance.
func (n *Node) WithIntervalGeoIPLocation(v time.Duration) *Node {
	n.IntervalGeoIPLocation = v
	return n
}

// WithIntervalSessionUsageSyncWithBlockchain sets the IntervalSessionUsageSyncWithBlockchain field and returns the modified Node instance.
func (n *Node) WithIntervalSessionUsageSyncWithBlockchain(v time.Duration) *Node {
	n.IntervalSessionUsageSyncWithBlockchain = v
	return n
}

// WithIntervalSessionUsageSyncWithDatabase sets the IntervalSessionUsageSyncWithDatabase field and returns the modified Node instance.
func (n *Node) WithIntervalSessionUsageSyncWithDatabase(v time.Duration) *Node {
	n.IntervalSessionUsageSyncWithDatabase = v
	return n
}

// WithIntervalSessionUsageValidate sets the IntervalSessionUsageValidate field and returns the modified Node instance.
func (n *Node) WithIntervalSessionUsageValidate(v time.Duration) *Node {
	n.IntervalSessionUsageValidate = v
	return n
}

// WithIntervalSessionValidate sets the IntervalSessionValidate field and returns the modified Node instance.
func (n *Node) WithIntervalSessionValidate(v time.Duration) *Node {
	n.IntervalSessionValidate = v
	return n
}

// WithIntervalSpeedtest sets the IntervalSpeedtest field and returns the modified Node instance.
func (n *Node) WithIntervalSpeedtest(v time.Duration) *Node {
	n.IntervalSpeedtest = v
	return n
}

// WithIntervalStatusUpdate sets the IntervalStatusUpdate field and returns the modified Node instance.
func (n *Node) WithIntervalStatusUpdate(v time.Duration) *Node {
	n.IntervalStatusUpdate = v
	return n
}

// WithListenOn sets the ListenOn field and returns the modified Node instance.
func (n *Node) WithListenOn(v string) *Node {
	n.ListenOn = v
	return n
}

// WithMoniker sets the Moniker field and returns the modified Node instance.
func (n *Node) WithMoniker(v string) *Node {
	n.Moniker = v
	return n
}

// WithPublicIPv4Addr sets the PublicIPv4Addr field and returns the modified Node instance.
func (n *Node) WithPublicIPv4Addr(v string) *Node {
	n.PublicIPv4Addr = v
	return n
}

// WithRemoteURL sets the RemoteURL field and returns the modified Node instance.
func (n *Node) WithRemoteURL(v string) *Node {
	n.RemoteURL = v
	return n
}

// WithType sets the Type field and returns the modified Node instance.
func (n *Node) WithType(v sdk.ServiceType) *Node {
	n.Type = v
	return n
}

// AddNodeFlagsToCmd adds all node-related flags to the given Cobra command.
func AddNodeFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagNodeGigabytePrices(cmd)
	flags.SetFlagNodeHourlyPrices(cmd)
	flags.SetFlagNodeIntervalSessionUsageSyncWithBlockchain(cmd)
	flags.SetFlagNodeIntervalSessionUsageSyncWithDatabase(cmd)
	flags.SetFlagNodeIntervalSessionUsageValidate(cmd)
	flags.SetFlagNodeIntervalSessionValidate(cmd)
	flags.SetFlagNodeIntervalStatusUpdate(cmd)
	flags.SetFlagNodeIntervalBestRPCAddr(cmd)
	flags.SetFlagNodeIntervalGeoIPLocation(cmd)
	flags.SetFlagNodeIntervalSpeedtest(cmd)
	flags.SetFlagNodeListenOn(cmd)
	flags.SetFlagNodeMoniker(cmd)
	flags.SetFlagNodePublicIPv4Addr(cmd)
	flags.SetFlagNodeRemoteURL(cmd)
	flags.SetFlagNodeTLSCertPath(cmd)
	flags.SetFlagNodeTLSKeyPath(cmd)
	flags.SetFlagNodeType(cmd)
}

// NewNodeFromCmd creates and returns Node from the given Cobra command's flags.
func NewNodeFromCmd(cmd *cobra.Command) (*Node, error) {
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

	// Retrieve the interval best RPC address flag value from the command.
	intervalBestRPCAddr, err := flags.GetNodeIntervalBestRPCAddrFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval GeoIP location flag value from the command.
	intervalGeoIPLocation, err := flags.GetNodeIntervalGeoIPLocationFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval session usage sync with blockchain flag value from the command.
	intervalSessionUsageSyncWithBlockchain, err := flags.GetNodeIntervalSessionUsageSyncWithBlockchainFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval session usage sync with database flag value from the command.
	intervalSessionUsageSyncWithDatabase, err := flags.GetNodeIntervalSessionUsageSyncWithDatabaseFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval session usage validate flag value from the command.
	intervalSessionUsageValidate, err := flags.GetNodeIntervalSessionUsageValidateFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval session validate flag value from the command.
	intervalSessionValidate, err := flags.GetNodeIntervalSessionValidateFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval speedtest flag value from the command.
	intervalSpeedtest, err := flags.GetNodeIntervalSpeedtestFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval status update flag value from the command.
	intervalStatusUpdate, err := flags.GetNodeIntervalStatusUpdateFromCmd(cmd)
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

	// Retrieve the public IPv4 address flag value from the command.
	publicIPv4Addr, err := flags.GetNodePublicIPv4AddrFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the remote URL flag value from the command.
	remoteURL, err := flags.GetNodeRemoteURLFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the TLS certificate path flag value from the command.
	tlsCertPath, err := flags.GetNodeTLSCertPathFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the TLS key path flag value from the command.
	tlsKeyPath, err := flags.GetNodeTLSKeyPathFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the node type flag value from the command.
	nodeType, err := flags.GetNodeTypeFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new Node instance populated with the retrieved flag values.
	return &Node{
		GigabytePrices:                         gigabytePrices,
		HourlyPrices:                           hourlyPrices,
		IntervalBestRPCAddr:                    intervalBestRPCAddr,
		IntervalGeoIPLocation:                  intervalGeoIPLocation,
		IntervalSessionUsageSyncWithBlockchain: intervalSessionUsageSyncWithBlockchain,
		IntervalSessionUsageSyncWithDatabase:   intervalSessionUsageSyncWithDatabase,
		IntervalSessionUsageValidate:           intervalSessionUsageValidate,
		IntervalSessionValidate:                intervalSessionValidate,
		IntervalSpeedtest:                      intervalSpeedtest,
		IntervalStatusUpdate:                   intervalStatusUpdate,
		ListenOn:                               listenOn,
		Moniker:                                moniker,
		PublicIPv4Addr:                         publicIPv4Addr,
		RemoteURL:                              remoteURL,
		TLSCertPath:                            tlsCertPath,
		TLSKeyPath:                             tlsKeyPath,
		Type:                                   nodeType,
	}, nil
}
