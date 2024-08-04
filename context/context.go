package context

import (
	"sync"

	sdkmath "cosmossdk.io/math"
	"github.com/sentinel-official/sentinel-go-sdk/client"
	"github.com/sentinel-official/sentinel-go-sdk/libs/geoip"
	"github.com/sentinel-official/sentinel-go-sdk/types"
	"gorm.io/gorm"
)

type Context struct {
	client  *client.Client
	db      *gorm.DB
	service types.ServerService

	location *geoip.Location
	rpcAddr  string
	dlSpeed  sdkmath.Int
	ulSpeed  sdkmath.Int
	rw       sync.RWMutex
}

func New() *Context {
	return &Context{}
}

func (c *Context) Client() *client.Client {
	return c.client
}

func (c *Context) DB() *gorm.DB {
	return c.db
}

func (c *Context) Service() types.ServerService {
	return c.service
}

func (c *Context) SpeedtestResults() (dlSpeed, ulSpeed sdkmath.Int) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	return c.dlSpeed, c.ulSpeed
}

func (c *Context) Location() *geoip.Location {
	c.rw.RLock()
	defer c.rw.RUnlock()

	return c.location
}

func (c *Context) RPCAddr() string {
	c.rw.RLock()
	defer c.rw.RUnlock()

	return c.rpcAddr
}

func (c *Context) SetSpeedtestResults(dlSpeed, ulSpeed sdkmath.Int) error {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.dlSpeed = dlSpeed
	c.ulSpeed = ulSpeed
	return nil
}

func (c *Context) SetLocation(v *geoip.Location) error {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.location = v
	return nil
}

func (c *Context) SetRPCAddr(v string) error {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.rpcAddr = v
	return nil
}
