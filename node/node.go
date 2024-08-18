package node

import (
	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"

	nodecontext "github.com/sentinel-official/dvpn-node/context"
	"github.com/sentinel-official/dvpn-node/utils"
)

// Node represents the application node, holding its context, router, and scheduler.
type Node struct {
	ctx       *nodecontext.Context
	router    *gin.Engine
	scheduler *cron.Scheduler
}

// New creates a new Node with the provided context.
func New(ctx *nodecontext.Context) *Node {
	return &Node{
		ctx: ctx,
	}
}

// WithRouter sets the router for the Node and returns the updated Node.
func (n *Node) WithRouter(v *gin.Engine) *Node {
	n.router = v
	return n
}

// WithScheduler sets the scheduler for the Node and returns the updated Node.
func (n *Node) WithScheduler(v *cron.Scheduler) *Node {
	n.scheduler = v
	return n
}

// Start starts the scheduler and the Node server.
func (n *Node) Start() error {
	// Start the cron scheduler.
	if err := n.scheduler.Start(); err != nil {
		return err
	}

	// Start the HTTPS server using the configured TLS certificates.
	return utils.ListenAndServeTLS(
		n.ctx.ListenOn(),
		n.ctx.TLSCertPath(),
		n.ctx.TLSKeyPath(),
		n.router,
	)
}

// Stop stops the Node.
func (n *Node) Stop() error {
	return nil
}
