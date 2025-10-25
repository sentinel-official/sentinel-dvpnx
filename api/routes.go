package api

import (
	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/sentinel-dvpnx/api/handshake"
	"github.com/sentinel-official/sentinel-dvpnx/api/info"
	"github.com/sentinel-official/sentinel-dvpnx/core"
)

func RegisterRoutes(c *core.Context, r gin.IRouter) {
	handshake.RegisterRoutes(c, r)
	info.RegisterRoutes(c, r)
}
