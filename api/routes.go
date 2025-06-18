package api

import (
	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/sentinel-dvpnx/api/info"
	"github.com/sentinel-official/sentinel-dvpnx/api/session"
	"github.com/sentinel-official/sentinel-dvpnx/context"
)

func RegisterRoutes(ctx *context.Context, r gin.IRouter) {
	info.RegisterRoutes(ctx, r)
	session.RegisterRoutes(ctx, r)
}
