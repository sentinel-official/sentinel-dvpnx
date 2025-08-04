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

const nameNodeStatusUpdate = "node_status_update"

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
			return fmt.Errorf("failed to broadcast update node status tx: %w", err)
		}

		return nil
	}

	// Error handling function to log failures.
	onErrorFunc := func(err error) bool {
		return false
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker().
		WithName(nameNodeStatusUpdate).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}
