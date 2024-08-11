package options

import (
	sdkflags "github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
	"github.com/spf13/cobra"
)

// Keyring represents options for keyring creation.
type Keyring struct {
	Backend string `json:"backend" toml:"backend"` // Backend is the keyring backend to use.
}

// NewKeyring creates a new Keyring instance with default values.
func NewKeyring() *Keyring {
	return &Keyring{
		Backend: sdkflags.DefaultKeyringBackend,
	}
}

// WithBackend sets the Backend field and returns the modified Keyring instance.
func (k *Keyring) WithBackend(v string) *Keyring {
	k.Backend = v
	return k
}

// AddKeyringFlagsToCmd adds keyring-related flags to the given cobra command.
func AddKeyringFlagsToCmd(cmd *cobra.Command) {
	sdkflags.SetFlagKeyringBackend(cmd)
}

// NewKeyringFromCmd creates and returns Keyring from the given cobra command's flags.
func NewKeyringFromCmd(cmd *cobra.Command) (*Keyring, error) {
	// Retrieve the backend flag value from the command.
	backend, err := sdkflags.GetKeyringBackendFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new Keyring instance populated with the retrieved flag values.
	return &Keyring{
		Backend: backend,
	}, nil
}
