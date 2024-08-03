package options

import (
	"github.com/spf13/cobra"

	"github.com/sentinel-official/dvpn-node/cmd/flags"
)

// HandshakeOptions represents options for configuring the Handshake DNS system.
type HandshakeOptions struct {
	Enable bool   `json:"enable" toml:"enable"` // Enable indicates whether the Handshake DNS functionality is enabled.
	Peers  uint64 `json:"peers" toml:"peers"`   // Peers specifies the number of peers to connect with in the Handshake network.
}

// NewDefaultHandshake creates a new HandshakeOptions instance with default values.
func NewDefaultHandshake() *HandshakeOptions {
	return &HandshakeOptions{
		Enable: flags.DefaultHandshakeEnable,
		Peers:  flags.DefaultHandshakePeers,
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

// AddHandshakeFlagsToCmd adds Handshake-related flags to the given Cobra command.
func AddHandshakeFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagHandshakeEnable(cmd)
	flags.SetFlagHandshakePeers(cmd)
}

// NewHandshakeOptionsFromCmd creates and returns HandshakeOptions from the given Cobra command's flags.
func NewHandshakeOptionsFromCmd(cmd *cobra.Command) (*HandshakeOptions, error) {
	// Retrieve the enable flag value from the command.
	enable, err := flags.GetHandshakeEnableFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the peers flag value from the command.
	peers, err := flags.GetHandshakePeersFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new HandshakeOptions instance populated with the retrieved flag values.
	return &HandshakeOptions{
		Enable: enable,
		Peers:  peers,
	}, nil
}
