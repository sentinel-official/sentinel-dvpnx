package cmd

import (
	"fmt"

	"github.com/sentinel-official/sentinel-go-sdk/app"
	"github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/sentinel-official/sentinel-go-sdk/openvpn"
	"github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/v2ray"
	"github.com/sentinel-official/sentinel-go-sdk/wireguard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"

	"github.com/sentinel-official/sentinel-dvpnx/config"
	"github.com/sentinel-official/sentinel-dvpnx/core"
	"github.com/sentinel-official/sentinel-dvpnx/node"
)

// NewStartCmd creates and returns a new Cobra command to start the node application.
func NewStartCmd(cfg *config.Config) *cobra.Command {
	// Initialize default server configs for all supported services
	cfg.Services = map[types.ServiceType]types.ServiceConfig{
		types.ServiceTypeOpenVPN:   openvpn.DefaultServerConfig(),
		types.ServiceTypeV2Ray:     v2ray.DefaultServerConfig(),
		types.ServiceTypeWireGuard: wireguard.DefaultServerConfig(),
	}

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the Sentinel dVPN node",
		Long: `Starts the Sentinel dVPN node. Initializes the logger, sets up the context and node,
explicitly starts the node, and handles SIGINT/SIGTERM for graceful shutdown.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Retrieve the home directory from the configuration
			homeDir := viper.GetString("home")

			// Create and configure the application context
			c := core.NewContext().
				WithHomeDir(homeDir).
				WithInput(cmd.InOrStdin())

			log.Info("Setting up context")
			if err := c.Setup(cfg); err != nil {
				return fmt.Errorf("setting up context: %w", err)
			}

			// Seal the context to prevent further modifications
			c.Seal()

			// Create and initialize the node with the configured context
			n := node.New(c)

			log.Info("Setting up node")
			if err := n.Setup(cfg); err != nil {
				return fmt.Errorf("setting up node: %w", err)
			}

			// Register the node and update its details
			if err := n.Register(cmd.Context()); err != nil {
				return fmt.Errorf("registering node: %w", err)
			}
			if err := n.UpdateDetails(cmd.Context()); err != nil {
				return fmt.Errorf("updating node details: %w", err)
			}

			// Use errgroup to manage concurrent start/wait and shutdown operations
			eg, ctx := errgroup.WithContext(cmd.Context())

			// Goroutine to handle graceful shutdown on signal
			eg.Go(func() error {
				<-ctx.Done()

				log.Info("Stopping node")
				if err := n.Stop(); err != nil {
					return &app.StopError{Err: err}
				}

				log.Info("Node stopped successfully")
				return nil
			})

			// Goroutine to start and wait on the node
			eg.Go(func() error {
				log.Info("Starting node")
				if err := n.Start(); err != nil {
					return &app.StartError{Err: err}
				}

				log.Info("Node started successfully")
				eg.Go(func() error {
					if err := n.Wait(); err != nil {
						return &app.WaitError{Err: err}
					}

					return nil
				})

				return nil
			})

			// Wait for all goroutines to finish
			if err := eg.Wait(); err != nil {
				return err
			}

			return nil
		},
	}

	// Set CLI flags for application and service configuration
	cfg.SetForFlags(cmd.Flags())
	cfg.Services[types.ServiceTypeOpenVPN].SetForFlags(cmd.Flags(), "openvpn")
	cfg.Services[types.ServiceTypeV2Ray].SetForFlags(cmd.Flags(), "v2ray")
	cfg.Services[types.ServiceTypeWireGuard].SetForFlags(cmd.Flags(), "wireguard")

	return cmd
}
