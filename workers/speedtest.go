package workers

import (
	"fmt"
	"time"

	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	logger "github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/sentinel-official/sentinel-go-sdk/libs/speedtest"

	"github.com/sentinel-official/dvpn-node/context"
)

const nameSpeedtest = "speedtest"

// NewSpeedtestWorker creates a worker that performs periodic speed tests and updates the context with the results.
// This worker measures the download and upload speeds and sets the results in the application's context.
func NewSpeedtestWorker(c *context.Context, interval time.Duration) cron.Worker {
	log := logger.With("name", nameSpeedtest)

	// Handler function that performs the speed test and updates the context.
	handlerFunc := func() error {
		log.Info("Running scheduler worker")

		// Run the speed test to measure download and upload speeds.
		dlSpeed, ulSpeed, err := speedtest.Run()
		if err != nil {
			return fmt.Errorf("failed to run speed test: %w", err)
		}

		// Update the context with the obtained speed test results.
		c.SetSpeedtestResults(dlSpeed, ulSpeed)

		return nil
	}

	// Error handling function to log failures.
	onErrorFunc := func(err error) bool {
		log.Error("Failed to run scheduler worker", "msg", err)
		return false
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker().
		WithName(nameSpeedtest).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}
