package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sentinel-official/sentinel-go-sdk/cmd"
	"github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/sentinel-official/sentinel-go-sdk/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/sentinel-official/sentinel-dvpnx/config"
)

// NewRootCmd returns the main/root command for the Sentinel dVPN node CLI.
func NewRootCmd(userDir string) *cobra.Command {
	// Declare variables for CLI flags
	var (
		homeDir   = filepath.Join(userDir, ".sentinel-dvpnx")
		logFormat = "text"
		logLevel  = "info"
	)

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
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// Initialize logger with selected format and level
			logger, err := log.NewLogger(cmd.OutOrStdout(), logFormat, logLevel)
			if err != nil {
				return fmt.Errorf("initializing logger: %w", err)
			}

			// Set the global logger instance
			log.SetLogger(logger)

			// Create a new viper instance
			v := viper.New()

			// Bind flags to Viper with normalized keys
			r := strings.NewReplacer("-", "_")
			cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
				_ = v.BindPFlag(r.Replace(f.Name), f)
			})
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				_ = v.BindPFlag(r.Replace(f.Name), f)
			})

			// Construct the full path to the config file
			cfgFile := filepath.Join(homeDir, "config.toml")

			// Check if the config file exists at the specified path
			exists, err := utils.IsFileExists(cfgFile)
			if err != nil {
				return fmt.Errorf("checking if config file %q exists: %w", cfgFile, err)
			}

			// If the config file exists, proceed to read its contents
			if exists {
				v.SetConfigFile(cfgFile)
				if err := v.ReadInConfig(); err != nil {
					return fmt.Errorf("reading config file %q: %w", cfgFile, err)
				}
			}

			// Unmarshal configuration into the config object
			if err := v.Unmarshal(cfg); err != nil {
				return fmt.Errorf("unmarshaling config: %w", err)
			}

			// Update the keyring configuration
			cfg.Keyring.HomeDir = homeDir
			cfg.Keyring.Input = cmd.InOrStdin()

			log.Info("Validating configuration")
			if err := cfg.Validate(); err != nil {
				return fmt.Errorf("validating config: %w", err)
			}

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
	rootCmd.PersistentFlags().StringVar(&logFormat, "log.format", logFormat, "format of the log output (json or text)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log.level", logLevel, "log level for output (debug, error, info, none, warn)")

	// Bind flags to global viper instance
	_ = viper.BindPFlag("home", rootCmd.PersistentFlags().Lookup("home"))
	_ = viper.BindPFlag("log.format", rootCmd.PersistentFlags().Lookup("log.format"))
	_ = viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log.level"))

	return rootCmd
}
