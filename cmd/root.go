package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sentinel-official/sentinel-go-sdk/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/sentinel-official/dvpn-node/config"
)

// NewRootCmd returns the main/root command for the Sentinel dVPN node CLI.
func NewRootCmd(homeDir string) *cobra.Command {
	// Initialize default configuration
	cfg := config.DefaultConfig()

	// Create the root command
	rootCmd := &cobra.Command{
		Use:   "sentinel-dvpnx",
		Short: "Run and manage the Sentinel dVPN node",
		Long: `The Sentinel dVPN node software lets users join the decentralized VPN network on the Sentinel Hub
blockchain, providing secure, private, and censorship-resistant internet access while earning
cryptocurrency rewards. It integrates with Cosmos-SDK, supports robust configuration, and offers
tools for key management and node initialization, ensuring privacy, performance, and ease of use.`,
		SilenceUsage: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// Initialize viper instance
			v := viper.New()

			// Skip loading if the config file does not exist
			cfgPath := filepath.Join(homeDir, "config.toml")
			if _, err := os.Stat(cfgPath); err != nil {
				if !os.IsNotExist(err) {
					return fmt.Errorf("failed to stat config file: %w", err)
				}
			} else {
				// Read the config from the specified file
				v.SetConfigFile(cfgPath)
				if err := v.ReadInConfig(); err != nil {
					return fmt.Errorf("failed to read config file: %w", err)
				}
			}

			// Bind flags to Viper with normalized keys
			r := strings.NewReplacer("-", "_")
			cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
				_ = v.BindPFlag(r.Replace(f.Name), f)
			})
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				_ = v.BindPFlag(r.Replace(f.Name), f)
			})

			// Unmarshal configuration into the config object
			if err := v.Unmarshal(cfg); err != nil {
				return fmt.Errorf("failed to unmarshal config file: %w", err)
			}

			// Write the updated configuration to the file
			if err := cfg.WriteToFile(cfgPath); err != nil {
				return fmt.Errorf("failed to write config file: %w", err)
			}

			// Update the keyring configuration
			cfg.Keyring.HomeDir = homeDir
			cfg.Keyring.Input = cmd.InOrStdin()

			// Set TLS paths in the configuration
			cfg.Node.TLSCertPath = filepath.Join(homeDir, "tls.crt")
			cfg.Node.TLSKeyPath = filepath.Join(homeDir, "tls.key")

			return nil
		},
	}

	// Add subcommands
	rootCmd.AddCommand(
		cmd.NewKeysCmd(cfg.Keyring),
		cmd.NewVersionCmd(),
		NewInitCmd(cfg),
		NewStartCmd(cfg),
	)

	// Add persistent flags
	rootCmd.PersistentFlags().StringVar(&homeDir, "home", homeDir, "home directory for application config and data")
	_ = viper.BindPFlag("home", rootCmd.PersistentFlags().Lookup("home"))

	return rootCmd
}
