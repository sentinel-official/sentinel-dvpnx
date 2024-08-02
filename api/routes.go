package api

import (
	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/dvpn-node/api/info"
	"github.com/sentinel-official/dvpn-node/api/session"
	"github.com/sentinel-official/dvpn-node/context"
)

func RegisterRoutes(ctx *context.Context, r gin.IRouter) {
	info.RegisterRoutes(ctx, r)
	session.RegisterRoutes(ctx, r)
}
