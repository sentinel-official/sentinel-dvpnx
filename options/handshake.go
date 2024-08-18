package options

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/sentinel-official/dvpn-node/flags"
)

// Handshake represents options for configuring the Handshake DNS system.
type Handshake struct {
	Enable bool   `json:"enable" toml:"enable"` // Enable indicates whether the Handshake DNS functionality is enabled.
	Peers  uint64 `json:"peers" toml:"peers"`   // Peers specifies the number of peers to connect with in the Handshake network.
}

// NewHandshake creates a new Handshake instance with default values.
func NewHandshake() *Handshake {
	return &Handshake{
		Enable: flags.DefaultHandshakeEnable,
		Peers:  flags.DefaultHandshakePeers,
	}
}

// WithEnable sets the Enable field and returns the modified Handshake instance.
func (h *Handshake) WithEnable(v bool) *Handshake {
	h.Enable = v
	return h
}

// WithPeers sets the Peers field and returns the modified Handshake instance.
func (h *Handshake) WithPeers(v uint64) *Handshake {
	h.Peers = v
	return h
}

// GetEnable returns the value of the Enable field.
func (h *Handshake) GetEnable() bool {
	return h.Enable
}

// GetPeers returns the value of the Peers field.
func (h *Handshake) GetPeers() uint64 {
	return h.Peers
}

// ValidateHandshakePeers checks if the Peers field is valid.
func ValidateHandshakePeers(v uint64) error {
	if v == 0 {
		return errors.New("peers must be greater than zero")
	}

	return nil
}

// Validate validates all fields of the Handshake struct.
func (h *Handshake) Validate() error {
	if err := ValidateHandshakePeers(h.Peers); err != nil {
		return err
	}

	return nil
}

// NewHandshakeFromCmd creates and returns Handshake from the given Cobra command's flags.
func NewHandshakeFromCmd(cmd *cobra.Command) (*Handshake, error) {
	// Retrieve the enable flag value from the command.
	enable, err := flags.GetHandshakeEnable(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the peers flag value from the command.
	peers, err := flags.GetHandshakePeers(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new Handshake instance populated with the retrieved flag values.
	return &Handshake{
		Enable: enable,
		Peers:  peers,
	}, nil
}
