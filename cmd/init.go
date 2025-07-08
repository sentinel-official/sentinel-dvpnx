package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sentinel-official/sentinel-dvpnx/config"
)

// NewInitCmd creates and returns a new Cobra command for initializing the application configuration.
func NewInitCmd(cfg *config.Config) *cobra.Command {
	// Declare variables for flags
	var force bool
	var withTLS bool

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the application configuration",
		Long: `Creates the application home directory and generates a default config.toml file.
If a configuration file already exists, this command will abort unless the "force" flag
is set to overwrite the existing configuration.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Validate the configuration.
			if err := cfg.Validate(); err != nil {
				return fmt.Errorf("invalid configuration: %w", err)
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Create the home directory if it doesn't exist
			homeDir := viper.GetString("home")
			if err := os.MkdirAll(homeDir, 0755); err != nil {
				return fmt.Errorf("failed to create application directory: %w", err)
			}

			// Construct the config file path.
			cfgFile := filepath.Join(homeDir, "config.toml")

			// Check if the configuration file already exists
			if _, err := os.Stat(cfgFile); err != nil {
				// If an error other than "file not found" occurs, return it
				if !os.IsNotExist(err) {
					return fmt.Errorf("failed to stat config file: %w", err)
				}
			} else {
				if !force {
					return errors.New("config file already exists")
				}
			}

			// Write the default configuration to the configuration file
			if err := cfg.WriteToFile(cfgFile); err != nil {
				return fmt.Errorf("failed to write config file: %w", err)
			}

			// Generate TLS keys if "withTLS" is enabled
			if withTLS {
				// todo: generate tls key and certificate
			}

			cmd.Println("Configuration initialized successfully")
			return nil
		},
	}

	// Set configuration flags using the default configuration.
	cfg.SetForFlags(cmd.Flags())

	// Bind command-line flags to variables
	cmd.Flags().BoolVar(&force, "force", force, "overwrite the existing configuration file if it exists")
	cmd.Flags().BoolVar(&withTLS, "with-tls", withTLS, "generate tls keys for secure communication")

	return cmd
}
