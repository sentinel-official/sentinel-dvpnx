package cmd

import (
	sdkflags "github.com/sentinel-official/sentinel-go-sdk/flags"
	sdklog "github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/spf13/cobra"

	nodecontext "github.com/sentinel-official/dvpn-node/context"
	"github.com/sentinel-official/dvpn-node/flags"
	"github.com/sentinel-official/dvpn-node/node"
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
			opts, err := node.NewOptionsFromCmd(cmd)
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
				WithGigabytePrices(opts.GetGigabytePrices()).
				WithHourlyPrices(opts.GetHourlyPrices()).
				WithIntervalBestRPCAddr(opts.GetIntervalBestRPCAddr()).
				WithIntervalGeoIPLocation(opts.GetIntervalGeoIPLocation()).
				WithIntervalSessionUsageSyncWithBlockchain(opts.GetIntervalSessionUsageSyncWithBlockchain()).
				WithIntervalSessionUsageSyncWithDatabase(opts.GetIntervalSessionUsageSyncWithDatabase()).
				WithIntervalSessionUsageValidate(opts.GetIntervalSessionUsageValidate()).
				WithIntervalSessionValidate(opts.GetIntervalSessionValidate()).
				WithIntervalSpeedtest(opts.GetIntervalSpeedtest()).
				WithIntervalStatusUpdate(opts.GetIntervalStatusUpdate()).
				WithKeyringBackend(opts.Keyring.GetBackend()).
				WithListenOn(opts.GetListenOn()).
				WithMoniker(opts.GetMoniker()).
				WithNodeType(opts.GetType()).
				WithQueryMaxRetries(opts.Query.GetMaxRetries()).
				WithQueryRPCAddrs(opts.Query.GetRPCAddrs()).
				WithQueryTimeout(opts.Query.GetTimeout()).
				WithRemoteURL(opts.GetRemoteURL()).
				WithTxChainID(opts.Tx.GetChainID()).
				WithTxFeeGranterAddr(opts.Tx.GetFeeGranterAddr()).
				WithTxFromName(opts.Tx.GetFromName()).
				WithTxGasAdjustment(opts.Tx.GetGasAdjustment()).
				WithTxGas(opts.Tx.GetGas()).
				WithTxGasPrices(opts.Tx.GetGasPrices()).
				WithTxSimulateAndExecute(opts.Tx.GetSimulateAndExecute())

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

	flags.AddHandshakeFlags(cmd)
	flags.AddKeyringFlags(cmd)
	sdkflags.AddLogFlags(cmd)
	flags.AddNodeFlags(cmd)
	flags.AddQOSFlags(cmd)
	flags.AddQueryFlags(cmd)
	flags.AddTxFlags(cmd)
	cmd.Flags().String("home-dir", "", "Directory for configuration and application data storage.")

	return cmd
}
