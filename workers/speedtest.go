package workers

import (
	"github.com/sentinel-official/sentinel-go-sdk/libs/speedtest"

	"github.com/sentinel-official/dvpn-node/node"
)

// HandlerSpeedtest returns a function that performs a speed test and sets the results in the context.
func HandlerSpeedtest(ctx *node.Context) func() error {
	return func() error {
		// Run the speed test and get the download and upload speeds
		dlSpeed, ulSpeed, err := speedtest.Run()
		if err != nil {
			return err
		}

		// Set the speed test results in the context
		ctx.SetSpeedtestResults(dlSpeed, ulSpeed)

		return nil
	}
}
