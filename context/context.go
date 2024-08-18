package context

import (
	"path/filepath"
	"sync"
	"time"

	cosmossdklog "cosmossdk.io/log"
	cosmossdkmath "cosmossdk.io/math"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	sdkclient "github.com/sentinel-official/sentinel-go-sdk/client"
	"github.com/sentinel-official/sentinel-go-sdk/libs/geoip"
	sdkoptions "github.com/sentinel-official/sentinel-go-sdk/options"
	sdk "github.com/sentinel-official/sentinel-go-sdk/types"
	"gorm.io/gorm"
)

type Context struct {
	appName     string
	client      *sdkclient.Client
	database    *gorm.DB
	geoIPClient geoip.Client
	homeDir     string
	log         cosmossdklog.Logger
	service     sdk.ServerService

	gigabytePrices                         cosmossdk.Coins
	hourlyPrices                           cosmossdk.Coins
	intervalBestRPCAddr                    time.Duration
	intervalGeoIPLocation                  time.Duration
	intervalSessionUsageSyncWithBlockchain time.Duration
	intervalSessionUsageSyncWithDatabase   time.Duration
	intervalSessionUsageValidate           time.Duration
	intervalSessionValidate                time.Duration
	intervalSpeedtest                      time.Duration
	intervalStatusUpdate                   time.Duration
	keyringBackend                         string
	listenOn                               string
	moniker                                string
	nodeType                               sdk.ServiceType
	queryMaxRetries                        int
	queryRPCAddrs                          []string
	queryTimeout                           time.Duration
	remoteURL                              string
	txChainID                              string
	txFeeGranterAddr                       cosmossdk.AccAddress
	txFromAddr                             cosmossdk.AccAddress
	txFromName                             string
	txGasAdjustment                        float64
	txGasPrices                            cosmossdk.DecCoins
	txGas                                  uint64
	txSimulateAndExecute                   bool

	rw       sync.RWMutex
	location *geoip.Location
	rpcAddr  string
	dlSpeed  cosmossdkmath.Int
	ulSpeed  cosmossdkmath.Int

	sealed bool
}

// New creates a new Context instance with default values.
func New() *Context {
	return &Context{
		dlSpeed: cosmossdkmath.ZeroInt(),
		ulSpeed: cosmossdkmath.ZeroInt(),
	}
}

// checkSealed verifies if the context is sealed to prevent modification.
func (c *Context) checkSealed() {
	if c.sealed {
		panic("context is sealed and cannot be modified")
	}
}

// Seal marks the context as sealed, preventing further modifications.
func (c *Context) Seal() *Context {
	c.sealed = true
	return c
}

// WithAppName sets the application name in the context and returns the updated context.
func (c *Context) WithAppName(appName string) *Context {
	c.checkSealed()
	c.appName = appName
	return c
}

// WithClient sets the client in the context and returns the updated context.
func (c *Context) WithClient(client *sdkclient.Client) *Context {
	c.checkSealed()
	c.client = client
	return c
}

// WithDatabase sets the database connection in the context and returns the updated context.
func (c *Context) WithDatabase(database *gorm.DB) *Context {
	c.checkSealed()
	c.database = database
	return c
}

// WithGeoIPClient sets the GeoIP client in the context and returns the updated context.
func (c *Context) WithGeoIPClient(client geoip.Client) *Context {
	c.checkSealed()
	c.geoIPClient = client
	return c
}

// WithHomeDir sets the home directory in the context and returns the updated context.
func (c *Context) WithHomeDir(dir string) *Context {
	c.checkSealed()
	c.homeDir = dir
	return c
}

// WithLogger sets the logger in the context and returns the updated context.
func (c *Context) WithLogger(log cosmossdklog.Logger) *Context {
	c.checkSealed()
	c.log = log
	return c
}

// WithService sets the server service in the context and returns the updated context.
func (c *Context) WithService(service sdk.ServerService) *Context {
	c.checkSealed()
	c.service = service
	return c
}

// WithGigabytePrices sets the gigabyte prices for nodes and returns the updated context.
func (c *Context) WithGigabytePrices(prices cosmossdk.Coins) *Context {
	c.checkSealed()
	c.gigabytePrices = prices
	return c
}

// WithHourlyPrices sets the hourly prices for nodes and returns the updated context.
func (c *Context) WithHourlyPrices(prices cosmossdk.Coins) *Context {
	c.checkSealed()
	c.hourlyPrices = prices
	return c
}

