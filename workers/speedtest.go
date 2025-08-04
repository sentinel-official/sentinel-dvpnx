package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/sentinel-official/sentinel-go-sdk/libs/cron"
	"github.com/sentinel-official/sentinel-go-sdk/libs/speedtest"

	"github.com/sentinel-official/sentinel-dvpnx/core"
)

const nameSpeedtest = "speedtest"

// NewSpeedtestWorker creates a worker that performs periodic speed tests and updates the context with the results.
// This worker measures the download and upload speeds and sets the results in the application's context.
func NewSpeedtestWorker(c *core.Context, interval time.Duration) cron.Worker {
	// Handler function that performs the speed test and updates the context.
	handlerFunc := func(ctx context.Context) error {
		// Run the speed test to measure download and upload speeds.
		dlSpeed, ulSpeed, err := speedtest.Run(ctx)
		if err != nil {
			return fmt.Errorf("failed to run speed test: %w", err)
		}

		// Update the context with the obtained speed test results.
		c.SetSpeedtestResults(dlSpeed, ulSpeed)

		return nil
	}

	// Error handling function to log failures.
	onErrorFunc := func(err error) bool {
		return false
	}

	// Initialize and return the worker.
	return cron.NewBasicWorker().
		WithName(nameSpeedtest).
		WithHandler(handlerFunc).
		WithInterval(interval).
		WithOnError(onErrorFunc)
}
