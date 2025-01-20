package context

import (
	"context"
	"fmt"

	"github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/v2ray"
	"github.com/sentinel-official/sentinel-go-sdk/wireguard"
)

// HasPeerForKey checks if a peer with the given key exists in the current service.
func (c *Context) HasPeerForKey(ctx context.Context, s string) (bool, error) {
	// Determine the service type.
	t := c.Service().Type()

	switch t {
	case types.ServiceTypeV2Ray:
		// Create a peer existence request for V2Ray service.
		req, err := v2ray.NewHasPeerRequestFromKey(s)
		if err != nil {
			return false, fmt.Errorf("failed to create v2ray has peer request: %w", err)
		}

		// Check if the peer exists for V2Ray service.
		ok, err := c.Service().HasPeer(ctx, req)
		if err != nil {
			return false, fmt.Errorf("failed to check v2ray peer existence: %w", err)
		}

		return ok, nil
	case types.ServiceTypeWireGuard:
		// Create a peer existence request for WireGuard service.
		req, err := wireguard.NewHasPeerRequestFromKey(s)
		if err != nil {
			return false, fmt.Errorf("failed to create wireguard has peer request: %w", err)
		}

		// Check if the peer exists for WireGuard service.
		ok, err := c.Service().HasPeer(ctx, req)
		if err != nil {
			return false, fmt.Errorf("failed to check wireguard peer existence: %w", err)
		}

		return ok, nil
	default:
		// Return an error for unsupported service types.
		return false, fmt.Errorf("invalid service type %s", t)
	}
}

// RemovePeerForKey removes the peer with the given key from the current service.
func (c *Context) RemovePeerForKey(ctx context.Context, s string) error {
	// Determine the service type.
	t := c.Service().Type()

	switch t {
	case types.ServiceTypeV2Ray:
		// Create a peer removal request for V2Ray service.
		req, err := v2ray.NewRemovePeerRequestFromKey(s)
		if err != nil {
			return fmt.Errorf("failed to create v2ray peer removal request: %w", err)
		}

		// Remove the peer for V2Ray service.
		if err := c.Service().RemovePeer(ctx, req); err != nil {
			return fmt.Errorf("failed to remove v2ray peer: %w", err)
		}

		return nil
	case types.ServiceTypeWireGuard:
		// Create a peer removal request for WireGuard service.
		req, err := wireguard.NewRemovePeerRequestFromKey(s)
		if err != nil {
			return fmt.Errorf("failed to create wireguard peer removal request: %w", err)
		}

		// Remove the peer for WireGuard service.
		if err := c.Service().RemovePeer(ctx, req); err != nil {
			return fmt.Errorf("failed to remove wireguard peer: %w", err)
		}

		return nil
	default:
		// Return an error for unsupported service types.
		return fmt.Errorf("invalid service type %s", t)
	}
}

// RemovePeerIfExistsForKey checks if a peer exists, and removes it if found.
func (c *Context) RemovePeerIfExistsForKey(ctx context.Context, s string) error {
	// Check if the peer exists.
	ok, err := c.HasPeerForKey(ctx, s)
	if err != nil {
		return fmt.Errorf("failed to check peer existence: %w", err)
	}
	if !ok {
		return nil
	}

	// Remove the peer if it exists.
	if err := c.RemovePeerForKey(ctx, s); err != nil {
		return fmt.Errorf("failed to remove peer: %w", err)
	}

	return nil
}
