package options

import (
	"github.com/spf13/cobra"
)

// Default values for QOS options.
const (
	DefaultQOSMaxPeers = 250
)

// QOSOptions represents options for configuring Quality of Service (QOS).
type QOSOptions struct {
	MaxPeers int `json:"max_peers" toml:"max_peers"` // MaxPeers specifies the maximum number of peers for Node.
}

// NewDefaultQOSOptions creates a new QOSOptions instance with default values.
func NewDefaultQOSOptions() *QOSOptions {
	return &QOSOptions{
		MaxPeers: DefaultQOSMaxPeers,
	}
}

// WithMaxPeers sets the MaxPeers field and returns the modified QOSOptions instance.
func (q *QOSOptions) WithMaxPeers(v int) *QOSOptions {
	q.MaxPeers = v
	return q
}

// GetQOSMaxPeersFromCmd retrieves the "qos.max-peers" flag value from the command.
func GetQOSMaxPeersFromCmd(cmd *cobra.Command) (int, error) {
	return cmd.Flags().GetInt("qos.max-peers")
}

// SetFlagQOSMaxPeers adds the "qos.max-peers" flag to the command.
func SetFlagQOSMaxPeers(cmd *cobra.Command) {
	cmd.Flags().Int("qos.max-peers", DefaultQOSMaxPeers, "Maximum number of peers for node.")
}

// AddQOSFlagsToCmd adds QOS-related flags to the given Cobra command.
func AddQOSFlagsToCmd(cmd *cobra.Command) {
	SetFlagQOSMaxPeers(cmd)
}

// NewQOSOptionsFromCmd creates and returns QOSOptions from the given Cobra command's flags.
func NewQOSOptionsFromCmd(cmd *cobra.Command) (*QOSOptions, error) {
	// Retrieve the value of the "qos.max-peers" flag.
	maxPeers, err := GetQOSMaxPeersFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new QOSOptions instance populated with the retrieved flag value.
	return &QOSOptions{
		MaxPeers: maxPeers,
	}, nil
}
