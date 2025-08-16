package node

import (
	"context"
	"fmt"
	"sync"

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

// Register registers the node on the network if not already registered.
func (n *Node) Register(ctx context.Context) error {
	// Query the network to check if the node is already registered.
	node, err := n.Client().Node(ctx, n.NodeAddr())
	if err != nil {
		return fmt.Errorf("failed to query node: %w", err)
	}
	if node != nil {
		log.Info("Node already registered", "addr", n.NodeAddr())
		return nil
	}

	log.Info("Registering node",
		"gigabyte_price", n.GigabytePrices(),
		"hourly_price", n.HourlyPrices(),
		"remote_addrs", n.APIAddrs(),
	)

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

	log.Info("Node registered successfully", "addr", n.NodeAddr())
	return nil
}

// UpdateDetails updates the node's pricing and address details on the network.
func (n *Node) UpdateDetails(ctx context.Context) error {
	log.Info("Updating node details",
		"gigabyte_prices", n.GigabytePrices(),
		"hourly_prices", n.HourlyPrices(),
		"remote_addrs", n.APIAddrs(),
	)

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

	log.Info("Node details updated successfully", "addr", n.NodeAddr())
	return nil
}

// Start initializes the Node's services, scheduler, and API server.
func (n *Node) Start(ctx context.Context) error {
	// Merge passed-in context with node's primary context
	ctx, _ = utils.AnyDoneContext(n.ctx, ctx)

	// Register the node and update its details.
	if err := n.Register(ctx); err != nil {
		return fmt.Errorf("failed to register node: %w", err)
	}
	if err := n.UpdateDetails(ctx); err != nil {
		return fmt.Errorf("failed to update node details: %w", err)
	}

	var start sync.WaitGroup
	start.Add(3)

	// Launch the service stack as a background goroutine.
	n.eg.Go(func() error {
		log.Info("Starting service routine")
		defer start.Done()

		log.Info("Running service up task")
		if err := n.Service().Up(ctx); err != nil {
			return fmt.Errorf("failed to run service up task: %w", err)
		}

		log.Info("Running service post-up task")
		if err := n.Service().PostUp(ctx); err != nil {
			return fmt.Errorf("failed to run service post-up task: %w", err)
		}

		n.eg.Go(func() error {
			defer log.Info("Exiting service routine")
			if err := n.Service().Wait(); err != nil {
				return fmt.Errorf("failed to wait service: %w", err)
			}

			return nil
		})

		return nil
	})

	// Launch the cron-based job scheduler in the background.
	n.eg.Go(func() error {
		log.Info("Starting scheduler routine")
		defer start.Done()

		log.Info("Starting scheduler")
		if err := n.Scheduler().Start(ctx); err != nil {
			return fmt.Errorf("failed to start scheduler: %w", err)
		}

		n.eg.Go(func() error {
			defer log.Info("Exiting scheduler routine")
			if err := n.Scheduler().Wait(); err != nil {
				return fmt.Errorf("failed to wait scheduler: %w", err)
			}

			return nil
		})

		return nil
	})

	// Launch the API server using the configured TLS certificates and router.
	n.eg.Go(func() error {
		log.Info("Starting API server routine")
		defer start.Done()

		log.Info("Starting API server",
			"listen_on", n.APIListenAddr(),
			"tls_cert_file", n.TLSCertFile(),
			"tls_key_file", n.TLSKeyFile(),
		)
		n.eg.Go(func() error {
			defer log.Info("Exiting API server routine")
			if err := cmux.ListenAndServeTLS(ctx, n.APIListenAddr(), n.TLSCertFile(), n.TLSKeyFile(), n.Router()); err != nil {
				return fmt.Errorf("failed to listen and serve tls: %w", err)
			}

			return nil
		})

		return nil
	})

	// Wait until all routines started
	start.Wait()
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

	log.Info("Running service pre-down task")
	if err := n.Service().PreDown(); err != nil {
		return fmt.Errorf("failed to run service pre-down task: %w", err)
	}

	log.Info("Running service down task")
	if err := n.Service().Down(); err != nil {
		return fmt.Errorf("failed to run service down task: %w", err)
	}

	log.Info("Running service post-down task")
	if err := n.Service().PostDown(); err != nil {
		return fmt.Errorf("failed to run service post-up task: %w", err)
	}

	log.Info("Stopping scheduler")
	if err := n.scheduler.Stop(); err != nil {
		return fmt.Errorf("failed to stop scheduler: %w", err)
	}

	return nil
}
