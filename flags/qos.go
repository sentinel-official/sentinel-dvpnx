package flags

import (
	"github.com/spf13/cobra"
)

// Default values for QOS flags.
const (
	DefaultQOSMaxPeers = 200
)

// GetQOSMaxPeers retrieves the "qos.max-peers" flag value from the command.
func GetQOSMaxPeers(cmd *cobra.Command) (int, error) {
	return cmd.Flags().GetInt("qos.max-peers")
}

// SetFlagQOSMaxPeers adds the "qos.max-peers" flag to the command.
func SetFlagQOSMaxPeers(cmd *cobra.Command) {
	cmd.Flags().Int("qos.max-peers", DefaultQOSMaxPeers, "Maximum number of peers for node.")
}

// AddQOSFlags adds QOS-related flags to the given Cobra command.
func AddQOSFlags(cmd *cobra.Command) {
	SetFlagQOSMaxPeers(cmd)
}
