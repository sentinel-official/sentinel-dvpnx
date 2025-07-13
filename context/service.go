package context

import (
	"context"
	"fmt"
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
	if err := c.Service().RemovePeer(ctx, req); err != nil {
		return fmt.Errorf("failed to remove peer: %w", err)
	}

	return nil
}
