package info

import (
	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/sentinel-dvpnx/context"
)

func RegisterRoutes(ctx *context.Context, r gin.IRouter) {
	r.GET("/", handlerGetInfo(ctx))
}
