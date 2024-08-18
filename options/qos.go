package options

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/sentinel-official/dvpn-node/flags"
)

// QOS represents options for configuring Quality of Service (QOS).
type QOS struct {
	MaxPeers int `json:"max_peers" toml:"max_peers"` // MaxPeers specifies the maximum number of peers for Node.
}

// NewQOS creates a new QOS instance with default values.
func NewQOS() *QOS {
	return &QOS{
		MaxPeers: flags.DefaultQOSMaxPeers,
	}
}

// WithMaxPeers sets the MaxPeers field and returns the modified QOS instance.
func (q *QOS) WithMaxPeers(v int) *QOS {
	q.MaxPeers = v
	return q
}

// GetMaxPeers returns the maximum number of peers.
func (q *QOS) GetMaxPeers() int {
	return q.MaxPeers
}

// ValidateQOSMaxPeers checks if the MaxPeers field is valid.
func ValidateQOSMaxPeers(maxPeers int) error {
	if maxPeers <= 0 {
		return errors.New("max peers must be greater than zero")
	}

	return nil
}

// Validate validates all fields of the QOS struct.
func (q *QOS) Validate() error {
	if err := ValidateQOSMaxPeers(q.MaxPeers); err != nil {
		return err
	}

	return nil
}

// NewQOSFromCmd creates and returns QOS from the given Cobra command's flags.
func NewQOSFromCmd(cmd *cobra.Command) (*QOS, error) {
	// Retrieve the max peers flag value from the command.
	maxPeers, err := flags.GetQOSMaxPeers(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new QOS instance populated with the retrieved flag value.
	return &QOS{
		MaxPeers: maxPeers,
	}, nil
}
