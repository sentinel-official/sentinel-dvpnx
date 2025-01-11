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

// InitCmd returns a new Cobra command for initializing the application configuration.
func InitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the application configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Retrieve Viper instance for managing configuration
			v := viper.GetViper()

			// Get the home directory path from the configuration
			homeDir := v.GetString("home")

			// Create the home directory if it doesn't exist
			if err := os.MkdirAll(homeDir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", homeDir, err)
			}

			// Define the path to the configuration file
			cfgFile := filepath.Join(homeDir, "config.toml")

			// Check if the configuration file already exists
			if _, err := os.Stat(cfgFile); err != nil {
				// If an error other than "file not found" occurs, return it
				if !os.IsNotExist(err) {
					return fmt.Errorf("failed to stat file %s: %w", cfgFile, err)
				}
			} else {
				// If the file exists, check if the `force` flag is set
				force := v.GetBool("force")
				if !force {
					return errors.New("config file already exists")
				}
			}

			// Generate the default configuration
			cfg := config.DefaultConfig()

			// Write the default configuration to the configuration file
			if err := cfg.WriteToFile(cfgFile); err != nil {
				return fmt.Errorf("failed to write file %s: %w", cfgFile, err)
			}

			return nil
		},
	}

	// Add a flag to allow overwriting the existing configuration file
	cmd.Flags().Bool("force", false, "overwrite the existing configuration file if it exists")

	return cmd
}
