package workers

import (
	gocontext "context"
	"fmt"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	logger "github.com/sentinel-official/sentinel-go-sdk/libs/log"

	"github.com/sentinel-official/dvpn-node/context"
	"github.com/sentinel-official/dvpn-node/database/operations"
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
func NewSessionUsageSyncWithBlockchainWorker(c *context.Context, interval time.Duration) cron.Worker {
	log := logger.With("name", nameSessionUsageSyncWithBlockchain)

	handlerFunc := func() error {
		// Retrieve session records from the database.
		items, err := operations.SessionFind(c.Database(), nil)
		if err != nil {
			return fmt.Errorf("failed to retrieve sessions from the database: %w", err)
		}
		if len(items) == 0 {
			return nil
		}

		var msgs []types.Msg
		// Iterate over sessions and prepare messages for updates.
		for _, item := range items {
			session, err := c.Client().Session(gocontext.TODO(), item.GetID())
			if err != nil {
				return fmt.Errorf("failed to query session from the blockchain: %w", err)
			}

			if session != nil {
				// Generate an update message for the session.
				msg := item.MsgUpdateSessionRequest()
				msgs = append(msgs, msg)
			}
		}

		// If no update messages were generated, exit early.
		if len(msgs) == 0 {
			return nil
		}

		// Broadcast the prepared messages as a transaction.
		if err := c.BroadcastTx(gocontext.TODO(), msgs...); err != nil {
			return fmt.Errorf("failed to broadcast update session tx: %w", err)
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
		WithName(nameSessionUsageSyncWithBlockchain).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}

// NewSessionUsageSyncWithDatabaseWorker creates a worker that updates session usage in the database.
// This worker fetches usage data from the peer service and updates the corresponding database records.
func NewSessionUsageSyncWithDatabaseWorker(c *context.Context, interval time.Duration) cron.Worker {
	log := logger.With("name", nameSessionUsageSyncWithDatabase)

	handlerFunc := func() error {
		// Create a context with a timeout to fetch peer statistics.
		ctx, cancel := gocontext.WithTimeout(gocontext.TODO(), 5*time.Second)
		defer cancel()

		// Fetch peer usage statistics from the service.
		items, err := c.Service().PeerStatistics(ctx)
		if err != nil {
			return fmt.Errorf("failed to fetch peer statistics: %w", err)
		}
		if len(items) == 0 {
			return nil
		}

		// Update the database with the fetched statistics.
		for _, item := range items {
			// Convert usage statistics to strings for database storage.
			downloadBytes := math.NewInt(item.DownloadBytes).String()
			uploadBytes := math.NewInt(item.UploadBytes).String()

			// Define query to find the session by peer key.
			query := map[string]interface{}{
				"peer_key": item.Key,
			}

			// Define updates to apply to the session record.
			updates := map[string]interface{}{
				"download_bytes": downloadBytes,
				"upload_bytes":   uploadBytes,
			}

			// Update the session in the database.
			if _, err := operations.SessionFindOneAndUpdate(c.Database(), query, updates); err != nil {
				return fmt.Errorf("failed to update session with peer key %s: %w", item.Key, err)
			}
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
		WithName(nameSessionUsageSyncWithDatabase).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}

// NewSessionUsageValidateWorker creates a worker that validates session usage limits and removes peers if necessary.
// This worker checks if sessions exceed their maximum byte or duration limits and removes peers accordingly.
func NewSessionUsageValidateWorker(c *context.Context, interval time.Duration) cron.Worker {
	log := logger.With("name", nameSessionUsageValidate)

	handlerFunc := func() error {
		// Retrieve session records from the database.
		items, err := operations.SessionFind(c.Database(), nil)
		if err != nil {
			return fmt.Errorf("failed to retrieve sessions from the database: %w", err)
		}
		if len(items) == 0 {
			return nil
		}

		// Validate session limits and remove peers if needed.
		for _, item := range items {
			removePeer := false

			// Check if the session exceeds the maximum allowed bytes.
			maxBytes := item.GetMaxBytes()
			if !maxBytes.IsZero() && item.GetBytes().GTE(maxBytes) {
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
				if err := c.RemovePeerIfExistsForKey(gocontext.TODO(), item.PeerKey); err != nil {
					return fmt.Errorf("failed to remove peer with key %s: %w", item.PeerKey, err)
				}
			}
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
		WithName(nameSessionUsageValidate).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}

// NewSessionValidateWorker creates a worker that validates session status and removes peers if necessary.
// This worker ensures sessions are active and consistent between the database and blockchain.
func NewSessionValidateWorker(c *context.Context, interval time.Duration) cron.Worker {
	log := logger.With("name", nameSessionValidate)

	handlerFunc := func() error {
		// Retrieve session records from the database.
		items, err := operations.SessionFind(c.Database(), nil)
		if err != nil {
			return fmt.Errorf("failed to retrieve sessions from the database: %w", err)
		}
		if len(items) == 0 {
			return nil
		}

		// Validate session status and consistency.
		for _, item := range items {
			session, err := c.Client().Session(gocontext.TODO(), item.GetID())
			if err != nil {
				return fmt.Errorf("failed to query session from the blockchain: %w", err)
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
				if err := c.RemovePeerIfExistsForKey(gocontext.TODO(), item.PeerKey); err != nil {
					return fmt.Errorf("failed to remove peer with key %s: %w", item.PeerKey, err)
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
					return fmt.Errorf("failed to delete session %d: %w", item.ID, err)
				}
			}
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
		WithName(nameSessionValidate).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}
