package node

import (
	"errors"
	"fmt"
	"time"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	sdkoptions "github.com/sentinel-official/sentinel-go-sdk/options"
	sdk "github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/dvpn-node/flags"
	"github.com/sentinel-official/dvpn-node/options"
)

// Options aggregates all the individual option structs.
type Options struct {
	*options.Handshake                     `json:"handshake" toml:"handshake"` // Options related to handshake.
	*options.Keyring                       `json:"keyring" toml:"keyring"`     // Options related to keyring configuration.
	*sdkoptions.Log                        `json:"log" toml:"log"`             // Options related to logging.
	*options.QOS                           `json:"qos" toml:"qos"`             // Options related to quality of service (QoS).
	*options.Query                         `json:"query" toml:"query"`         // Options related to querying.
	*options.Tx                            `json:"tx" toml:"tx"`               // Options related to transactions.
	GigabytePrices                         string                              `json:"gigabyte_prices" toml:"gigabyte_prices"`                                                         // GigabytePrices is the pricing information for gigabytes.
	HourlyPrices                           string                              `json:"hourly_prices" toml:"hourly_prices"`                                                             // HourlyPrices is the pricing information for hourly usage.
	IntervalBestRPCAddr                    string                              `json:"interval_best_rpc_addr" toml:"interval_best_rpc_addr"`                                           // IntervalBestRPCAddr is the duration between checking the best RPC address.
	IntervalGeoIPLocation                  string                              `json:"interval_geoip_location" toml:"interval_geoip_location"`                                         // IntervalGeoIPLocation is the duration between checking the GeoIP location.
	IntervalSessionUsageSyncWithBlockchain string                              `json:"interval_session_usage_sync_with_blockchain" toml:"interval_session_usage_sync_with_blockchain"` // IntervalSessionUsageSyncWithBlockchain is the duration between syncing session usage with the blockchain.
	IntervalSessionUsageSyncWithDatabase   string                              `json:"interval_session_usage_sync_with_database" toml:"interval_session_usage_sync_with_database"`     // IntervalSessionUsageSyncWithDatabase is the duration between syncing session usage with the database.
	IntervalSessionUsageValidate           string                              `json:"interval_session_usage_validate" toml:"interval_session_usage_validate"`                         // IntervalSessionUsageValidate is the duration between validating session usage.
	IntervalSessionValidate                string                              `json:"interval_session_validate" toml:"interval_session_validate"`                                     // IntervalSessionValidate is the duration between validating sessions.
	IntervalSpeedtest                      string                              `json:"interval_speedtest" toml:"interval_speedtest"`                                                   // IntervalSpeedtest is the duration between performing speed tests.
	IntervalStatusUpdate                   string                              `json:"interval_status_update" toml:"interval_status_update"`                                           // IntervalStatusUpdate is the duration between updating the status of the node.
	ListenOn                               string                              `json:"listen_on" toml:"listen_on"`                                                                     // ListenOn is the address on which the node listens.
	Moniker                                string                              `json:"moniker" toml:"moniker"`                                                                         // Moniker is the name or identifier for the node.
	PublicIPv4Addr                         string                              `json:"public_ipv4_addr" toml:"public_ipv4_addr"`                                                       // PublicIPv4Addr is the IPv4 address of the node.
	RemoteURL                              string                              `json:"remote_url" toml:"remote_url"`                                                                   // RemoteURL is the URL for remote operations.
	Type                                   string                              `json:"type" toml:"type"`                                                                               // Type indicates the type of node.
}

// NewOptions creates and returns a new instance of Options with default values.
func NewOptions() *Options {
	return &Options{}
}

// WithHandshake sets the Handshake option and returns the updated Options.
func (o *Options) WithHandshake(v *options.Handshake) *Options {
	o.Handshake = v
	return o
}

// WithKeyring sets the Keyring option and returns the updated Options.
func (o *Options) WithKeyring(v *options.Keyring) *Options {
	o.Keyring = v
	return o
}

// WithLog sets the Log option and returns the updated Options.
func (o *Options) WithLog(v *sdkoptions.Log) *Options {
	o.Log = v
	return o
}

// WithQOS sets the QoS option and returns the updated Options.
func (o *Options) WithQOS(v *options.QOS) *Options {
	o.QOS = v
	return o
}

// WithQuery sets the Query option and returns the updated Options.
func (o *Options) WithQuery(v *options.Query) *Options {
	o.Query = v
	return o
}

// WithTx sets the Tx option and returns the updated Options.
func (o *Options) WithTx(v *options.Tx) *Options {
	o.Tx = v
	return o
}

// WithGigabytePrices sets the GigabytePrices field and returns the modified Node instance.
func (o *Options) WithGigabytePrices(v cosmossdk.Coins) *Options {
	o.GigabytePrices = v.String()
	return o
}

