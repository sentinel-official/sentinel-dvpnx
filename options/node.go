package options

import (
	"errors"
	"time"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/dvpn-node/flags"
)

// Node represents options for node configuration.
type Node struct {
	GigabytePrices                         string `json:"gigabyte_prices" toml:"gigabyte_prices"`                                                         // GigabytePrices is the pricing information for gigabytes.
	HourlyPrices                           string `json:"hourly_prices" toml:"hourly_prices"`                                                             // HourlyPrices is the pricing information for hourly usage.
	IntervalBestRPCAddr                    string `json:"interval_best_rpc_addr" toml:"interval_best_rpc_addr"`                                           // IntervalBestRPCAddr is the duration between checking the best RPC address.
	IntervalGeoIPLocation                  string `json:"interval_geoip_location" toml:"interval_geoip_location"`                                         // IntervalGeoIPLocation is the duration between checking the GeoIP location.
	IntervalSessionUsageSyncWithBlockchain string `json:"interval_session_usage_sync_with_blockchain" toml:"interval_session_usage_sync_with_blockchain"` // IntervalSessionUsageSyncWithBlockchain is the duration between syncing session usage with the blockchain.
	IntervalSessionUsageSyncWithDatabase   string `json:"interval_session_usage_sync_with_database" toml:"interval_session_usage_sync_with_database"`     // IntervalSessionUsageSyncWithDatabase is the duration between syncing session usage with the database.
	IntervalSessionUsageValidate           string `json:"interval_session_usage_validate" toml:"interval_session_usage_validate"`                         // IntervalSessionUsageValidate is the duration between validating session usage.
	IntervalSessionValidate                string `json:"interval_session_validate" toml:"interval_session_validate"`                                     // IntervalSessionValidate is the duration between validating sessions.
	IntervalSpeedtest                      string `json:"interval_speedtest" toml:"interval_speedtest"`                                                   // IntervalSpeedtest is the duration between performing speed tests.
	IntervalStatusUpdate                   string `json:"interval_status_update" toml:"interval_status_update"`                                           // IntervalStatusUpdate is the duration between updating the status of the node.
	ListenOn                               string `json:"listen_on" toml:"listen_on"`                                                                     // ListenOn is the address on which the node listens.
	Moniker                                string `json:"moniker" toml:"moniker"`                                                                         // Moniker is the name or identifier for the node.
	PublicIPv4Addr                         string `json:"public_ipv4_addr" toml:"public_ipv4_addr"`                                                       // PublicIPv4Addr is the IPv4 address of the node.
	RemoteURL                              string `json:"remote_url" toml:"remote_url"`                                                                   // RemoteURL is the URL for remote operations.
	Type                                   string `json:"type" toml:"type"`                                                                               // Type indicates the type of node.
}

// NewNode creates a new Node instance with default values.
func NewNode() *Node {
	return &Node{}
}

// WithGigabytePrices sets the GigabytePrices field and returns the modified Node instance.
func (n *Node) WithGigabytePrices(v cosmossdk.Coins) *Node {
	n.GigabytePrices = v.String()
	return n
}

// WithHourlyPrices sets the HourlyPrices field and returns the modified Node instance.
func (n *Node) WithHourlyPrices(v cosmossdk.Coins) *Node {
	n.HourlyPrices = v.String()
	return n
}

// WithIntervalBestRPCAddr sets the IntervalBestRPCAddr field and returns the modified Node instance.
func (n *Node) WithIntervalBestRPCAddr(v time.Duration) *Node {
	n.IntervalBestRPCAddr = v.String()
	return n
}

// WithIntervalGeoIPLocation sets the IntervalGeoIPLocation field and returns the modified Node instance.
func (n *Node) WithIntervalGeoIPLocation(v time.Duration) *Node {
	n.IntervalGeoIPLocation = v.String()
	return n
}

// WithIntervalSessionUsageSyncWithBlockchain sets the IntervalSessionUsageSyncWithBlockchain field and returns the modified Node instance.
func (n *Node) WithIntervalSessionUsageSyncWithBlockchain(v time.Duration) *Node {
	n.IntervalSessionUsageSyncWithBlockchain = v.String()
	return n
}

