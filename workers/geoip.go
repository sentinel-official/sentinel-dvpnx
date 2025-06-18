package workers

import (
	"fmt"
	"time"

	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	logger "github.com/sentinel-official/sentinel-go-sdk/libs/log"

	"github.com/sentinel-official/sentinel-dvpnx/context"
)

const nameGeoIPLocation = "geoip_location"

// NewGeoIPLocationWorker creates a worker to periodically update the GeoIP location in the context.
// This worker fetches the GeoIP location and updates the context at regular intervals.
func NewGeoIPLocationWorker(c *context.Context, interval time.Duration) cron.Worker {
	log := logger.With("name", nameGeoIPLocation)

	// Handler function that fetches the GeoIP location and updates the context.
	handlerFunc := func() error {
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
		log.Error("Failed to run scheduler worker", "msg", err)
		return false
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker().
		WithName(nameGeoIPLocation).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}
