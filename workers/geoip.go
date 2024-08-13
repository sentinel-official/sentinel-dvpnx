package workers

import (
	"github.com/sentinel-official/dvpn-node/node"
)

// HandlerGeoIPLocation fetches the geo IP location using the provided client and sets it in the given context.
func HandlerGeoIPLocation(ctx *node.Context) func() error {
	return func() error {
		// Fetch the geo IP location.
		location, err := ctx.GeoIPClient().Get("")
		if err != nil {
			return err
		}

		// Set the fetched location in the context.
		ctx.SetLocation(location)

		return nil
	}
}
