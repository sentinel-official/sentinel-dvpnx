package node

import (
	gocontext "context"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cmux"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	"github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types/v3"

	"github.com/sentinel-official/sentinel-dvpnx/context"
)

// Node represents the application node, holding its context, router, and scheduler.
type Node struct {
	*context.Context
	laddr     string          // Address the Node listens on for incoming requests.
	router    *gin.Engine     // HTTP router for handling API requests.
	scheduler *cron.Scheduler // Scheduler for managing periodic tasks.
}

// New creates a new Node with the provided context.
func New(ctx *context.Context) *Node {
	return &Node{Context: ctx}
}

// WithListenAddr sets the listen address for the Node and returns the updated Node.
func (n *Node) WithListenAddr(v string) *Node {
	n.laddr = v
	return n
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

// ListenAddr returns the listen address of the Node.
func (n *Node) ListenAddr() string {
	return n.laddr
}

// Router returns the router configured for the Node.
func (n *Node) Router() *gin.Engine {
	return n.router
}

// Scheduler returns the scheduler configured for the Node.
func (n *Node) Scheduler() *cron.Scheduler {
	return n.scheduler
}

// TLSCertFile returns the TLS certificate path of the Node.
func (n *Node) TLSCertFile() string {
	return filepath.Join(n.HomeDir(), "tls.crt")
}

// TLSKeyFile returns the TLS key path of the Node.
func (n *Node) TLSKeyFile() string {
	return filepath.Join(n.HomeDir(), "tls.key")
}

// Register registers the node on the network if not already registered.
func (n *Node) Register() error {
	node, err := n.Client().Node(gocontext.TODO(), n.NodeAddr())
	if err != nil {
		return fmt.Errorf("failed to query node: %w", err)
	}
	if node != nil {
		return nil
	}

	log.Info("Registering node...")

	// Prepare a message to register the node.
	msg := v3.NewMsgRegisterNodeRequest(
		n.AccAddr(),
		n.GigabytePrices(),
		n.HourlyPrices(),
		n.APIAddrs()[0],
	)

	// Broadcast the registration transaction.
	if err := n.BroadcastTx(gocontext.TODO(), msg); err != nil {
		return fmt.Errorf("failed to broadcast register node tx: %w", err)
	}

	return nil
}

// UpdateDetails updates the node's pricing and address details on the network.
func (n *Node) UpdateDetails() error {
	log.Info("Updating node details...")

	// Prepare a message to update the node's details.
	msg := v3.NewMsgUpdateNodeDetailsRequest(
		n.NodeAddr(),
		n.GigabytePrices(),
		n.HourlyPrices(),
		n.APIAddrs()[0],
	)

	// Broadcast the update transaction.
	if err := n.BroadcastTx(gocontext.TODO(), msg); err != nil {
		return fmt.Errorf("failed to broadcast update node details tx: %w", err)
	}

	return nil
}

// Start initializes the Node's services, scheduler, and HTTPS server.
func (n *Node) Start(errChan chan error) error {
	log.Info("Starting node...")

	go func() {
		// Bring up the service by running pre-defined tasks.
		if err := n.Service().Up(gocontext.TODO()); err != nil {
			errChan <- fmt.Errorf("failed to run service up task: %w", err)
			return
		}
		if err := n.Service().PostUp(); err != nil {
			errChan <- fmt.Errorf("failed to run service post-up task: %w", err)
			return
		}
	}()

	go func() {
		// Start the HTTPS server using the configured TLS certificates and router.
		if err := cmux.ListenAndServeTLS(n.ListenAddr(), n.TLSCertFile(), n.TLSKeyFile(), n.Router()); err != nil {
			errChan <- fmt.Errorf("failed to listen and serve tls: %w", err)
			return
		}
	}()

	// Register the node and update its details.
	if err := n.Register(); err != nil {
		return fmt.Errorf("failed to register node: %w", err)
	}
	if err := n.UpdateDetails(); err != nil {
		return fmt.Errorf("failed to update node details: %w", err)
	}

	// Start the cron scheduler to execute periodic tasks.
	if err := n.Scheduler().Start(); err != nil {
		return fmt.Errorf("failed to start scheduler: %w", err)
	}

	return nil
}

// Stop gracefully stops the Node's operations.
func (n *Node) Stop() error {
	if err := n.Service().PreDown(); err != nil {
		return fmt.Errorf("failed to run service pre-down task: %w", err)
	}
	if err := n.Service().Down(gocontext.TODO()); err != nil {
		return fmt.Errorf("failed to run service down task: %w", err)
	}
	if err := n.Service().PostDown(); err != nil {
		return fmt.Errorf("failed to run service post-up task: %w", err)
	}

	return nil
}
