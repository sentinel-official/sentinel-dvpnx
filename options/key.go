package options

import (
	"github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
	"github.com/spf13/cobra"
)

// KeyringOptions represents options for keyring creation.
type KeyringOptions struct {
	Backend string `json:"backend" toml:"backend"`   // Backend is the keyring backend to use.
	HomeDir string `json:"home_dir" toml:"home_dir"` // HomeDir is the directory to store keys.
}

// NewDefaultKeyring creates a new KeyringOptions instance with default values.
func NewDefaultKeyring() *KeyringOptions {
	return &KeyringOptions{
		Backend: flags.DefaultKeyringBackend,
	}
}

// WithBackend sets the Backend field and returns the modified KeyringOptions instance.
func (k *KeyringOptions) WithBackend(v string) *KeyringOptions {
	k.Backend = v
	return k
}

// WithHomeDir sets the HomeDir field and returns the modified KeyringOptions instance.
func (k *KeyringOptions) WithHomeDir(v string) *KeyringOptions {
	k.HomeDir = v
	return k
}

// AddKeyringFlagsToCmd adds keyring-related flags to the given cobra command.
func AddKeyringFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagKeyringBackend(cmd)
	flags.SetFlagKeyringHomeDir(cmd)
}

// NewKeyringOptionsFromCmd creates and returns KeyringOptions from the given cobra command's flags.
func NewKeyringOptionsFromCmd(cmd *cobra.Command) (*KeyringOptions, error) {
	// Retrieve the backend flag value from the command.
	backend, err := flags.GetKeyringBackendFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the home directory flag value from the command.
	homeDir, err := flags.GetKeyringHomeDirFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new KeyringOptions instance populated with the retrieved flag values.
	return &KeyringOptions{
		Backend: backend,
		HomeDir: homeDir,
	}, nil
}
