package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sentinel-official/sentinel-dvpnx/config"
	"github.com/sentinel-official/sentinel-dvpnx/context"
	"github.com/sentinel-official/sentinel-dvpnx/node"
)

// NewStartCmd creates and returns a new Cobra command to start the node application.
func NewStartCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the Sentinel dVPN node",
		Long: `Starts the Sentinel dVPN node. Initializes the logger, sets up the context and node,
explicitly starts the node, and handles SIGINT/SIGTERM for graceful shutdown.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Validate the provided configuration.
			if err := cfg.Validate(); err != nil {
				return fmt.Errorf("invalid configuration: %w", err)
			}

			// Initialize the logger based on the configuration.
			logger, err := log.NewLogger(cmd.OutOrStderr(), cfg.Log.GetFormat(), cfg.Log.GetLevel())
			if err != nil {
				return fmt.Errorf("failed to initialize logger: %w", err)
			}

			// Set the global logger for the application.
			log.SetLogger(logger)
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Retrieve the home directory from the configuration.
			homeDir := viper.GetString("home")

			// Create and configure the application context.
			ctx := context.New().
				WithHomeDir(homeDir).
				WithInput(cmd.InOrStdin())

			// Setup the application context.
			if err := ctx.Setup(cfg); err != nil {
				return fmt.Errorf("failed to setup context: %w", err)
			}

			// Seal the context to prevent further modifications.
			ctx.Seal()

			// Create and setup the node.
			n := node.New(ctx)
			if err := n.Setup(cfg); err != nil {
				return fmt.Errorf("failed to setup node: %w", err)
			}

			// Channel to capture errors from the node.
			errChan := make(chan error, 1)

			// Start the node and handle any startup errors.
			if err := n.Start(errChan); err != nil {
				return fmt.Errorf("failed to start node: %w", err)
			}

			log.Info("Node started successfully")

			// Channel to capture OS signals (SIGINT, SIGTERM).
			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

			// Wait for a signal or an error.
			select {
			case <-signalChan:
			case err := <-errChan:
				log.Error(err.Error())
			}

			// Stop the node gracefully.
			if err := n.Stop(); err != nil {
				return fmt.Errorf("failed to stop node: %w", err)
			}

			log.Info("Node stopped successfully")
			return nil
		},
	}

	// Set configuration flags with the command.
	cfg.SetForFlags(cmd.Flags())

	return cmd
}
