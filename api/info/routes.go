package info

import (
	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/dvpn-node/node"
)

func RegisterRoutes(ctx *node.Context, r gin.IRouter) {
	r.GET("/", HandlerGetInfo(ctx))
}
