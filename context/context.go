package context

import (
	"sync"
	"time"

	"cosmossdk.io/log"
	cosmossdkmath "cosmossdk.io/math"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/client"
	"github.com/sentinel-official/sentinel-go-sdk/client/options"
	"github.com/sentinel-official/sentinel-go-sdk/libs/geoip"
	sdk "github.com/sentinel-official/sentinel-go-sdk/types"
	"gorm.io/gorm"
)

type Context struct {
	client  *client.Client
	db      *gorm.DB
	log     log.Logger
	service sdk.ServerService

	nodeGigabytePrices   cosmossdk.Coins
	nodeHourlyPrices     cosmossdk.Coins
	nodeRemoteURL        string
	queryMaxRetries      int
	queryTimeout         time.Duration
	txChainID            string
	txFeeGranterAddr     cosmossdk.AccAddress
	txFromAddr           cosmossdk.AccAddress
	txFromName           string
	txGas                uint64
	txGasAdjustment      float64
	txGasPrices          cosmossdk.DecCoins
	txSimulateAndExecute bool

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
func (c *Context) Seal() {
	c.sealed = true
}

// WithClient sets the client in the context and returns the updated context.
func (c *Context) WithClient(client *client.Client) *Context {
	c.checkSealed()
	c.client = client
	return c
}

// WithDB sets the database connection in the context and returns the updated context.
func (c *Context) WithDB(db *gorm.DB) *Context {
	c.checkSealed()
	c.db = db
	return c
}

// WithLogger sets the logger in the context and returns the updated context.
func (c *Context) WithLogger(log log.Logger) *Context {
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

// WithNodeGigabytePrices sets the gigabyte prices for nodes and returns the updated context.
func (c *Context) WithNodeGigabytePrices(prices cosmossdk.Coins) *Context {
	c.checkSealed()
	c.nodeGigabytePrices = prices
	return c
}

// WithNodeHourlyPrices sets the hourly prices for nodes and returns the updated context.
func (c *Context) WithNodeHourlyPrices(prices cosmossdk.Coins) *Context {
	c.checkSealed()
	c.nodeHourlyPrices = prices
	return c
}

// WithNodeRemoteURL sets the remote URL for nodes and returns the updated context.
func (c *Context) WithNodeRemoteURL(url string) *Context {
	c.checkSealed()
	c.nodeRemoteURL = url
	return c
}

// WithQueryMaxRetries sets the maximum number of retries for queries and returns the updated context.
func (c *Context) WithQueryMaxRetries(retries int) *Context {
	c.checkSealed()
	c.queryMaxRetries = retries
	return c
}

// WithQueryTimeout sets the timeout duration for queries and returns the updated context.
func (c *Context) WithQueryTimeout(timeout time.Duration) *Context {
	c.checkSealed()
	c.queryTimeout = timeout
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

// Client returns the client instance set in the context.
func (c *Context) Client() *client.Client {
	return c.client
}

// DB returns the database connection set in the context.
func (c *Context) DB() *gorm.DB {
	return c.db
}

// Log returns the logger instance set in the context.
func (c *Context) Log() log.Logger {
	return c.log
}

// Service returns the server service instance set in the context.
func (c *Context) Service() sdk.ServerService {
	return c.service
}

// NodeGigabytePrices returns the gigabyte prices for nodes.
func (c *Context) NodeGigabytePrices() cosmossdk.Coins {
	return c.nodeGigabytePrices
}

// NodeHourlyPrices returns the hourly prices for nodes.
func (c *Context) NodeHourlyPrices() cosmossdk.Coins {
	return c.nodeHourlyPrices
}

// NodeRemoteURL returns the remote URL for nodes.
func (c *Context) NodeRemoteURL() string {
	return c.nodeRemoteURL
}

// QueryMaxRetries returns the maximum number of retries for queries.
func (c *Context) QueryMaxRetries() int {
	return c.queryMaxRetries
}

// QueryTimeout returns the timeout duration for queries.
func (c *Context) QueryTimeout() time.Duration {
	return c.queryTimeout
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

// SetLocation sets the geo-location data in the context and returns the updated context.
func (c *Context) SetLocation(location *geoip.Location) *Context {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.location = location
	return c
}

// SetRPCAddr sets the RPC address in the context and returns the updated context.
func (c *Context) SetRPCAddr(rpcAddr string) *Context {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.rpcAddr = rpcAddr
	return c
}

// SetSpeedtestResults sets the download and upload speeds in the context and returns the updated context.
func (c *Context) SetSpeedtestResults(dlSpeed, ulSpeed cosmossdkmath.Int) *Context {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.dlSpeed = dlSpeed
	c.ulSpeed = ulSpeed
	return c
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

// ClientOptions returns an `options.Options` instance configured with the context's settings.
func (c *Context) ClientOptions() *options.Options {
	// Create default query options and configure them using the context settings.
	query := options.NewQuery().
		WithMaxRetries(c.QueryMaxRetries()).
		WithRPCAddr(c.RPCAddr()).
		WithTimeout(c.QueryTimeout())

	// Create default transaction options and configure them using the context settings.
	tx := options.NewTx().
		WithChainID(c.TxChainID()).
		WithFeeGranterAddr(c.TxFeeGranterAddr().String()).
		WithFromName(c.TxFromName()).
		WithGas(c.TxGas()).
		WithGasAdjustment(c.TxGasAdjustment()).
		WithGasPrices(c.TxGasPrices().String()).
		WithSimulateAndExecute(c.TxSimulateAndExecute())

	// Return combined options
	return options.New().
		WithQuery(query).
		WithTx(tx)
}
