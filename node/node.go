package node

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cmux"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	"github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/sentinel-official/sentinel-go-sdk/utils"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types/v3"
	"golang.org/x/sync/errgroup"

	"github.com/sentinel-official/sentinel-dvpnx/core"
)

// Node represents the application node, holding its context, router, and scheduler.
type Node struct {
	*core.Context
	router    *gin.Engine     // HTTP router for handling API requests.
	scheduler *cron.Scheduler // Scheduler for managing periodic tasks.

	ctx    context.Context
	cancel context.CancelFunc
	eg     *errgroup.Group
}

// New creates a new Node with the provided context.
func New(c *core.Context) *Node {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)

	return &Node{
		Context: c,
		cancel:  cancel,
		eg:      eg,
		ctx:     ctx,
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
func (n *Node) Register(ctx context.Context) error {
	// Query the network to check if the node is already registered.
	node, err := n.Client().Node(ctx, n.NodeAddr())
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
		n.APIAddrs(),
	)

	// Broadcast the registration transaction.
	if err := n.BroadcastTx(ctx, msg); err != nil {
		return fmt.Errorf("failed to broadcast register node tx: %w", err)
	}

	return nil
}

// UpdateDetails updates the node's pricing and address details on the network.
func (n *Node) UpdateDetails(ctx context.Context) error {
	log.Info("Updating node details...")

	// Prepare a message to update the node's details.
	msg := v3.NewMsgUpdateNodeDetailsRequest(
		n.NodeAddr(),
		n.GigabytePrices(),
		n.HourlyPrices(),
		n.APIAddrs(),
	)

	// Broadcast the update transaction.
	if err := n.BroadcastTx(ctx, msg); err != nil {
		return fmt.Errorf("failed to broadcast update node details tx: %w", err)
	}

	return nil
}

// Start initializes the Node's services, scheduler, and API server.
func (n *Node) Start(ctx context.Context) error {
	log.Info("Starting node...")

	// Merge passed-in context with node's primary context
	ctx, _ = utils.AnyDoneContext(n.ctx, ctx)

	// Register the node and update its details.
	if err := n.Register(ctx); err != nil {
		return fmt.Errorf("failed to register node: %w", err)
	}
	if err := n.UpdateDetails(ctx); err != nil {
		return fmt.Errorf("failed to update node details: %w", err)
	}

	// Launch the service stack as a background goroutine.
	n.eg.Go(func() error {
		if err := n.Service().Up(ctx); err != nil {
			return fmt.Errorf("failed to run service up task: %w", err)
		}
		if err := n.Service().PostUp(ctx); err != nil {
			return fmt.Errorf("failed to run service post-up task: %w", err)
		}
		if err := n.Service().Wait(); err != nil {
			return fmt.Errorf("failed to wait for service: %w", err)
		}

		return nil
	})

	// Launch the API server using the configured TLS certificates and router.
	n.eg.Go(func() error {
		if err := cmux.ListenAndServeTLS(
			ctx, n.APIListenAddr(), n.TLSCertFile(), n.TLSKeyFile(), n.Router(),
		); err != nil {
			return fmt.Errorf("failed to listen and serve tls: %w", err)
		}

		return nil
	})

	// Launch the cron-based job scheduler in the background.
	n.eg.Go(func() error {
		if err := n.Scheduler().Start(ctx); err != nil {
			return fmt.Errorf("failed to start scheduler: %w", err)
		}
		if err := n.Scheduler().Wait(); err != nil {
			return fmt.Errorf("failed to wait for scheduler: %w", err)
		}

		return nil
	})

	log.Info("Node started successfully")
	return nil
}

// Wait blocks until all background goroutines launched exit.
func (n *Node) Wait() error {
	if err := n.eg.Wait(); err != nil {
		return err
	}

	return nil
}

// Stop gracefully stops the Node's operations.
func (n *Node) Stop() error {
	// Cancel the node's context to signal all routines to exit.
	n.cancel()

	// Attempt to gracefully stop the service.
	if err := n.Service().PreDown(); err != nil {
		return fmt.Errorf("failed to run service pre-down task: %w", err)
	}
	if err := n.Service().Down(); err != nil {
		return fmt.Errorf("failed to run service down task: %w", err)
	}
	if err := n.Service().PostDown(); err != nil {
		return fmt.Errorf("failed to run service post-up task: %w", err)
	}

	// Attempt to gracefully stop the scheduler.
	if err := n.scheduler.Stop(); err != nil {
		return fmt.Errorf("failed to stop scheduler: %w", err)
	}

	log.Info("Node stopped successfully")
	return nil
}
