package info

import (
	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/dvpn-node/context"
)

func RegisterRoutes(ctx *context.Context, r gin.IRouter) {
	r.GET("/", HandlerGetInfo(ctx))
}
