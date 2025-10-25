package handshake

import (
	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/sentinel-dvpnx/core"
)

// RegisterRoutes registers the routes for the handshake API.
func RegisterRoutes(c *core.Context, r gin.IRouter) {
	r.POST("/", handlerInitHandshake(c))
}
