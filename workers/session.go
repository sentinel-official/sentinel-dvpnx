package workers

import (
	gocontext "context"

	sdkmath "cosmossdk.io/math"

	"github.com/sentinel-official/dvpn-node/context"
	"github.com/sentinel-official/dvpn-node/database/operations"
)

// handleSyncSessionUsage returns a function that synchronizes peer session usage statistics with the database.
func handleSyncSessionUsage(ctx *context.Context) func() error {
	return func() error {
		// Retrieve peer statistics from the service.
		statistics, err := ctx.Service().PeerStatistics(gocontext.TODO())
		if err != nil {
			return err
		}

		// Iterate over each peer statistic.
		for _, s := range statistics {
			// Convert download and upload bytes to strings for database storage.
			downloadBytes := sdkmath.NewInt(s.DownloadBytes).String()
			uploadBytes := sdkmath.NewInt(s.UploadBytes).String()

			// Define the query to find the relevant session by peer_key.
			query := map[string]interface{}{
				"peer_key": s.Key,
			}

			// Define the updates to be applied to the found session.
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

// handleValidateSessionUsage returns a function that validates session usage and removes peers if necessary.
func handleValidateSessionUsage(ctx *context.Context) func() error {
	return func() error {
		// Retrieve sessions from the database
		sessions, err := operations.SessionFind(ctx.DB(), nil)
		if err != nil {
			return err
		}

		// Iterate over each session
		for _, s := range sessions {
			// Determine if the peer should be removed
			removePeer := false

			// Check if the session has exceeded its byte limit
			if s.GetBytes().GTE(s.GetMaxBytes()) {
				removePeer = true
			}

			// Check if the session has exceeded its duration limit
			if s.GetDuration() >= s.GetMaxDuration() {
				removePeer = true
			}

			// If the peer should not be removed, skip to the next session
			if !removePeer {
				continue
			}

			// Attempt to remove the peer identified by session.PeerKey
			if err := ctx.RemovePeerIfExistsForKey(gocontext.TODO(), s.PeerKey); err != nil {
				return err
			}
		}

		return nil
	}
}
