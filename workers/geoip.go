package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"

	"github.com/sentinel-official/sentinel-dvpnx/core"
)

const nameGeoIPLocation = "geoip_location"

// NewGeoIPLocationWorker creates a worker to periodically update the GeoIP location in the context.
// This worker fetches the GeoIP location and updates the context at regular intervals.
func NewGeoIPLocationWorker(c *core.Context, interval time.Duration) cron.Worker {
	// Handler function that fetches the GeoIP location and updates the context.
	handlerFunc := func(_ context.Context) error {
		// Fetch the GeoIP location using the GeoIP client.
		location, err := c.GeoIPClient().Get("")
		if err != nil {
			return fmt.Errorf("failed to get geoip location: %w", err)
		}

		// Update the context with the fetched location.
		c.SetLocation(location)

		return nil
	}

	// Error handling function to log failures.
	onErrorFunc := func(err error) bool {
		return false
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker().
		WithName(nameGeoIPLocation).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}
