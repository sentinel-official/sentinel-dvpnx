package workers

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	logger "github.com/sentinel-official/sentinel-go-sdk/libs/log"

	nodecontext "github.com/sentinel-official/dvpn-node/context"
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
func NewSessionUsageSyncWithBlockchainWorker(c *nodecontext.Context, interval time.Duration) cron.Worker {
	log := logger.With("name", nameSessionUsageSyncWithBlockchain)

	handlerFunc := func() error {
		log.Info("Running scheduler worker")

		// Retrieve session records from the database.
		items, err := operations.SessionFind(c.Database(), nil)
		if err != nil {
			return fmt.Errorf("failed to retrieve sessions from database: %w", err)
		}
		if len(items) == 0 {
			return nil
		}

		var msgs []types.Msg
		// Iterate over sessions and prepare messages for updates.
		for _, item := range items {
			session, err := c.Client().Session(context.TODO(), item.GetID())
			if err != nil {
				return fmt.Errorf("failed to query session: %w", err)
			}

			if session != nil {
				// Generate an update message for the session.
				msg := item.MsgUpdateSessionRequest()
				msgs = append(msgs, msg)
			}
		}

		// Broadcast the prepared messages as a transaction.
		res, err := c.BroadcastTx(context.TODO(), msgs...)
		if err != nil {
			return fmt.Errorf("failed to broadcast update session details tx: %w", err)
		}
		if !res.TxResult.IsOK() {
			return fmt.Errorf("update session details tx failed with code %d: %s", res.TxResult.Code, res.TxResult.Log)
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
func NewSessionUsageSyncWithDatabaseWorker(c *nodecontext.Context, interval time.Duration) cron.Worker {
	log := logger.With("name", nameSessionUsageSyncWithDatabase)

	handlerFunc := func() error {
		log.Info("Running scheduler worker")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
			downloadBytes := math.NewInt(item.DownloadBytes).String()
			uploadBytes := math.NewInt(item.UploadBytes).String()

			query := map[string]interface{}{
				"peer_key": item.Key,
			}
			updates := map[string]interface{}{
				"download_bytes": downloadBytes,
				"upload_bytes":   uploadBytes,
			}

			if _, err := operations.SessionFindOneAndUpdate(c.Database(), query, updates); err != nil {
				return fmt.Errorf("failed to update session: %w", err)
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
func NewSessionUsageValidateWorker(c *nodecontext.Context, interval time.Duration) cron.Worker {
	log := logger.With("name", nameSessionUsageValidate)

	handlerFunc := func() error {
		log.Info("Running scheduler worker")

		// Retrieve session records from the database.
		items, err := operations.SessionFind(c.Database(), nil)
		if err != nil {
			return fmt.Errorf("failed to retrieve sessions from database: %w", err)
		}
		if len(items) == 0 {
			return nil
		}

		// Validate session limits and remove peers if needed.
		for _, item := range items {
			removePeer := false

			if item.GetBytes().GTE(item.GetMaxBytes()) {
				removePeer = true
			}
			if item.GetDuration() >= item.GetMaxDuration() {
				removePeer = true
			}

			if removePeer {
				if err := c.RemovePeerIfExistsForKey(context.TODO(), item.PeerKey); err != nil {
					return fmt.Errorf("failed to remove peer: %w", err)
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
func NewSessionValidateWorker(c *nodecontext.Context, interval time.Duration) cron.Worker {
	log := logger.With("name", nameSessionValidate)

	handlerFunc := func() error {
		log.Info("Running scheduler worker")

		// Retrieve session records from the database.
		items, err := operations.SessionFind(c.Database(), nil)
		if err != nil {
			return fmt.Errorf("failed to retrieve sessions from database: %w", err)
		}
		if len(items) == 0 {
			return nil
		}

		// Validate session status and consistency.
		for _, item := range items {
			session, err := c.Client().Session(context.TODO(), item.GetID())
			if err != nil {
				return fmt.Errorf("failed to query session: %w", err)
			}

			// Remove peers if sessions are inactive or missing on the blockchain.
			if session == nil || !session.GetStatus().Equal(v1.StatusActive) {
				if err := c.RemovePeerIfExistsForKey(context.TODO(), item.PeerKey); err != nil {
					return fmt.Errorf("failed to remove peer: %w", err)
				}
			}

			// Delete sessions missing on the blockchain from the database.
			if session == nil {
				query := map[string]interface{}{
					"id": item.ID,
				}
				if _, err := operations.SessionFindOneAndDelete(c.Database(), query); err != nil {
					return fmt.Errorf("failed to delete session: %w", err)
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
