package node

import (
	gocontext "context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	"github.com/sentinel-official/sentinel-go-sdk/libs/log"

	"github.com/sentinel-official/dvpn-node/context"
	"github.com/sentinel-official/dvpn-node/utils"
)

// Node represents the application node, holding its context, router, and scheduler.
type Node struct {
	listenAddr  string          // Address the Node listens on for incoming requests.
	router      *gin.Engine     // HTTP router for handling API requests.
	scheduler   *cron.Scheduler // Scheduler for managing periodic tasks.
	tlsCertPath string          // Path to the TLS certificate file.
	tlsKeyPath  string          // Path to the TLS key file.
}

// New creates a new Node with the provided context.
func New() *Node {
	return &Node{}
}

// WithListenAddr sets the listen address for the Node and returns the updated Node.
func (n *Node) WithListenAddr(v string) *Node {
	n.listenAddr = v
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

// WithTLSCertPath sets the TLS certificate path for the Node and returns the updated Node.
func (n *Node) WithTLSCertPath(v string) *Node {
	n.tlsCertPath = v
	return n
}

// WithTLSKeyPath sets the TLS key path for the Node and returns the updated Node.
func (n *Node) WithTLSKeyPath(v string) *Node {
	n.tlsKeyPath = v
	return n
}

// ListenAddr returns the listen address of the Node.
func (n *Node) ListenAddr() string {
	return n.listenAddr
}

// Router returns the router configured for the Node.
func (n *Node) Router() *gin.Engine {
	return n.router
}

// Scheduler returns the scheduler configured for the Node.
func (n *Node) Scheduler() *cron.Scheduler {
	return n.scheduler
}

// TLSCertPath returns the TLS certificate path of the Node.
func (n *Node) TLSCertPath() string {
	return n.tlsCertPath
}

// TLSKeyPath returns the TLS key path of the Node.
func (n *Node) TLSKeyPath() string {
	return n.tlsKeyPath
}

// Register registers the node on the network if not already registered.
func (n *Node) Register(c *context.Context) error {
	node, err := c.Client().Node(gocontext.TODO(), c.NodeAddr())
	if err != nil {
		return fmt.Errorf("failed to query node: %w", err)
	}
	if node != nil {
		return nil
	}

	log.Info("Registering node...")

	// Prepare a message to register the node.
	msg := v3.NewMsgRegisterNodeRequest(
		c.AccAddr(),
		c.GigabytePrices(),
		c.HourlyPrices(),
		c.APIAddrs()[0],
	)

	// Broadcast the registration transaction.
	res, err := c.BroadcastTx(gocontext.TODO(), msg)
	if err != nil {
		return fmt.Errorf("failed to broadcast register node tx: %w", err)
	}
	if !res.TxResult.IsOK() {
		err := errors.New(res.TxResult.Log)
		return fmt.Errorf("register node tx failed with code %d: %w", res.TxResult.Code, err)
	}

	return nil
}

// UpdateDetails updates the node's pricing and address details on the network.
func (n *Node) UpdateDetails(c *context.Context) error {
	log.Info("Updating node details...")

	// Prepare a message to update the node's details.
	msg := v3.NewMsgUpdateNodeDetailsRequest(
		c.NodeAddr(),
		c.GigabytePrices(),
		c.HourlyPrices(),
		c.APIAddrs()[0],
	)

	// Broadcast the update transaction.
	res, err := c.BroadcastTx(gocontext.TODO(), msg)
	if err != nil {
		return fmt.Errorf("failed to broadcast update node details tx: %w", err)
	}
	if !res.TxResult.IsOK() {
		err := errors.New(res.TxResult.Log)
		return fmt.Errorf("update node deatils tx failed with code %d: %w", res.TxResult.Code, err)
	}

	return nil
}

// Start initializes the Node's services, scheduler, and HTTPS server.
func (n *Node) Start(c *context.Context, errChan chan error) error {
	log.Info("Starting node...")

	go func() {
		// Bring up the service by running pre-defined tasks.
		if err := c.Service().Up(gocontext.TODO()); err != nil {
			errChan <- fmt.Errorf("failed to run service up task: %w", err)
			return
		}
		if err := c.Service().PostUp(); err != nil {
			errChan <- fmt.Errorf("failed to run service post-up task: %w", err)
			return
		}
	}()

	go func() {
		// Start the HTTPS server using the configured TLS certificates and router.
		if err := utils.ListenAndServeTLS(n.ListenAddr(), n.TLSCertPath(), n.TLSKeyPath(), n.Router()); err != nil {
			errChan <- fmt.Errorf("failed to listen and serve tls: %w", err)
			return
		}
	}()

	// Register the node and update its details.
	if err := n.Register(c); err != nil {
		return fmt.Errorf("failed to register node: %w", err)
	}
	if err := n.UpdateDetails(c); err != nil {
		return fmt.Errorf("failed to update node details: %w", err)
	}

	// Start the cron scheduler to execute periodic tasks.
	if err := n.Scheduler().Start(); err != nil {
		return fmt.Errorf("failed to start scheduler: %w", err)
	}

	return nil
}

// Stop gracefully stops the Node's operations.
func (n *Node) Stop(c *context.Context) error {
	if err := c.Service().PreDown(); err != nil {
		return fmt.Errorf("failed to run service pre-down task: %w", err)
	}
	if err := c.Service().Down(gocontext.TODO()); err != nil {
		return fmt.Errorf("failed to run service down task: %w", err)
	}
	if err := c.Service().PostDown(); err != nil {
		return fmt.Errorf("failed to run service post-up task: %w", err)
	}

	return nil
}
