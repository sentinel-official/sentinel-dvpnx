package session

import (
	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/dvpn-node/node"
)

func RegisterRoutes(ctx *node.Context, r gin.IRouter) {
	r.GET("/accounts/:acc_address/sessions/:id", HandlerAddSession(ctx))
}