// WithIntervalSessionUsageSyncWithDatabase sets the IntervalSessionUsageSyncWithDatabase field and returns the modified Node instance.
func (n *Node) WithIntervalSessionUsageSyncWithDatabase(v time.Duration) *Node {
	n.IntervalSessionUsageSyncWithDatabase = v.String()
	return n
}

// WithIntervalSessionUsageValidate sets the IntervalSessionUsageValidate field and returns the modified Node instance.
func (n *Node) WithIntervalSessionUsageValidate(v time.Duration) *Node {
	n.IntervalSessionUsageValidate = v.String()
	return n
}

// WithIntervalSessionValidate sets the IntervalSessionValidate field and returns the modified Node instance.
func (n *Node) WithIntervalSessionValidate(v time.Duration) *Node {
	n.IntervalSessionValidate = v.String()
	return n
}

// WithIntervalSpeedtest sets the IntervalSpeedtest field and returns the modified Node instance.
func (n *Node) WithIntervalSpeedtest(v time.Duration) *Node {
	n.IntervalSpeedtest = v.String()
	return n
}

// WithIntervalStatusUpdate sets the IntervalStatusUpdate field and returns the modified Node instance.
func (n *Node) WithIntervalStatusUpdate(v time.Duration) *Node {
	n.IntervalStatusUpdate = v.String()
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
	n.Type = v.String()
	return n
}

// GetGigabytePrices returns the GigabytePrices field.
func (n *Node) GetGigabytePrices() cosmossdk.Coins {
	v, err := cosmossdk.ParseCoinsNormalized(n.GigabytePrices)
	if err != nil {
		panic(err)
	}

	return v
}

