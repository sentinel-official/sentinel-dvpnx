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
	t := c.Service().Type()

	switch t {
	case types.ServiceTypeV2Ray:
		req, err := v2ray.NewHasPeerRequestFromKey(s)
		if err != nil {
			return false, err
		}

		// Check if the peer exists for V2Ray service
		return c.Service().HasPeer(ctx, req)
	case types.ServiceTypeWireGuard:
		req, err := wireguard.NewHasPeerRequestFromKey(s)
		if err != nil {
			return false, err
		}

		// Check if the peer exists for WireGuard service
		return c.Service().HasPeer(ctx, req)
	default:
		return false, fmt.Errorf("invalid service type: %s", t)
	}
}

// RemovePeerForKey removes the peer with the given key from the current service.
func (c *Context) RemovePeerForKey(ctx context.Context, s string) error {
	t := c.Service().Type()

	switch t {
	case types.ServiceTypeV2Ray:
		req, err := v2ray.NewRemovePeerRequestFromKey(s)
		if err != nil {
			return err
		}

		// Remove the peer for V2Ray service
		return c.Service().RemovePeer(ctx, req)
	case types.ServiceTypeWireGuard:
		req, err := wireguard.NewRemovePeerRequestFromKey(s)
		if err != nil {
			return err
		}

		// Remove the peer for WireGuard service
		return c.Service().RemovePeer(ctx, req)
	default:
		return fmt.Errorf("invalid service type: %s", t)
	}
}

// RemovePeerIfExistsForKey first checks if a peer with the given key exists. If it does, the peer is removed.
func (c *Context) RemovePeerIfExistsForKey(ctx context.Context, s string) error {
	ok, err := c.HasPeerForKey(ctx, s)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	if err := c.RemovePeerForKey(ctx, s); err != nil {
		return err
	}

	return nil
}