// WithIntervalBestRPCAddr sets the interval for the best RPC address check and returns the updated context.
func (c *Context) WithIntervalBestRPCAddr(interval time.Duration) *Context {
	c.checkSealed()
	c.intervalBestRPCAddr = interval
	return c
}

// WithIntervalGeoIPLocation sets the interval for GeoIP location updates and returns the updated context.
func (c *Context) WithIntervalGeoIPLocation(interval time.Duration) *Context {
	c.checkSealed()
	c.intervalGeoIPLocation = interval
	return c
}

// WithIntervalSessionUsageSyncWithBlockchain sets the interval for syncing session usage with the blockchain and returns the updated context.
func (c *Context) WithIntervalSessionUsageSyncWithBlockchain(interval time.Duration) *Context {
	c.checkSealed()
	c.intervalSessionUsageSyncWithBlockchain = interval
	return c
}

// WithIntervalSessionUsageSyncWithDatabase sets the interval for syncing session usage with the database and returns the updated context.
func (c *Context) WithIntervalSessionUsageSyncWithDatabase(interval time.Duration) *Context {
	c.checkSealed()
	c.intervalSessionUsageSyncWithDatabase = interval
	return c
}

// WithIntervalSessionUsageValidate sets the interval for validating session usage and returns the updated context.
func (c *Context) WithIntervalSessionUsageValidate(interval time.Duration) *Context {
	c.checkSealed()
	c.intervalSessionUsageValidate = interval
	return c
}

// WithIntervalSessionValidate sets the interval for validating sessions and returns the updated context.
func (c *Context) WithIntervalSessionValidate(interval time.Duration) *Context {
	c.checkSealed()
	c.intervalSessionValidate = interval
	return c
}

// WithIntervalSpeedtest sets the interval for running speed tests and returns the updated context.
func (c *Context) WithIntervalSpeedtest(interval time.Duration) *Context {
	c.checkSealed()
	c.intervalSpeedtest = interval
	return c
}

// WithIntervalStatusUpdate sets the interval for updating the node status and returns the updated context.
func (c *Context) WithIntervalStatusUpdate(interval time.Duration) *Context {
	c.checkSealed()
	c.intervalStatusUpdate = interval
	return c
}

// WithKeyringBackend sets the keyring backend in the context and returns the updated context.
func (c *Context) WithKeyringBackend(backend string) *Context {
	c.checkSealed()
	c.keyringBackend = backend
	return c
}

// WithListenOn sets the address or port the node should listen on and returns the updated context.
func (c *Context) WithListenOn(address string) *Context {
	c.checkSealed()
	c.listenOn = address
	return c
}

// WithMoniker sets the name or identifier for the node and returns the updated context.
func (c *Context) WithMoniker(moniker string) *Context {
	c.checkSealed()
	c.moniker = moniker
	return c
}

// WithNodeType sets the type or role of the node and returns the updated context.
func (c *Context) WithNodeType(nodeType sdk.ServiceType) *Context {
	c.checkSealed()
	c.nodeType = nodeType
	return c
}

// WithQueryMaxRetries sets the maximum number of retries for queries and returns the updated context.
func (c *Context) WithQueryMaxRetries(retries int) *Context {
	c.checkSealed()
	c.queryMaxRetries = retries
	return c
}

// WithQueryRPCAddrs sets the RPC addresses for queries in the context and returns the updated context.
func (c *Context) WithQueryRPCAddrs(addrs []string) *Context {
	c.checkSealed()
	c.queryRPCAddrs = addrs
	return c
}

// WithQueryTimeout sets the timeout duration for queries and returns the updated context.
func (c *Context) WithQueryTimeout(timeout time.Duration) *Context {
	c.checkSealed()
	c.queryTimeout = timeout
	return c
}

// WithRemoteURL sets the remote URL for nodes and returns the updated context.
func (c *Context) WithRemoteURL(url string) *Context {
	c.checkSealed()
	c.remoteURL = url
	return c
}

// WithTxChainID sets the chain ID for transactions and returns the updated context.
func (c *Context) WithTxChainID(chainID string) *Context {
	c.checkSealed()
	c.txChainID = chainID
	return c
}

// WithTxFeeGranterAddr sets the address of the transaction fee granter and returns the updated context.
func (c *Context) WithTxFeeGranterAddr(addr cosmossdk.AccAddress) *Context {
	c.checkSealed()
	c.txFeeGranterAddr = addr
	return c
}

