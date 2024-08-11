package options

import (
	"github.com/spf13/cobra"

	"github.com/sentinel-official/dvpn-node/cmd/flags"
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

// AddQOSFlagsToCmd adds QOS-related flags to the given Cobra command.
func AddQOSFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagQOSMaxPeers(cmd)
}

// NewQOSFromCmd creates and returns QOS from the given Cobra command's flags.
func NewQOSFromCmd(cmd *cobra.Command) (*QOS, error) {
	// Retrieve the max peers flag value from the command.
	maxPeers, err := flags.GetQOSMaxPeersFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new QOS instance populated with the retrieved flag value.
	return &QOS{
		MaxPeers: maxPeers,
	}, nil
}
