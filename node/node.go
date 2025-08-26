package node

import (
	"context"
	"fmt"

	"github.com/sentinel-official/sentinel-go-sdk/libs/cmux"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	"github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types/v3"
	"golang.org/x/sync/errgroup"

	"github.com/sentinel-official/sentinel-dvpnx/core"
)

// Node represents the application node, holding its context, scheduler, and server.
type Node struct {
	*core.Context

	scheduler *cron.Scheduler // Scheduler for managing periodic tasks.
	server    *cmux.Server    // HTTP server for handling API requests.

	eg *errgroup.Group
}

// New creates a new Node with the provided context.
func New(c *core.Context) *Node {
	return &Node{
		Context: c,
		eg:      &errgroup.Group{},
	}
}

// WithScheduler sets the scheduler for the Node and returns the updated Node.
func (n *Node) WithScheduler(v *cron.Scheduler) *Node {
	n.scheduler = v
	return n
}

// WithServer sets the server for the Node and returns the updated Node.
func (n *Node) WithServer(v *cmux.Server) *Node {
	n.server = v
	return n
}

// Scheduler returns the scheduler configured for the Node.
func (n *Node) Scheduler() *cron.Scheduler {
	return n.scheduler
}

// Server returns the server configured for the Node.
func (n *Node) Server() *cmux.Server {
	return n.server
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
		return fmt.Errorf("broadcasting tx with register_node msg: %w", err)
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
		return fmt.Errorf("broadcasting tx with update_node_details msg: %w", err)
	}

	log.Info("Node details updated successfully", "addr", n.NodeAddr())
	return nil
}

// Start initializes the Node's services, scheduler, and API server.
func (n *Node) Start() error {
	sg := &errgroup.Group{}

	// Launch the service stack as a background goroutine.
	sg.Go(func() error {
		log.Info("Running service pre-up task")
		if err := n.Service().PreUp(); err != nil {
			return fmt.Errorf("running service pre-up task: %w", err)
		}

		log.Info("Running service up task")
		if err := n.Service().Up(); err != nil {
			return fmt.Errorf("running service up task: %w", err)
		}

		log.Info("Running service post-up task")
		if err := n.Service().PostUp(); err != nil {
			return fmt.Errorf("running service post-up task: %w", err)
		}

		n.eg.Go(func() error {
			if err := n.Service().Wait(); err != nil {
				return fmt.Errorf("waiting service: %w", err)
			}

			return nil
		})

		return nil
	})

	// Launch the cron-based job scheduler in the background.
	sg.Go(func() error {
		log.Info("Starting scheduler")
		if err := n.Scheduler().Start(); err != nil {
			return fmt.Errorf("starting scheduler: %w", err)
		}

		n.eg.Go(func() error {
			if err := n.Scheduler().Wait(); err != nil {
				return fmt.Errorf("waiting scheduler: %w", err)
			}

			return nil
		})

		return nil
	})

	// Launch the API server using the configured TLS certificates and router.
	sg.Go(func() error {
		log.Info("Starting API server", "addr", n.APIListenAddr())
		if err := n.Server().Start(); err != nil {
			return fmt.Errorf("starting API server: %w", err)
		}

		n.eg.Go(func() error {
			if err := n.Server().Wait(); err != nil {
				return fmt.Errorf("waiting API server: %w", err)
			}

			return nil
		})

		return nil
	})

	// Wait until all routines started
	if err := sg.Wait(); err != nil {
		return err
	}

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
	sg := &errgroup.Group{}

	sg.Go(func() error {
		log.Info("Running service pre-down task")
		if err := n.Service().PreDown(); err != nil {
			return fmt.Errorf("running service pre-down task: %w", err)
		}

		log.Info("Running service down task")
		if err := n.Service().Down(); err != nil {
			return fmt.Errorf("running service down task: %w", err)
		}

		log.Info("Running service post-down task")
		if err := n.Service().PostDown(); err != nil {
			return fmt.Errorf("running service post-up task: %w", err)
		}

		return nil
	})

	sg.Go(func() error {
		log.Info("Stopping scheduler")
		if err := n.Scheduler().Stop(); err != nil {
			return fmt.Errorf("stopping scheduler: %w", err)
		}

		return nil
	})

	sg.Go(func() error {
		log.Info("Stopping API server")
		if err := n.Server().Stop(); err != nil {
			return fmt.Errorf("stopping API server: %w", err)
		}

		return nil
	})

	if err := sg.Wait(); err != nil {
		return err
	}

	return nil
}
