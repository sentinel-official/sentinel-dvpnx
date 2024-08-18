package flags

import (
	"github.com/spf13/cobra"
)

// Default values for handshake flags.
const (
	DefaultHandshakeEnable = false
	DefaultHandshakePeers  = 10
)

// GetHandshakeEnable retrieves the "handshake.enable" flag value from the command.
func GetHandshakeEnable(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("handshake.enable")
}

// GetHandshakePeers retrieves the "handshake.peers" flag value from the command.
func GetHandshakePeers(cmd *cobra.Command) (uint64, error) {
	return cmd.Flags().GetUint64("handshake.peers")
}

// SetFlagHandshakeEnable adds the "handshake.enable" flag to the command.
func SetFlagHandshakeEnable(cmd *cobra.Command) {
	cmd.Flags().Bool("handshake.enable", DefaultHandshakeEnable, "Enable or disable the Handshake DNS functionality.")
}

// SetFlagHandshakePeers adds the "handshake.peers" flag to the command.
func SetFlagHandshakePeers(cmd *cobra.Command) {
	cmd.Flags().Uint64("handshake.peers", DefaultHandshakePeers, "Number of peers to connect with in the Handshake DNS network.")
}

// AddHandshakeFlags adds Handshake-related flags to the given Cobra command.
func AddHandshakeFlags(cmd *cobra.Command) {
	SetFlagHandshakeEnable(cmd)
	SetFlagHandshakePeers(cmd)
}
