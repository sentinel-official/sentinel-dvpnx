package workers

import (
	"context"
	"errors"

	cosmossdkmath "cosmossdk.io/math"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"

	nodecontext "github.com/sentinel-official/dvpn-node/context"
	"github.com/sentinel-official/dvpn-node/database/operations"
)

// handleSyncSessionUsageWithDB returns a function that updates session usage statistics in the database.
func handleSyncSessionUsageWithDB(ctx *nodecontext.Context) func() error {
	return func() error {
		// Fetch peer statistics from the service.
		items, err := ctx.Service().PeerStatistics(context.TODO())
		if err != nil {
			return err
		}

		// Iterate through each peer statistic and update the database.
		for _, item := range items {
			// Convert download and upload bytes to strings for database storage.
			downloadBytes := cosmossdkmath.NewInt(item.DownloadBytes).String()
			uploadBytes := cosmossdkmath.NewInt(item.UploadBytes).String()

			// Query to find the session by peer_key.
			query := map[string]interface{}{
				"peer_key": item.Key,
			}

			// Define the updates to apply to the session.
			updates := map[string]interface{}{
				"download_bytes": downloadBytes,
				"upload_bytes":   uploadBytes,
			}

			// Apply the updates to the session in the database.
			if _, err := operations.SessionFindOneAndUpdate(ctx.DB(), query, updates); err != nil {
				return err
			}
		}

		// Return nil if all updates are successful.
		return nil
	}
}

// handleValidateSessionUsage returns a function that checks session usage and removes peers if they exceed limits.
func handleValidateSessionUsage(ctx *nodecontext.Context) func() error {
	return func() error {
		// Retrieve all sessions from the database.
		items, err := operations.SessionFind(ctx.DB(), nil)
		if err != nil {
			return err
		}

		// Iterate through each session to validate usage.
		for _, item := range items {
			removePeer := false

			// Check if the session has exceeded its byte limit.
			if item.GetBytes().GTE(item.GetMaxBytes()) {
				removePeer = true
			}

			// Check if the session has exceeded its duration limit.
			if item.GetDuration() >= item.GetMaxDuration() {
				removePeer = true
			}

			// Skip to the next session if the peer should not be removed.
			if removePeer {
				if err := ctx.RemovePeerIfExistsForKey(context.TODO(), item.PeerKey); err != nil {
					return err
				}
			}
		}

		return nil
	}
}

// handleSyncSessionUsageWithBlockchain returns a function that synchronizes session usage with the blockchain.
func handleSyncSessionUsageWithBlockchain(ctx *nodecontext.Context) func() error {
	return func() error {
		// Retrieve all sessions from the database.
		items, err := operations.SessionFind(ctx.DB(), nil)
		if err != nil {
			return err
		}

		// Prepare a list of transaction messages for session updates.
		var messages []cosmossdk.Msg
		for _, item := range items {
			// Query session details from the blockchain.
			session, err := ctx.QuerySession(context.TODO(), item.GetID())
			if err != nil {
				return err
			}

			if session != nil {
				// Create a message to update the session and add it to the list of messages.
				msg := item.MsgUpdateSessionRequest()
				messages = append(messages, msg)
			}
		}

		// Broadcast the transaction messages to the blockchain.
		res, err := ctx.BroadcastTx(context.TODO(), messages...)
		if err != nil {
			return err
		}
		if !res.TxResult.IsOK() {
			return errors.New(res.TxResult.GetLog())
		}

		return nil
	}
}

// handleValidateSessions returns a function that verifies sessions and removes peers if necessary.
func handleValidateSessions(ctx *nodecontext.Context) func() error {
	return func() error {
		// Retrieve all sessions from the database.
		items, err := operations.SessionFind(ctx.DB(), nil)
		if err != nil {
			return err
		}

		// Iterate through each session to check its status.
		for _, item := range items {
			// Query session details from the blockchain.
			session, err := ctx.QuerySession(context.TODO(), item.GetID())
			if err != nil {
				return err
			}

			if session == nil || !session.GetStatus().Equal(v1base.StatusActive) {
				// Attempt to remove the peer identified by session.PeerKey.
				if err := ctx.RemovePeerIfExistsForKey(context.TODO(), item.PeerKey); err != nil {
					return err
				}
			}

			if session == nil {
				// Delete the session from the database if it does not exist on the blockchain.
				query := map[string]interface{}{
					"id": item.ID,
				}

				if _, err := operations.SessionFindOneAndDelete(ctx.DB(), query); err != nil {
					return err
				}
			}
		}

		return nil
	}
}
