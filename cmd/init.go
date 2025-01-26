package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sentinel-official/dvpn-node/config"
)

// NewInitCmd creates and returns a new Cobra command for initializing the application configuration.
func NewInitCmd() *cobra.Command {
	// Initialize default configuration
	cfg := config.DefaultConfig()

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the application configuration",
		Long: `Creates the application home directory and generates a default config.toml file.
If a configuration file already exists, this command will abort unless the "force" flag
is set to overwrite the existing configuration.`,
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
				// If the file exists, check if the `force` flag is set
				force, err := cmd.Flags().GetBool("force")
				if err != nil {
					return fmt.Errorf("failed to get 'force' flag: %w", err)
				}

				if !force {
					return errors.New("config file already exists")
				}
			}

			// Validate the configuration.
			if err := cfg.Validate(); err != nil {
				return fmt.Errorf("invalid configuration: %w", err)
			}

			// Write the default configuration to the configuration file
			if err := cfg.WriteToFile(cfgFile); err != nil {
				return fmt.Errorf("failed to write config file: %w", err)
			}

			cmd.Println("Configuration initialized successfully")
			return nil
		},
	}

	// Set configuration flags using the default configuration.
	cfg.SetForFlags(cmd.Flags())

	// Add a flag to allow overwriting the existing configuration file
	cmd.Flags().Bool("force", false, "overwrite the existing configuration file if it exists")

	return cmd
}
