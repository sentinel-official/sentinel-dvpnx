package info

import (
	"github.com/gin-gonic/gin"

	nodecontext "github.com/sentinel-official/dvpn-node/context"
)

func RegisterRoutes(ctx *nodecontext.Context, r gin.IRouter) {
	r.GET("/", HandlerGetInfo(ctx))
}
