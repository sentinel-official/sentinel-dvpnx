package flags

import (
	"github.com/spf13/cobra"
)

// Default values for handshake options.
const (
	DefaultHandshakePeers = 10
)

// GetHandshakeEnableFromCmd retrieves the "handshake.enable" flag value from the command.
func GetHandshakeEnableFromCmd(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("handshake.enable")
}

// GetHandshakePeersFromCmd retrieves the "handshake.peers" flag value from the command.
func GetHandshakePeersFromCmd(cmd *cobra.Command) (uint64, error) {
	return cmd.Flags().GetUint64("handshake.peers")
}

// SetFlagHandshakeEnable adds the "handshake.enable" flag to the command.
func SetFlagHandshakeEnable(cmd *cobra.Command) {
	cmd.Flags().Bool("handshake.enable", false, "Enable or disable the Handshake DNS functionality.")
}

// SetFlagHandshakePeers adds the "handshake.peers" flag to the command.
func SetFlagHandshakePeers(cmd *cobra.Command) {
	cmd.Flags().Uint64("handshake.peers", DefaultHandshakePeers, "Number of peers to connect with in the Handshake DNS network.")
}
