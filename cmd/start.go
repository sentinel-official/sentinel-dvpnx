package cmd

import (
	"fmt"
	"os/signal"
	"syscall"

	"github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"

	"github.com/sentinel-official/sentinel-dvpnx/config"
	"github.com/sentinel-official/sentinel-dvpnx/core"
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

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Retrieve the home directory from the configuration.
			homeDir := viper.GetString("home")

			// Create and configure the application context.
			c := core.NewContext().
				WithHomeDir(homeDir).
				WithInput(cmd.InOrStdin())

			// Set up the application context.
			if err := c.Setup(cfg); err != nil {
				return fmt.Errorf("failed to setup context: %w", err)
			}

			// Seal the context to prevent further modifications.
			c.Seal()

			// Create and set up the node.
			n := node.New(c)
			if err := n.Setup(cfg); err != nil {
				return fmt.Errorf("failed to setup node: %w", err)
			}

			// Create a context that listens for SIGINT and SIGTERM signals.
			ctx, cancel := signal.NotifyContext(cmd.Context(), syscall.SIGINT, syscall.SIGTERM)
			defer cancel()

			// Create an errgroup with the signal-aware context.
			eg, ctx := errgroup.WithContext(ctx)

			// Launch goroutine to stop the node gracefully.
			eg.Go(func() error {
				// Wait until signal is received
				<-ctx.Done()

				if err := n.Stop(); err != nil {
					log.Error("Failed to stop node", "cause", err)
				}

				return nil
			})

			// Launch goroutine to start the node and wait.
			eg.Go(func() error {
				if err := n.Start(ctx); err != nil {
					return fmt.Errorf("failed to start node: %w", err)
				}
				if err := n.Wait(); err != nil {
					return fmt.Errorf("failed to wait node: %w", err)
				}

				return nil
			})

			// Wait for all goroutines to complete.
			if err := eg.Wait(); err != nil {
				return err
			}

			return nil
		},
	}

	// Set configuration flags with the command.
	cfg.SetForFlags(cmd.Flags())

	return cmd
}
