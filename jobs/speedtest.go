package jobs

import (
	"github.com/sentinel-official/sentinel-go-sdk/libs/speedtest"

	"github.com/sentinel-official/dvpn-node/context"
)

// handleSpeedtest returns a function that performs a speed test and sets the results in the context.
func handleSpeedtest(ctx *context.Context) func() error {
	return func() error {
		// Run the speed test and get the download and upload speeds
		dlSpeed, ulSpeed, err := speedtest.Run()
		if err != nil {
			return err
		}

		// Set the speed test results in the context
		if err := ctx.SetSpeedtestResults(dlSpeed, ulSpeed); err != nil {
			return err
		}

		return nil
	}
}

// onErrorSpeedtest returns a function to handle errors.
func onErrorSpeedtest(_ *context.Context) func(error) bool {
	return func(err error) bool {
		// Do nothing on error
		return false
	}
}