// WithTxFromAddr sets the address of the account sending the transaction and returns the updated context.
func (c *Context) WithTxFromAddr(addr cosmossdk.AccAddress) *Context {
	c.checkSealed()
	c.txFromAddr = addr
	return c
}

// WithTxFromName sets the name of the account sending the transaction and returns the updated context.
func (c *Context) WithTxFromName(name string) *Context {
	c.checkSealed()
	c.txFromName = name
	return c
}

// WithTxGas sets the gas limit for transactions and returns the updated context.
func (c *Context) WithTxGas(gas uint64) *Context {
	c.checkSealed()
	c.txGas = gas
	return c
}

// WithTxGasAdjustment sets the gas adjustment factor for transactions and returns the updated context.
func (c *Context) WithTxGasAdjustment(adjustment float64) *Context {
	c.checkSealed()
	c.txGasAdjustment = adjustment
	return c
}

// WithTxGasPrices sets the gas prices for transactions and returns the updated context.
func (c *Context) WithTxGasPrices(prices cosmossdk.DecCoins) *Context {
	c.checkSealed()
	c.txGasPrices = prices
	return c
}

// WithTxSimulateAndExecute sets the flag for simulating and executing transactions and returns the updated context.
func (c *Context) WithTxSimulateAndExecute(simulate bool) *Context {
	c.checkSealed()
	c.txSimulateAndExecute = simulate
	return c
}

// SetLocation sets the geo-location data in the context.
func (c *Context) SetLocation(location *geoip.Location) {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.location = location
}

// SetRPCAddr sets the RPC address in the context.
func (c *Context) SetRPCAddr(rpcAddr string) {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.rpcAddr = rpcAddr
}

// SetSpeedtestResults sets the download and upload speeds in the context.
func (c *Context) SetSpeedtestResults(dlSpeed, ulSpeed cosmossdkmath.Int) {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.dlSpeed = dlSpeed
	c.ulSpeed = ulSpeed
}

// AppName returns the application name set in the context.
func (c *Context) AppName() string {
	return c.appName
}

// Client returns the client instance set in the context.
func (c *Context) Client() *sdkclient.Client {
	return c.client
}

// Database returns the database connection set in the context.
func (c *Context) Database() *gorm.DB {
	return c.database
}

// DatabaseFilePath returns the file path for the database file.
func (c *Context) DatabaseFilePath() string {
	return filepath.Join(c.homeDir, "data.db")
}

// GeoIPClient returns the GeoIP client set in the context.
func (c *Context) GeoIPClient() geoip.Client {
	return c.geoIPClient
}

// KeyringDir returns the directory path where the keyring is located.
func (c *Context) KeyringDir() string {
	return c.homeDir
}

// Log returns the logger instance set in the context.
func (c *Context) Log() cosmossdklog.Logger {
	return c.log
}

// Service returns the server service instance set in the context.
func (c *Context) Service() sdk.ServerService {
	return c.service
}

// TLSCertPath returns the path to the TLS certificate for secure communication.
func (c *Context) TLSCertPath() string {
	return filepath.Join(c.homeDir, "tls.crt")
}

// TLSKeyPath returns the path to the TLS key for secure communication.
func (c *Context) TLSKeyPath() string {
	return filepath.Join(c.homeDir, "tls.key")
}

// GigabytePrices returns the gigabyte prices for nodes.
func (c *Context) GigabytePrices() cosmossdk.Coins {
	return c.gigabytePrices
}

// HourlyPrices returns the hourly prices for nodes.
func (c *Context) HourlyPrices() cosmossdk.Coins {
	return c.hourlyPrices
}

// IntervalBestRPCAddr returns the interval for the best RPC address check.
func (c *Context) IntervalBestRPCAddr() time.Duration {
	return c.intervalBestRPCAddr
}

// IntervalGeoIPLocation returns the interval for GeoIP location updates.
func (c *Context) IntervalGeoIPLocation() time.Duration {
	return c.intervalGeoIPLocation
}

// IntervalSessionUsageSyncWithBlockchain returns the interval for syncing session usage with the blockchain.
func (c *Context) IntervalSessionUsageSyncWithBlockchain() time.Duration {
	return c.intervalSessionUsageSyncWithBlockchain
}

// IntervalSessionUsageSyncWithDatabase returns the interval for syncing session usage with the database.
func (c *Context) IntervalSessionUsageSyncWithDatabase() time.Duration {
	return c.intervalSessionUsageSyncWithDatabase
}

