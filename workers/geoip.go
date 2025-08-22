package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"

	"github.com/sentinel-official/sentinel-dvpnx/core"
)

const NameGeoIPLocation = "geoip_location"

// NewGeoIPLocationWorker creates a worker to periodically update the GeoIP location in the context.
// This worker fetches the GeoIP location and updates the context at regular intervals.
func NewGeoIPLocationWorker(c *core.Context, interval time.Duration) cron.Worker {
	// Handler function that fetches the GeoIP location and updates the context.
	handlerFunc := func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		// Fetch the GeoIP location using the GeoIP client.
		location, err := c.GeoIPClient().Get(ctx, "")
		if err != nil {
			return fmt.Errorf("getting GeoIP location: %w", err)
		}

		// Update the context with the fetched location.
		c.SetLocation(location)

		return nil
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker(NameGeoIPLocation).
		WithHandler(handlerFunc).
		WithInterval(interval)
}
