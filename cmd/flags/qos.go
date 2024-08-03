package flags

import (
	"github.com/spf13/cobra"
)

// Default values for QOS options.
const (
	DefaultQOSMaxPeers = 200
)

// GetQOSMaxPeersFromCmd retrieves the "qos.max-peers" flag value from the command.
func GetQOSMaxPeersFromCmd(cmd *cobra.Command) (int, error) {
	return cmd.Flags().GetInt("qos.max-peers")
}

// SetFlagQOSMaxPeers adds the "qos.max-peers" flag to the command.
func SetFlagQOSMaxPeers(cmd *cobra.Command) {
	cmd.Flags().Int("qos.max-peers", DefaultQOSMaxPeers, "Maximum number of peers for node.")
}
