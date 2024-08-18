package flags

import (
	sdkflags "github.com/sentinel-official/sentinel-go-sdk/flags"
	"github.com/spf13/cobra"
)

// AddTxFlags configures all transaction-related flags for the given command.
func AddTxFlags(cmd *cobra.Command) {
	sdkflags.SetFlagTxChainID(cmd)
	sdkflags.SetFlagTxFeeGranterAddr(cmd)
	sdkflags.SetFlagTxFromName(cmd)
	sdkflags.SetFlagTxGas(cmd)
	sdkflags.SetFlagTxGasAdjustment(cmd)
	sdkflags.SetFlagTxGasPrices(cmd)
	sdkflags.SetFlagTxSimulateAndExecute(cmd)
}
