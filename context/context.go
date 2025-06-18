package context

import (
	"errors"
	"io"
	"sync"

	"cosmossdk.io/math"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/core"
	"github.com/sentinel-official/sentinel-go-sdk/libs/geoip"
	sentinelsdk "github.com/sentinel-official/sentinel-go-sdk/types"
	sentinelhub "github.com/sentinel-official/sentinelhub/v12/types"
	"github.com/sentinel-official/sentinelhub/v12/types/v1"
	"gorm.io/gorm"
)

// Context defines the application context, holding configurations and shared components.
type Context struct {
	accAddr        cosmossdk.AccAddress
	apiAddrs       []string
	client         *core.Client
	database       *gorm.DB
	dlSpeed        math.Int
	geoIPClient    geoip.Client
	gigabytePrices v1.Prices
	homeDir        string
	hourlyPrices   v1.Prices
	input          io.Reader
	location       *geoip.Location
	moniker        string
	remoteAddrs    []string
	rpcAddrs       []string
	service        sentinelsdk.ServerService
	ulSpeed        math.Int

	sealed bool

	fm  sync.RWMutex
	txm sync.Mutex
}

// New creates a new Context instance with default values.
func New() *Context {
	return &Context{}
}

// AccAddr returns the transaction sender address set in the context.
func (c *Context) AccAddr() cosmossdk.AccAddress {
	return c.accAddr.Bytes()
}

// APIAddrs returns the api addresses set in the context.
func (c *Context) APIAddrs() []string {
	return c.apiAddrs
}

// Client returns the client instance set in the context.
func (c *Context) Client() *core.Client {
	return c.client.WithRPCAddr(c.RPCAddr())
}

// Database returns the database connection set in the context.
func (c *Context) Database() *gorm.DB {
	return c.database
}

// GeoIPClient returns the GeoIP client set in the context.
func (c *Context) GeoIPClient() geoip.Client {
	return c.geoIPClient
}

// GigabytePrices returns the gigabyte prices for nodes.
func (c *Context) GigabytePrices() v1.Prices {
	return c.gigabytePrices
}

// HomeDir returns the home directory set in the context.
func (c *Context) HomeDir() string {
	return c.homeDir
}

// HourlyPrices returns the hourly prices for nodes.
func (c *Context) HourlyPrices() v1.Prices {
	return c.hourlyPrices
}

// Input returns the keyring input set in the context.
func (c *Context) Input() io.Reader {
	return c.input
}

// Location returns the geo-location data set in the context.
func (c *Context) Location() *geoip.Location {
	c.fm.RLock()
	defer c.fm.RUnlock()
	return c.location
}

// Moniker returns the name or identifier for the node.
func (c *Context) Moniker() string {
	return c.moniker
}

func (c *Context) NodeAddr() sentinelhub.NodeAddress {
	return c.accAddr.Bytes()
}

// RemoteAddrs returns the remote addresses set in the context.
func (c *Context) RemoteAddrs() []string {
	return c.remoteAddrs
}

// RPCAddr returns the first RPC address from the list or an empty string if no addresses are available.
func (c *Context) RPCAddr() string {
	addrs := c.RPCAddrs()
	if len(addrs) == 0 {
		return ""
	}

	return addrs[0]
}

// RPCAddrs returns the RPC addresses used for queries in the context.
func (c *Context) RPCAddrs() []string {
	c.fm.RLock()
	defer c.fm.RUnlock()
	return c.rpcAddrs
}

// Seal marks the context as sealed, preventing further modifications.
func (c *Context) Seal() *Context {
	c.sealed = true
	return c
}

// Service returns the server service instance set in the context.
func (c *Context) Service() sentinelsdk.ServerService {
	return c.service
}

// SetLocation sets the geo-location data in the context.
func (c *Context) SetLocation(location *geoip.Location) {
	c.fm.Lock()
	defer c.fm.Unlock()
	c.location = location
}

// SetRPCAddrs sets the RPC addresses in the context and allows for thread-safe updates.
func (c *Context) SetRPCAddrs(addrs []string) {
	c.fm.Lock()
	defer c.fm.Unlock()
	c.rpcAddrs = addrs
}

// SetSpeedtestResults sets the download and upload speeds in the context.
func (c *Context) SetSpeedtestResults(dlSpeed, ulSpeed math.Int) {
	c.fm.Lock()
	defer c.fm.Unlock()
	c.dlSpeed = dlSpeed
	c.ulSpeed = ulSpeed
}

// SpeedtestResults returns the download and upload speeds set in the context.
func (c *Context) SpeedtestResults() (dlSpeed, ulSpeed math.Int) {
	c.fm.RLock()
	defer c.fm.RUnlock()
	return c.dlSpeed, c.ulSpeed
}

// WithAccAddr sets the transaction sender address in the context and returns the updated context.
func (c *Context) WithAccAddr(addr cosmossdk.AccAddress) *Context {
	c.checkSealed()
	c.accAddr = addr
	return c
}

// WithAPIAddrs sets the api addresses in the context and returns the updated context.
func (c *Context) WithAPIAddrs(addrs []string) *Context {
	c.checkSealed()
	c.apiAddrs = addrs
	return c
}

// WithClient sets the base client in the context and returns the updated context.
func (c *Context) WithClient(client *core.Client) *Context {
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

// WithGigabytePrices sets the gigabyte prices for nodes and returns the updated context.
func (c *Context) WithGigabytePrices(prices v1.Prices) *Context {
	c.checkSealed()
	c.gigabytePrices = prices
	return c
}

// WithHomeDir sets the home directory in the context and returns the updated context.
func (c *Context) WithHomeDir(dir string) *Context {
	c.checkSealed()
	c.homeDir = dir
	return c
}

// WithHourlyPrices sets the hourly prices for nodes and returns the updated context.
func (c *Context) WithHourlyPrices(prices v1.Prices) *Context {
	c.checkSealed()
	c.hourlyPrices = prices
	return c
}

// WithInput sets the keyring input in the context and returns the updated context.
func (c *Context) WithInput(input io.Reader) *Context {
	c.checkSealed()
	c.input = input
	return c
}

// WithMoniker sets the name or identifier for the node and returns the updated context.
func (c *Context) WithMoniker(moniker string) *Context {
	c.checkSealed()
	c.moniker = moniker
	return c
}

// WithRemoteAddrs sets the remote addresses in the context and returns the updated context.
func (c *Context) WithRemoteAddrs(addrs []string) *Context {
	c.checkSealed()
	c.remoteAddrs = addrs
	return c
}

// WithRPCAddrs sets the RPC addresses for queries in the context and returns the updated context.
func (c *Context) WithRPCAddrs(addrs []string) *Context {
	c.checkSealed()
	c.rpcAddrs = addrs
	return c
}

// WithService sets the server service in the context and returns the updated context.
func (c *Context) WithService(service sentinelsdk.ServerService) *Context {
	c.checkSealed()
	c.service = service
	return c
}

// checkSealed verifies if the context is sealed to prevent modification.
func (c *Context) checkSealed() {
	if c.sealed {
		panic(errors.New("context is sealed and cannot be modified"))
	}
}
