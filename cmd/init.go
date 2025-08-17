package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sentinel-official/sentinel-go-sdk/libs/crypto"
	"github.com/sentinel-official/sentinel-go-sdk/libs/log"
	"github.com/sentinel-official/sentinel-go-sdk/openvpn"
	"github.com/sentinel-official/sentinel-go-sdk/types"
	"github.com/sentinel-official/sentinel-go-sdk/utils"
	"github.com/sentinel-official/sentinel-go-sdk/v2ray"
	"github.com/sentinel-official/sentinel-go-sdk/wireguard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sentinel-official/sentinel-dvpnx/config"
)

// NewInitCmd creates and returns a new Cobra command for initializing the application configuration.
func NewInitCmd(cfg *config.Config) *cobra.Command {
	// Declare variables for CLI flags
	var (
		force       bool // Determines whether to overwrite existing config
		skipTLS     bool // Whether to skip the TLS key/cert generation
		skipService bool // Whether to skip the node service initialization
	)

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the application configuration",
		Long: `Creates the application home directory and generates a default config.toml file.
If a configuration file already exists, this command will abort unless the "force" flag
is set to overwrite the existing configuration.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			log.Info("Validating configuration")
			if err := cfg.Validate(); err != nil {
				return fmt.Errorf("validating configuration: %w", err)
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Create the home directory if it doesn't exist
			homeDir := viper.GetString("home")
			if err := os.MkdirAll(homeDir, 0755); err != nil {
				return fmt.Errorf("creating application directory %q: %w", homeDir, err)
			}

			// Construct the full path to the config file
			cfgFile := filepath.Join(homeDir, "config.toml")

			// Check if the config file exists at the specified path
			cfgFileExists, err := utils.IsFileExists(cfgFile)
			if err != nil {
				return fmt.Errorf("checking if config file %q exists: %w", cfgFile, err)
			}

			// Write default config only if file doesn't exist or force flag is set
			if !cfgFileExists || force {
				log.Info("Writing app config", "file", cfgFile)
				if err := cfg.WriteAppConfig(cfgFile); err != nil {
					return fmt.Errorf("writing config file %q: %w", cfgFile, err)
				}
			}

			// Generate TLS keys if "skipTLS" is disabled
			if !skipTLS {
				log.Info("Initializing PKI with CA certificate and key", "dir", homeDir)

				pki := crypto.NewPKI(homeDir)
				if err := pki.Init(); err != nil {
					return fmt.Errorf("initializing PKI: %w", err)
				}

				opts := []crypto.CertOption{
					crypto.CertSAN(cfg.Node.GetRemoteAddrs()...),
				}

				log.Info("Issuing certificate and key", "name", "tls")
				if _, _, err := pki.Issue("tls", opts...); err != nil {
					return fmt.Errorf("issuing TLS certificate and key: %w", err)
				}
			}

			// Initialize the node service config if "skipService" is disabled
			if !skipService {
				var (
					service     types.ServerService         // Interface for the node service
					serviceType = cfg.Node.GetServiceType() // Get the service type from config
				)

				log.Info("Initializing service", "type", serviceType)

				switch serviceType {
				case types.ServiceTypeV2Ray:
					service = v2ray.NewServer(homeDir)
				case types.ServiceTypeWireGuard:
					service = wireguard.NewServer(homeDir)
				case types.ServiceTypeOpenVPN:
					service = openvpn.NewServer(homeDir)
				default:
					return fmt.Errorf("unsupported service type %q", serviceType)
				}

				log.Info("Running service init task")
				if err := service.Init(force); err != nil {
					return fmt.Errorf("running service init task: %w", err)
				}
			}

			log.Info("Configuration initialized successfully")
			return nil
		},
	}

	// Set default configuration flags
	cfg.SetForFlags(cmd.Flags())

	// Bind command-line flags to local variables
	cmd.Flags().BoolVar(&force, "force", force, "overwrite the existing configuration file if it exists")
	cmd.Flags().BoolVar(&skipTLS, "skip-tls", false, "skip TLS key and certificate generation")
	cmd.Flags().BoolVar(&skipService, "skip-service", false, "skip initialization of the selected service")

	return cmd
}
