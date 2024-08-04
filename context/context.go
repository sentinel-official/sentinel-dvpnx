package context

import (
	"sync"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	"github.com/sentinel-official/sentinel-go-sdk/client"
	"github.com/sentinel-official/sentinel-go-sdk/libs/geoip"
	"github.com/sentinel-official/sentinel-go-sdk/types"
	"gorm.io/gorm"
)

type Context struct {
	client  *client.Client
	db      *gorm.DB
	log     log.Logger
	service types.ServerService

	location *geoip.Location
	rpcAddr  string
	dlSpeed  sdkmath.Int
	ulSpeed  sdkmath.Int
	rw       sync.RWMutex
}

// New initializes a new Context with default values.
func New() *Context {
	return &Context{
		dlSpeed: sdkmath.ZeroInt(),
		ulSpeed: sdkmath.ZeroInt(),
	}
}

// WithClient sets the client and returns the updated Context.
func (c *Context) WithClient(client *client.Client) *Context {
	c.client = client
	return c
}

// WithDB sets the database and returns the updated Context.
func (c *Context) WithDB(db *gorm.DB) *Context {
	c.db = db
	return c
}

// WithLog sets the logger and returns the updated Context.
func (c *Context) WithLog(log log.Logger) *Context {
	c.log = log
	return c
}

// WithService sets the server service and returns the updated Context.
func (c *Context) WithService(service types.ServerService) *Context {
	c.service = service
	return c
}

// Client returns the client instance.
func (c *Context) Client() *client.Client {
	return c.client
}

// DB returns the database instance.
func (c *Context) DB() *gorm.DB {
	return c.db
}

// Log returns the logger instance.
func (c *Context) Log() log.Logger {
	return c.log
}

// Service returns the server service instance.
func (c *Context) Service() types.ServerService {
	return c.service
}

// SetLocation sets the location and returns the updated Context.
func (c *Context) SetLocation(location *geoip.Location) *Context {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.location = location
	return c
}

// SetRPCAddr sets the RPC address and returns the updated Context.
func (c *Context) SetRPCAddr(rpcAddr string) *Context {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.rpcAddr = rpcAddr
	return c
}

// SetSpeedtestResults sets download and upload speeds and returns the updated Context.
func (c *Context) SetSpeedtestResults(dlSpeed, ulSpeed sdkmath.Int) *Context {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.dlSpeed = dlSpeed
	c.ulSpeed = ulSpeed
	return c
}

// Location returns the current location.
func (c *Context) Location() *geoip.Location {
	c.rw.RLock()
	defer c.rw.RUnlock()

	return c.location
}

// RPCAddr returns the current RPC address.
func (c *Context) RPCAddr() string {
	c.rw.RLock()
	defer c.rw.RUnlock()

	return c.rpcAddr
}

// SpeedtestResults returns the current download and upload speeds.
func (c *Context) SpeedtestResults() (dlSpeed, ulSpeed sdkmath.Int) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	return c.dlSpeed, c.ulSpeed
}
