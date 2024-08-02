package options

import (
	"github.com/spf13/cobra"
)

// Default values for handshake options.
const (
	DefaultHandshakeEnable = false
	DefaultHandshakePeers  = 10
)

// HandshakeOptions represents options for configuring the Handshake DNS system.
type HandshakeOptions struct {
	Enable bool   `json:"enable" toml:"enable"` // Enable indicates whether the Handshake DNS functionality is enabled.
	Peers  uint64 `json:"peers" toml:"peers"`   // Peers specifies the number of peers to connect with in the Handshake network.
}

// NewDefaultHandshakeOptions creates a new HandshakeOptions instance with default values.
func NewDefaultHandshakeOptions() *HandshakeOptions {
	return &HandshakeOptions{
		Enable: DefaultHandshakeEnable,
		Peers:  DefaultHandshakePeers,
	}
}

// WithEnable sets the Enable field and returns the modified HandshakeOptions instance.
func (h *HandshakeOptions) WithEnable(v bool) *HandshakeOptions {
	h.Enable = v
	return h
}

// WithPeers sets the Peers field and returns the modified HandshakeOptions instance.
func (h *HandshakeOptions) WithPeers(v uint64) *HandshakeOptions {
	h.Peers = v
	return h
}

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
	cmd.Flags().Bool("handshake.enable", DefaultHandshakeEnable, "Enable or disable the Handshake DNS functionality.")
}

// SetFlagHandshakePeers adds the "handshake.peers" flag to the command.
func SetFlagHandshakePeers(cmd *cobra.Command) {
	cmd.Flags().Uint64("handshake.peers", DefaultHandshakePeers, "Number of peers to connect with in the Handshake DNS network.")
}

// AddHandshakeFlagsToCmd adds Handshake-related flags to the given Cobra command.
func AddHandshakeFlagsToCmd(cmd *cobra.Command) {
	SetFlagHandshakeEnable(cmd)
	SetFlagHandshakePeers(cmd)
}

// NewHandshakeOptionsFromCmd creates and returns HandshakeOptions from the given Cobra command's flags.
func NewHandshakeOptionsFromCmd(cmd *cobra.Command) (*HandshakeOptions, error) {
	// Retrieve the value of the "handshake.enable" flag.
	enable, err := GetHandshakeEnableFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "handshake.peers" flag.
	peers, err := GetHandshakePeersFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new HandshakeOptions instance populated with the retrieved flag values.
	return &HandshakeOptions{
		Enable: enable,
		Peers:  peers,
	}, nil
}
