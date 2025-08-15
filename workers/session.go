package workers

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	"github.com/sentinel-official/sentinelhub/v12/types/v1"

	"github.com/sentinel-official/sentinel-dvpnx/core"
	"github.com/sentinel-official/sentinel-dvpnx/database/operations"
)

const (
	nameSessionUsageSyncWithBlockchain = "session_usage_sync_with_blockchain"
	nameSessionUsageSyncWithDatabase   = "session_usage_sync_with_database"
	nameSessionUsageValidate           = "session_usage_validate"
	nameSessionValidate                = "session_validate"
)

// NewSessionUsageSyncWithBlockchainWorker creates a worker that synchronizes session usage with the blockchain.
// This worker retrieves session data from the database, validates it against the blockchain,
// and broadcasts any updates as transactions.
func NewSessionUsageSyncWithBlockchainWorker(c *core.Context, interval time.Duration) cron.Worker {
	handlerFunc := func(ctx context.Context) error {
		// Retrieve session records from the database.
		items, err := operations.SessionFind(c.Database(), nil)
		if err != nil {
			return fmt.Errorf("retrieving sessions from database: %w", err)
		}

		var msgs []types.Msg
		// Iterate over sessions and prepare messages for updates.
		for _, item := range items {
			session, err := c.Client().Session(ctx, item.GetID())
			if err != nil {
				return fmt.Errorf("querying session %d on blockchain: %w", item.GetID(), err)
			}
			if session == nil {
				continue
			}

			// Generate an update message for the session.
			msg := item.MsgUpdateSessionRequest()
			msgs = append(msgs, msg)
		}

		// Broadcast the prepared messages as a transaction.
		if err := c.BroadcastTx(ctx, msgs...); err != nil {
			return fmt.Errorf("broadcasting tx: %w", err)
		}

		return nil
	}

	// Error handling function to log failures.
	onErrorFunc := func(err error) bool {
		return false
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker().
		WithName(nameSessionUsageSyncWithBlockchain).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}

// NewSessionUsageSyncWithDatabaseWorker creates a worker that updates session usage in the database.
// This worker fetches usage data from the peer service and updates the corresponding database records.
func NewSessionUsageSyncWithDatabaseWorker(c *core.Context, interval time.Duration) cron.Worker {
	handlerFunc := func(ctx context.Context) error {
		// Fetch peer usage statistics from the service.
		items, err := c.Service().PeerStatistics()
		if err != nil {
			return fmt.Errorf("retrieving peer statistics from service: %w", err)
		}

		// Update the database with the fetched statistics.
		for id, item := range items {
			// Convert usage statistics to strings for database storage.
			rxBytes := math.NewInt(item.RxBytes).String()
			txBytes := math.NewInt(item.TxBytes).String()

			// Define query to find the session by peer id.
			query := map[string]interface{}{
				"peer_id": id,
			}

			// Define updates to apply to the session record.
			updates := map[string]interface{}{
				"rx_bytes": rxBytes,
				"tx_bytes": txBytes,
			}

			// Update the session in the database.
			if _, err := operations.SessionFindOneAndUpdate(c.Database(), query, updates); err != nil {
				return fmt.Errorf("updating session for peer %q in database: %w", id, err)
			}
		}

		return nil
	}

	// Error handling function to log failures.
	onErrorFunc := func(err error) bool {
		return false
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker().
		WithName(nameSessionUsageSyncWithDatabase).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}

// NewSessionUsageValidateWorker creates a worker that validates session usage limits and removes peers if necessary.
// This worker checks if sessions exceed their maximum byte or duration limits and removes peers accordingly.
func NewSessionUsageValidateWorker(c *core.Context, interval time.Duration) cron.Worker {
	handlerFunc := func(ctx context.Context) error {
		// Retrieve session records from the database.
		items, err := operations.SessionFind(c.Database(), nil)
		if err != nil {
			return fmt.Errorf("retrieving sessions from database: %w", err)
		}

		// Validate session limits and remove peers if needed.
		for _, item := range items {
			removePeer := false

			// Check if the session exceeds the maximum allowed bytes.
			maxBytes := item.GetMaxBytes()
			if !maxBytes.IsZero() && item.GetTotalBytes().GTE(maxBytes) {
				removePeer = true
			}

			// Check if the session exceeds the maximum allowed duration.
			maxDuration := item.GetMaxDuration()
			if maxDuration != 0 && item.GetDuration() >= maxDuration {
				removePeer = true
			}

			// Ensure that only sessions of the current service type are validated.
			if item.GetServiceType() != c.Service().Type() {
				removePeer = false
			}

			// If the session exceeded any limits, remove the associated peer.
			if removePeer {
				req := item.GetPeerRequest()
				if err := c.RemovePeerIfExists(ctx, req); err != nil {
					return fmt.Errorf("removing peer for session %d from service: %w", item.GetID(), err)
				}
			}
		}

		return nil
	}

	// Error handling function to log failures.
	onErrorFunc := func(err error) bool {
		return false
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker().
		WithName(nameSessionUsageValidate).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}

// NewSessionValidateWorker creates a worker that validates session status and removes peers if necessary.
// This worker ensures sessions are active and consistent between the database and blockchain.
func NewSessionValidateWorker(c *core.Context, interval time.Duration) cron.Worker {
	handlerFunc := func(ctx context.Context) error {
		// Retrieve session records from the database.
		items, err := operations.SessionFind(c.Database(), nil)
		if err != nil {
			return fmt.Errorf("retrieving sessions from database: %w", err)
		}

		// Validate session status and consistency.
		for _, item := range items {
			session, err := c.Client().Session(ctx, item.GetID())
			if err != nil {
				return fmt.Errorf("querying session %d on blockchain: %w", item.GetID(), err)
			}

			removePeer := false

			// Remove peer if the session is missing on the blockchain.
			if session == nil {
				removePeer = true
			}
			// Remove peer if the session status is not active.
			if session != nil && !session.GetStatus().Equal(v1.StatusActive) {
				removePeer = true
			}
			// Validate only sessions of the current service type.
			if item.GetServiceType() != c.Service().Type() {
				removePeer = false
			}

			// Remove the associated peer if validation fails.
			if removePeer {
				req := item.GetPeerRequest()
				if err := c.RemovePeerIfExists(ctx, req); err != nil {
					return fmt.Errorf("removing peer for session %d from service: %w", item.GetID(), err)
				}
			}

			deleteSession := false

			// Delete session if the session is missing on the blockchain.
			if session == nil {
				deleteSession = true
			}

			// Delete the session record from the database if not found on the blockchain.
			if deleteSession {
				query := map[string]interface{}{
					"id": item.ID,
				}

				if _, err := operations.SessionFindOneAndDelete(c.Database(), query); err != nil {
					return fmt.Errorf("deleting session %d from database: %w", item.ID, err)
				}
			}
		}

		return nil
	}

	// Error handling function to log failures.
	onErrorFunc := func(err error) bool {
		return false
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker().
		WithName(nameSessionValidate).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}
