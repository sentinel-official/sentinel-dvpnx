package api

import (
	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/sentinel-dvpnx/api/handshake"
	"github.com/sentinel-official/sentinel-dvpnx/api/info"
	"github.com/sentinel-official/sentinel-dvpnx/context"
)

func RegisterRoutes(ctx *context.Context, r gin.IRouter) {
	handshake.RegisterRoutes(ctx, r)
	info.RegisterRoutes(ctx, r)
}
