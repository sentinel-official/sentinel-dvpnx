package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	"github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types/v3"

	"github.com/sentinel-official/sentinel-dvpnx/core"
)

const NameNodeStatusUpdate = "node_status_update"

// NewNodeStatusUpdateWorker creates a worker to periodically update the node's status to active on the blockchain.
// This worker broadcasts a transaction to mark the node as active at regular intervals.
func NewNodeStatusUpdateWorker(c *core.Context, interval time.Duration) cron.Worker {
	// Handler function that updates the node's status to active.
	handlerFunc := func(ctx context.Context) error {
		// Create a message to update the node's status to active.
		msg := v3.NewMsgUpdateNodeStatusRequest(
			c.AccAddr().Bytes(),
			v1.StatusActive,
		)

		// Broadcast the transaction message to the blockchain.
		if err := c.BroadcastTx(ctx, msg); err != nil {
			return fmt.Errorf("broadcasting tx with update_node_status msg: %w", err)
		}

		return nil
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker(NameNodeStatusUpdate).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithRetryDelay(5 * time.Second)
}
