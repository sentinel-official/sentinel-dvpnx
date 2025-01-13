package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	logger "github.com/sentinel-official/sentinel-go-sdk/libs/log"

	nodecontext "github.com/sentinel-official/dvpn-node/context"
)

const nameNodeStatusUpdate = "node_status_update"

// NewNodeStatusUpdateWorker creates a worker to periodically update the node's status to active on the blockchain.
// This worker broadcasts a transaction to mark the node as active at regular intervals.
func NewNodeStatusUpdateWorker(c *nodecontext.Context, interval time.Duration) cron.Worker {
	log := logger.With("name", nameNodeStatusUpdate)

	// Handler function that updates the node's status to active.
	handlerFunc := func() error {
		log.Info("Running scheduler worker")

		// Create a message to update the node's status to active.
		msg := v3.NewMsgUpdateNodeStatusRequest(
			c.AccAddr().Bytes(),
			v1.StatusActive,
		)

		// Broadcast the transaction message to the blockchain.
		res, err := c.BroadcastTx(context.TODO(), msg)
		if err != nil {
			return fmt.Errorf("failed to broadcast update node status tx: %w", err)
		}
		if !res.TxResult.IsOK() {
			return fmt.Errorf("update node status tx failed with code %d: %s", res.TxResult.Code, res.TxResult.Log)
		}

		return nil
	}

	// Error handling function to log failures.
	onErrorFunc := func(err error) bool {
		log.Error("Failed to run scheduler worker", "msg", err)
		return false
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker().
		WithName(nameNodeStatusUpdate).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}
