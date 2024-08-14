package cmd

import (
	sdklog "github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/spf13/cobra"

	nodecontext "github.com/sentinel-official/dvpn-node/context"
	"github.com/sentinel-official/dvpn-node/node"
	"github.com/sentinel-official/dvpn-node/options"
)

// StartCmd creates a new Cobra command for starting the node application.
func StartCmd(appName string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Retrieve the home directory from the command flags.
			homeDir, err := cmd.Flags().GetString("home-dir")
			if err != nil {
				return err
			}

			// Initialize the logger from the command.
			log, err := sdklog.NewLoggerFromCmd(cmd)
			if err != nil {
				return err
			}

			// Create options from the command.
			opts, err := options.NewFromCmd(cmd)
			if err != nil {
				return err
			}
			if err := opts.Validate(); err != nil {
				return err
			}

			// Create and configure the application context.
			ctx := nodecontext.New().
				WithAppName(appName).
				WithHomeDir(homeDir).
				WithLogger(log).
				WithQueryRPCAddrs(opts.Query.RPCAddrs).
				WithKeyringBackend(opts.Keyring.Backend).
				WithNodeGigabytePrices(opts.Node.GigabytePrices).
				WithNodeHourlyPrices(opts.Node.HourlyPrices).
				WithNodeIntervalBestRPCAddr(opts.Node.IntervalBestRPCAddr).
				WithNodeIntervalGeoIPLocation(opts.Node.IntervalGeoIPLocation).
				WithNodeIntervalSessionUsageSyncWithBlockchain(opts.Node.IntervalSessionUsageSyncWithBlockchain).
				WithNodeIntervalSessionUsageSyncWithDatabase(opts.Node.IntervalSessionUsageSyncWithDatabase).
				WithNodeIntervalSessionUsageValidate(opts.Node.IntervalSessionUsageValidate).
				WithNodeIntervalSessionValidate(opts.Node.IntervalSessionValidate).
				WithNodeIntervalSpeedtest(opts.Node.IntervalSpeedtest).
				WithNodeIntervalStatusUpdate(opts.Node.IntervalStatusUpdate).
				WithNodeListenOn(opts.Node.ListenOn).
				WithNodeMoniker(opts.Node.Moniker).
				WithNodeRemoteURL(opts.Node.RemoteURL).
				WithNodeType(opts.Node.Type).
				WithQueryMaxRetries(opts.Query.MaxRetries).
				WithQueryTimeout(opts.Query.Timeout).
				WithTxChainID(opts.Tx.ChainID).
				WithTxFeeGranterAddr(opts.Tx.FeeGranterAddr).
				WithTxFromName(opts.Tx.FromName).
				WithTxGas(opts.Tx.Gas).
				WithTxGasAdjustment(opts.Tx.GasAdjustment).
				WithTxGasPrices(opts.Tx.GasPrices).
				WithTxSimulateAndExecute(opts.Tx.SimulateAndExecute)

			// Setup the application context.
			if err := ctx.Setup(cmd.InOrStdin()); err != nil {
				return err
			}

			// Seal the context to prevent further modifications.
			ctx.Seal()

			// Create and initialize the node.
			n := node.New(ctx)

			// Setup router and scheduler for the node.
			if err := n.SetupRouter(); err != nil {
				return err
			}
			if err := n.SetupScheduler(); err != nil {
				return err
			}

			// Start the node.
			if err := n.Start(); err != nil {
				return err
			}

			return nil
		},
	}

	options.AddHandshakeFlagsToCmd(cmd)
	options.AddKeyringFlagsToCmd(cmd)
	options.AddLogFlagsToCmd(cmd)
	options.AddNodeFlagsToCmd(cmd)
	options.AddQOSFlagsToCmd(cmd)
	options.AddQueryFlagsToCmd(cmd)
	options.AddTxFlagsToCmd(cmd)
	cmd.Flags().String("home-dir", "", "Directory for configuration and application data storage.")

	return cmd
}