// WithHourlyPrices sets the HourlyPrices field and returns the modified Node instance.
func (o *Options) WithHourlyPrices(v cosmossdk.Coins) *Options {
	o.HourlyPrices = v.String()
	return o
}

// WithIntervalBestRPCAddr sets the IntervalBestRPCAddr field and returns the modified Node instance.
func (o *Options) WithIntervalBestRPCAddr(v time.Duration) *Options {
	o.IntervalBestRPCAddr = v.String()
	return o
}

// WithIntervalGeoIPLocation sets the IntervalGeoIPLocation field and returns the modified Node instance.
func (o *Options) WithIntervalGeoIPLocation(v time.Duration) *Options {
	o.IntervalGeoIPLocation = v.String()
	return o
}

// WithIntervalSessionUsageSyncWithBlockchain sets the IntervalSessionUsageSyncWithBlockchain field and returns the modified Node instance.
func (o *Options) WithIntervalSessionUsageSyncWithBlockchain(v time.Duration) *Options {
	o.IntervalSessionUsageSyncWithBlockchain = v.String()
	return o
}

// WithIntervalSessionUsageSyncWithDatabase sets the IntervalSessionUsageSyncWithDatabase field and returns the modified Node instance.
func (o *Options) WithIntervalSessionUsageSyncWithDatabase(v time.Duration) *Options {
	o.IntervalSessionUsageSyncWithDatabase = v.String()
	return o
}

// WithIntervalSessionUsageValidate sets the IntervalSessionUsageValidate field and returns the modified Node instance.
func (o *Options) WithIntervalSessionUsageValidate(v time.Duration) *Options {
	o.IntervalSessionUsageValidate = v.String()
	return o
}

// WithIntervalSessionValidate sets the IntervalSessionValidate field and returns the modified Node instance.
func (o *Options) WithIntervalSessionValidate(v time.Duration) *Options {
	o.IntervalSessionValidate = v.String()
	return o
}

// WithIntervalSpeedtest sets the IntervalSpeedtest field and returns the modified Node instance.
func (o *Options) WithIntervalSpeedtest(v time.Duration) *Options {
	o.IntervalSpeedtest = v.String()
	return o
}

// WithIntervalStatusUpdate sets the IntervalStatusUpdate field and returns the modified Node instance.
func (o *Options) WithIntervalStatusUpdate(v time.Duration) *Options {
	o.IntervalStatusUpdate = v.String()
	return o
}

// WithListenOn sets the ListenOn field and returns the modified Node instance.
func (o *Options) WithListenOn(v string) *Options {
	o.ListenOn = v
	return o
}

// WithMoniker sets the Moniker field and returns the modified Node instance.
func (o *Options) WithMoniker(v string) *Options {
	o.Moniker = v
	return o
}

// WithPublicIPv4Addr sets the PublicIPv4Addr field and returns the modified Node instance.
func (o *Options) WithPublicIPv4Addr(v string) *Options {
	o.PublicIPv4Addr = v
	return o
}

// WithRemoteURL sets the RemoteURL field and returns the modified Node instance.
func (o *Options) WithRemoteURL(v string) *Options {
	o.RemoteURL = v
	return o
}

// WithType sets the Type field and returns the modified Node instance.
func (o *Options) WithType(v sdk.ServiceType) *Options {
	o.Type = v.String()
	return o
}

// GetGigabytePrices returns the GigabytePrices field.
func (o *Options) GetGigabytePrices() cosmossdk.Coins {
	v, err := cosmossdk.ParseCoinsNormalized(o.GigabytePrices)
	if err != nil {
		panic(err)
	}

	return v
}