// IntervalSessionUsageValidate returns the interval for validating session usage.
func (c *Context) IntervalSessionUsageValidate() time.Duration {
	return c.intervalSessionUsageValidate
}

// IntervalSessionValidate returns the interval for validating sessions.
func (c *Context) IntervalSessionValidate() time.Duration {
	return c.intervalSessionValidate
}

// IntervalSpeedtest returns the interval for running speed tests.
func (c *Context) IntervalSpeedtest() time.Duration {
	return c.intervalSpeedtest
}

// IntervalStatusUpdate returns the interval for updating the node status.
func (c *Context) IntervalStatusUpdate() time.Duration {
	return c.intervalStatusUpdate
}

// KeyringBackend returns the keyring backend set in the context.
func (c *Context) KeyringBackend() string {
	return c.keyringBackend
}

// ListenOn returns the address or port the node is listening on.
func (c *Context) ListenOn() string {
	return c.listenOn
}

// Moniker returns the name or identifier for the node.
func (c *Context) Moniker() string {
	return c.moniker
}

// NodeType returns the type or role of the node.
func (c *Context) NodeType() sdk.ServiceType {
	return c.nodeType
}

// QueryMaxRetries returns the maximum number of retries for queries.
func (c *Context) QueryMaxRetries() int {
	return c.queryMaxRetries
}

// QueryRPCAddrs returns the RPC addresses used for queries in the context.
func (c *Context) QueryRPCAddrs() []string {
	return c.queryRPCAddrs
}

// QueryTimeout returns the timeout duration for queries.
func (c *Context) QueryTimeout() time.Duration {
	return c.queryTimeout
}

// RemoteURL returns the remote URL for nodes.
func (c *Context) RemoteURL() string {
	return c.remoteURL
}

// TxChainID returns the chain ID for transactions.
func (c *Context) TxChainID() string {
	return c.txChainID
}

// TxFeeGranterAddr returns the address of the transaction fee granter.
func (c *Context) TxFeeGranterAddr() cosmossdk.AccAddress {
	return c.txFeeGranterAddr
}

// TxFromAddr returns the address of the account sending the transaction.
func (c *Context) TxFromAddr() cosmossdk.AccAddress {
	return c.txFromAddr
}

// TxFromName returns the name of the account sending the transaction.
func (c *Context) TxFromName() string {
	return c.txFromName
}

// TxGas returns the gas limit for transactions.
func (c *Context) TxGas() uint64 {
	return c.txGas
}

// TxGasAdjustment returns the gas adjustment factor for transactions.
func (c *Context) TxGasAdjustment() float64 {
	return c.txGasAdjustment
}

// TxGasPrices returns the gas prices for transactions.
func (c *Context) TxGasPrices() cosmossdk.DecCoins {
	return c.txGasPrices
}

// TxSimulateAndExecute returns the flag indicating if transactions should be simulated and executed.
func (c *Context) TxSimulateAndExecute() bool {
	return c.txSimulateAndExecute
}

// Location returns the geo-location data set in the context.
func (c *Context) Location() *geoip.Location {
	c.rw.RLock()
	defer c.rw.RUnlock()

	return c.location
}

// RPCAddr returns the RPC address set in the context.
func (c *Context) RPCAddr() string {
	c.rw.RLock()
	defer c.rw.RUnlock()

	return c.rpcAddr
}

// SpeedtestResults returns the download and upload speeds set in the context.
func (c *Context) SpeedtestResults() (dlSpeed, ulSpeed cosmossdkmath.Int) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	return c.dlSpeed, c.ulSpeed
}

// ClientOptions returns an `sdkclient.Options` instance configured with the context's settings.
func (c *Context) ClientOptions() *sdkclient.Options {
	// Create default query options and configure them using the context settings.
	query := sdkoptions.NewQuery().
		WithMaxRetries(c.QueryMaxRetries()).
		WithRPCAddr(c.RPCAddr()).
		WithTimeout(c.QueryTimeout())

	// Create default transaction options and configure them using the context settings.
	tx := sdkoptions.NewTx().
		WithChainID(c.TxChainID()).
		WithFeeGranterAddr(c.TxFeeGranterAddr()).
		WithFromName(c.TxFromName()).
		WithGas(c.TxGas()).
		WithGasAdjustment(c.TxGasAdjustment()).
		WithGasPrices(c.TxGasPrices()).
		WithSimulateAndExecute(c.TxSimulateAndExecute())

	// Return combined options.
	return sdkclient.NewOptions().
		WithQuery(query).
		WithTx(tx)
}