// GetHourlyPrices returns the HourlyPrices field.
func (n *Node) GetHourlyPrices() cosmossdk.Coins {
	v, err := cosmossdk.ParseCoinsNormalized(n.HourlyPrices)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalBestRPCAddr returns the IntervalBestRPCAddr field.
func (n *Node) GetIntervalBestRPCAddr() time.Duration {
	v, err := time.ParseDuration(n.IntervalBestRPCAddr)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalGeoIPLocation returns the IntervalGeoIPLocation field.
func (n *Node) GetIntervalGeoIPLocation() time.Duration {
	v, err := time.ParseDuration(n.IntervalGeoIPLocation)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSessionUsageSyncWithBlockchain returns the IntervalSessionUsageSyncWithBlockchain field.
func (n *Node) GetIntervalSessionUsageSyncWithBlockchain() time.Duration {
	v, err := time.ParseDuration(n.IntervalSessionUsageSyncWithBlockchain)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSessionUsageSyncWithDatabase returns the IntervalSessionUsageSyncWithDatabase field.
func (n *Node) GetIntervalSessionUsageSyncWithDatabase() time.Duration {
	v, err := time.ParseDuration(n.IntervalSessionUsageSyncWithDatabase)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSessionUsageValidate returns the IntervalSessionUsageValidate field.
func (n *Node) GetIntervalSessionUsageValidate() time.Duration {
	v, err := time.ParseDuration(n.IntervalSessionUsageValidate)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSessionValidate returns the IntervalSessionValidate field.
func (n *Node) GetIntervalSessionValidate() time.Duration {
	v, err := time.ParseDuration(n.IntervalSessionValidate)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSpeedtest returns the IntervalSpeedtest field.
func (n *Node) GetIntervalSpeedtest() time.Duration {
	v, err := time.ParseDuration(n.IntervalSpeedtest)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalStatusUpdate returns the IntervalStatusUpdate field.
func (n *Node) GetIntervalStatusUpdate() time.Duration {
	v, err := time.ParseDuration(n.IntervalStatusUpdate)
	if err != nil {
		panic(err)
	}

	return v
}

// GetListenOn returns the ListenOn field.
func (n *Node) GetListenOn() string {
	return n.ListenOn
}

// GetMoniker returns the Moniker field.
func (n *Node) GetMoniker() string {
	return n.Moniker
}

// GetPublicIPv4Addr returns the PublicIPv4Addr field.
func (n *Node) GetPublicIPv4Addr() string {
	return n.PublicIPv4Addr
}

// GetRemoteURL returns the RemoteURL field.
func (n *Node) GetRemoteURL() string {
	return n.RemoteURL
}

// GetType returns the Type field.
func (n *Node) GetType() sdk.ServiceType {
	return sdk.ServiceTypeFromString(n.Type)
}

// ValidateGigabytePrices checks if the GigabytePrices field is valid.
func ValidateGigabytePrices(v string) error {
	if _, err := cosmossdk.ParseCoinsNormalized(v); err != nil {
		return errors.New("gigabyte prices must be a valid coins format")
	}

	return nil
}

// ValidateHourlyPrices checks if the HourlyPrices field is valid.
func ValidateHourlyPrices(v string) error {
	if _, err := cosmossdk.ParseCoinsNormalized(v); err != nil {
		return errors.New("hourly prices must be a valid coins format")
	}

	return nil
}

// ValidateIntervalBestRPCAddr checks if the IntervalBestRPCAddr field is valid.
func ValidateIntervalBestRPCAddr(v string) error {
	if _, err := time.ParseDuration(v); err != nil {
		return errors.New("invalid duration format for interval best RPC address")
	}

	return nil
}

// ValidateIntervalGeoIPLocation checks if the IntervalGeoIPLocation field is valid.
func ValidateIntervalGeoIPLocation(v string) error {
	if _, err := time.ParseDuration(v); err != nil {
		return errors.New("invalid duration format for interval GeoIP location")
	}

	return nil
}

// ValidateIntervalSessionUsageSyncWithBlockchain checks if the IntervalSessionUsageSyncWithBlockchain field is valid.
func ValidateIntervalSessionUsageSyncWithBlockchain(v string) error {
	if _, err := time.ParseDuration(v); err != nil {
		return errors.New("invalid duration format for interval session usage sync with blockchain")
	}

	return nil
}

// ValidateIntervalSessionUsageSyncWithDatabase checks if the IntervalSessionUsageSyncWithDatabase field is valid.
func ValidateIntervalSessionUsageSyncWithDatabase(v string) error {
	if _, err := time.ParseDuration(v); err != nil {
		return errors.New("invalid duration format for interval session usage sync with database")
	}

	return nil
}

// ValidateIntervalSessionUsageValidate checks if the IntervalSessionUsageValidate field is valid.
func ValidateIntervalSessionUsageValidate(v string) error {
	if _, err := time.ParseDuration(v); err != nil {
		return errors.New("invalid duration format for interval session usage validate")
	}

	return nil
}

// ValidateIntervalSessionValidate checks if the IntervalSessionValidate field is valid.
func ValidateIntervalSessionValidate(v string) error {
	if _, err := time.ParseDuration(v); err != nil {
		return errors.New("invalid duration format for interval session validate")
	}

	return nil
}

// ValidateIntervalSpeedtest checks if the IntervalSpeedtest field is valid.
func ValidateIntervalSpeedtest(v string) error {
	if _, err := time.ParseDuration(v); err != nil {
		return errors.New("invalid duration format for interval speedtest")
	}

	return nil
}

// ValidateIntervalStatusUpdate checks if the IntervalStatusUpdate field is valid.
func ValidateIntervalStatusUpdate(v string) error {
	if _, err := time.ParseDuration(v); err != nil {
		return errors.New("invalid duration format for interval status update")
	}

	return nil
}

// ValidateListenOn checks if the ListenOn field is valid.
func ValidateListenOn(v string) error {
	if v == "" {
		return errors.New("listen-on address must be non-empty")
	}

	return nil
}

// ValidateMoniker checks if the Moniker field is valid.
func ValidateMoniker(v string) error {
	if v == "" {
		return errors.New("moniker must be non-empty")
	}

	return nil
}

// ValidatePublicIPv4Addr checks if the PublicIPv4Addr field is valid.
func ValidatePublicIPv4Addr(v string) error {
	if v == "" {
		return errors.New("public IPv4 address must be non-empty")
	}

	return nil
}

// ValidateRemoteURL checks if the RemoteURL field is valid.
func ValidateRemoteURL(v string) error {
	if v == "" {
		return errors.New("remote URL must be non-empty")
	}

	return nil
}

// ValidateType checks if the Type field is valid.
func ValidateType(v string) error {
	if v == "" {
		return errors.New("type must be non-empty")
	}

	return nil
}

// Validate checks if all fields in the Node struct are valid.
func (n *Node) Validate() error {
	if err := ValidateGigabytePrices(n.GigabytePrices); err != nil {
		return err
	}
	if err := ValidateHourlyPrices(n.HourlyPrices); err != nil {
		return err
	}
	if err := ValidateIntervalBestRPCAddr(n.IntervalBestRPCAddr); err != nil {
		return err
	}
	if err := ValidateIntervalGeoIPLocation(n.IntervalGeoIPLocation); err != nil {
		return err
	}
	if err := ValidateIntervalSessionUsageSyncWithBlockchain(n.IntervalSessionUsageSyncWithBlockchain); err != nil {
		return err
	}
	if err := ValidateIntervalSessionUsageSyncWithDatabase(n.IntervalSessionUsageSyncWithDatabase); err != nil {
		return err
	}
	if err := ValidateIntervalSessionUsageValidate(n.IntervalSessionUsageValidate); err != nil {
		return err
	}
	if err := ValidateIntervalSessionValidate(n.IntervalSessionValidate); err != nil {
		return err
	}
	if err := ValidateIntervalSpeedtest(n.IntervalSpeedtest); err != nil {
		return err
	}
	if err := ValidateIntervalStatusUpdate(n.IntervalStatusUpdate); err != nil {
		return err
	}
	if err := ValidateListenOn(n.ListenOn); err != nil {
		return err
	}
	if err := ValidateMoniker(n.Moniker); err != nil {
		return err
	}
	if err := ValidatePublicIPv4Addr(n.PublicIPv4Addr); err != nil {
		return err
	}
	if err := ValidateRemoteURL(n.RemoteURL); err != nil {
		return err
	}
	if err := ValidateType(n.Type); err != nil {
		return err
	}

	return nil
}

// NewNodeFromCmd creates and returns Node from the given Cobra command's flags.
func NewNodeFromCmd(cmd *cobra.Command) (*Node, error) {
	// Retrieve the gigabyte prices flag value from the command.
	gigabytePrices, err := flags.GetNodeGigabytePrices(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the hourly prices flag value from the command.
	hourlyPrices, err := flags.GetNodeHourlyPrices(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval best RPC address flag value from the command.
	intervalBestRPCAddr, err := flags.GetNodeIntervalBestRPCAddr(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval GeoIP location flag value from the command.
	intervalGeoIPLocation, err := flags.GetNodeIntervalGeoIPLocation(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval session usage sync with blockchain flag value from the command.
	intervalSessionUsageSyncWithBlockchain, err := flags.GetNodeIntervalSessionUsageSyncWithBlockchain(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval session usage sync with database flag value from the command.
	intervalSessionUsageSyncWithDatabase, err := flags.GetNodeIntervalSessionUsageSyncWithDatabase(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval session usage validate flag value from the command.
	intervalSessionUsageValidate, err := flags.GetNodeIntervalSessionUsageValidate(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval session validate flag value from the command.
	intervalSessionValidate, err := flags.GetNodeIntervalSessionValidate(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval speedtest flag value from the command.
	intervalSpeedtest, err := flags.GetNodeIntervalSpeedtest(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the interval status update flag value from the command.
	intervalStatusUpdate, err := flags.GetNodeIntervalStatusUpdate(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the listen-on flag value from the command.
	listenOn, err := flags.GetNodeListenOn(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the moniker flag value from the command.
	moniker, err := flags.GetNodeMoniker(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the public IPv4 address flag value from the command.
	publicIPv4Addr, err := flags.GetNodePublicIPv4Addr(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the remote URL flag value from the command.
	remoteURL, err := flags.GetNodeRemoteURL(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the node type flag value from the command.
	nodeType, err := flags.GetNodeType(cmd)
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
		Type:                                   nodeType,
	}, nil
}
