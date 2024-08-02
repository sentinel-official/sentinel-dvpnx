package context

import (
	"github.com/sentinel-official/sentinel-go-sdk/client"
	"github.com/sentinel-official/sentinel-go-sdk/types"
	"gorm.io/gorm"
)

type Context struct {
	client  *client.Client
	db      *gorm.DB
	service types.ServerService
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
