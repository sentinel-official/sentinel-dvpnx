package options

import (
	"github.com/spf13/cobra"

	"github.com/sentinel-official/dvpn-node/cmd/flags"
)

// QOSOptions represents options for configuring Quality of Service (QOS).
type QOSOptions struct {
	MaxPeers int `json:"max_peers" toml:"max_peers"` // MaxPeers specifies the maximum number of peers for Node.
}

// NewDefaultQOS creates a new QOSOptions instance with default values.
func NewDefaultQOS() *QOSOptions {
	return &QOSOptions{
		MaxPeers: flags.DefaultQOSMaxPeers,
	}
}

// WithMaxPeers sets the MaxPeers field and returns the modified QOSOptions instance.
func (q *QOSOptions) WithMaxPeers(v int) *QOSOptions {
	q.MaxPeers = v
	return q
}

// AddQOSFlagsToCmd adds QOS-related flags to the given Cobra command.
func AddQOSFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagQOSMaxPeers(cmd)
}

// NewQOSOptionsFromCmd creates and returns QOSOptions from the given Cobra command's flags.
func NewQOSOptionsFromCmd(cmd *cobra.Command) (*QOSOptions, error) {
	// Retrieve the max peers flag value from the command.
	maxPeers, err := flags.GetQOSMaxPeersFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new QOSOptions instance populated with the retrieved flag value.
	return &QOSOptions{
		MaxPeers: maxPeers,
	}, nil
}
