package workers

import (
	nodecontext "github.com/sentinel-official/dvpn-node/context"
)

// HandlerGeoIPLocation fetches the geo IP location using the provided client and sets it in the given context.
func HandlerGeoIPLocation(ctx *nodecontext.Context) func() error {
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
