package core

import (
	"context"
	"fmt"

	"github.com/sentinel-official/sentinel-go-sdk/libs/log"
)

// RemovePeerIfExists checks if a peer exists, and removes it if found.
func (c *Context) RemovePeerIfExists(ctx context.Context, req []byte) error {
	// Check if the peer exists.
	ok, err := c.Service().HasPeer(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to check peer existence: %w", err)
	}
	if !ok {
		return nil
	}

	// Remove the peer if it exists.
	id, err := c.Service().RemovePeer(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to remove peer: %w", err)
	}

	log.Info("Peer has been removed from service", "id", id)
	return nil
}