// GetHourlyPrices returns the HourlyPrices field.
func (o *Options) GetHourlyPrices() cosmossdk.Coins {
	v, err := cosmossdk.ParseCoinsNormalized(o.HourlyPrices)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalBestRPCAddr returns the IntervalBestRPCAddr field.
func (o *Options) GetIntervalBestRPCAddr() time.Duration {
	v, err := time.ParseDuration(o.IntervalBestRPCAddr)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalGeoIPLocation returns the IntervalGeoIPLocation field.
func (o *Options) GetIntervalGeoIPLocation() time.Duration {
	v, err := time.ParseDuration(o.IntervalGeoIPLocation)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSessionUsageSyncWithBlockchain returns the IntervalSessionUsageSyncWithBlockchain field.
func (o *Options) GetIntervalSessionUsageSyncWithBlockchain() time.Duration {
	v, err := time.ParseDuration(o.IntervalSessionUsageSyncWithBlockchain)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSessionUsageSyncWithDatabase returns the IntervalSessionUsageSyncWithDatabase field.
func (o *Options) GetIntervalSessionUsageSyncWithDatabase() time.Duration {
	v, err := time.ParseDuration(o.IntervalSessionUsageSyncWithDatabase)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSessionUsageValidate returns the IntervalSessionUsageValidate field.
func (o *Options) GetIntervalSessionUsageValidate() time.Duration {
	v, err := time.ParseDuration(o.IntervalSessionUsageValidate)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSessionValidate returns the IntervalSessionValidate field.
func (o *Options) GetIntervalSessionValidate() time.Duration {
	v, err := time.ParseDuration(o.IntervalSessionValidate)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalSpeedtest returns the IntervalSpeedtest field.
func (o *Options) GetIntervalSpeedtest() time.Duration {
	v, err := time.ParseDuration(o.IntervalSpeedtest)
	if err != nil {
		panic(err)
	}

	return v
}

// GetIntervalStatusUpdate returns the IntervalStatusUpdate field.
func (o *Options) GetIntervalStatusUpdate() time.Duration {
	v, err := time.ParseDuration(o.IntervalStatusUpdate)
	if err != nil {
		panic(err)
	}

	return v
}

// GetListenOn returns the ListenOn field.
func (o *Options) GetListenOn() string {
	return o.ListenOn
}

// GetMoniker returns the Moniker field.
func (o *Options) GetMoniker() string {
	return o.Moniker
}

// GetPublicIPv4Addr returns the PublicIPv4Addr field.
func (o *Options) GetPublicIPv4Addr() string {
	return o.PublicIPv4Addr
}

// GetRemoteURL returns the RemoteURL field.
func (o *Options) GetRemoteURL() string {
	return o.RemoteURL
}

// GetType returns the Type field.
func (o *Options) GetType() sdk.ServiceType {
	return sdk.ServiceTypeFromString(o.Type)
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

// Validate ensures all option fields are valid.
func (o *Options) Validate() error {
	if err := o.Handshake.Validate(); err != nil {
		return fmt.Errorf("handshake validation failed: %w", err)
	}
	if err := o.Keyring.Validate(); err != nil {
		return fmt.Errorf("keyring validation failed: %w", err)
	}
	if err := o.Log.Validate(); err != nil {
		return fmt.Errorf("log validation failed: %w", err)
	}
	if err := o.QOS.Validate(); err != nil {
		return fmt.Errorf("QOS validation failed: %w", err)
	}
	if err := o.Query.Validate(); err != nil {
		return fmt.Errorf("query validation failed: %w", err)
	}
	if err := o.Tx.Validate(); err != nil {
		return fmt.Errorf("tx validation failed: %w", err)
	}

	if err := ValidateGigabytePrices(o.GigabytePrices); err != nil {
		return err
	}
	if err := ValidateHourlyPrices(o.HourlyPrices); err != nil {
		return err
	}
	if err := ValidateIntervalBestRPCAddr(o.IntervalBestRPCAddr); err != nil {
		return err
	}
	if err := ValidateIntervalGeoIPLocation(o.IntervalGeoIPLocation); err != nil {
		return err
	}
	if err := ValidateIntervalSessionUsageSyncWithBlockchain(o.IntervalSessionUsageSyncWithBlockchain); err != nil {
		return err
	}
	if err := ValidateIntervalSessionUsageSyncWithDatabase(o.IntervalSessionUsageSyncWithDatabase); err != nil {
		return err
	}
	if err := ValidateIntervalSessionUsageValidate(o.IntervalSessionUsageValidate); err != nil {
		return err
	}
	if err := ValidateIntervalSessionValidate(o.IntervalSessionValidate); err != nil {
		return err
	}
	if err := ValidateIntervalSpeedtest(o.IntervalSpeedtest); err != nil {
		return err
	}
	if err := ValidateIntervalStatusUpdate(o.IntervalStatusUpdate); err != nil {
		return err
	}
	if err := ValidateListenOn(o.ListenOn); err != nil {
		return err
	}
	if err := ValidateMoniker(o.Moniker); err != nil {
		return err
	}
	if err := ValidatePublicIPv4Addr(o.PublicIPv4Addr); err != nil {
		return err
	}
	if err := ValidateRemoteURL(o.RemoteURL); err != nil {
		return err
	}
	if err := ValidateType(o.Type); err != nil {
		return err
	}

	return nil
}

// NewOptionsFromCmd creates and returns an Options instance populated with values from the given command's flags.
func NewOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	// Retrieves Handshake from command flags.
	handshake, err := options.NewHandshakeFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Keyring from command flags.
	keyring, err := options.NewKeyringFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves LogOptions from command flags.
	log, err := sdkoptions.NewLogFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves QOS from command flags.
	qos, err := options.NewQOSFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Query from command flags.
	query, err := options.NewQueryFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Tx from command flags.
	tx, err := options.NewTxFromCmd(cmd)
	if err != nil {
		return nil, err
	}

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

	// Return a new Options instance populated with the retrieved flag values.
	return &Options{
		Handshake:                              handshake,
		Keyring:                                keyring,
		Log:                                    log,
		QOS:                                    qos,
		Query:                                  query,
		Tx:                                     tx,
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
