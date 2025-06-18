package session

import (
	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/sentinel-dvpnx/context"
)

// RegisterRoutes registers the routes for the session API.
func RegisterRoutes(ctx *context.Context, r gin.IRouter) {
	r.POST("/sessions", handlerAddSession(ctx))
}
