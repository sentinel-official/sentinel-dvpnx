package options

import (
	sdkoptions "github.com/sentinel-official/sentinel-go-sdk/client/options"
	"github.com/spf13/cobra"
)

// Options aggregates all the individual option structs.
type Options struct {
	*HandshakeOptions      `json:"handshake" toml:"handshake"` // Options related to key creation.
	*KeyringOptions        `json:"keyring" toml:"keyring"`     // Options related to keyring configuration.
	*sdkoptions.LogOptions `json:"log" toml:"log"`             // Options related to logging.
	*NodeOptions           `json:"node" toml:"node"`           // Options related to pagination.
	*QOSOptions            `json:"qos" toml:"qos"`             // Options related to quality of service (QOS).
	*QueryOptions          `json:"query" toml:"query"`         // Options related to querying.
	*TxOptions             `json:"tx" toml:"tx"`               // Options related to transactions.
}

// New creates and returns a new instance of Options with all fields initialized to nil.
func New() *Options {
	return &Options{}
}

// NewDefault creates and returns a new instance of Options with default values.
func NewDefault() *Options {
	return &Options{
		HandshakeOptions: NewDefaultHandshake(),      // Initializes with default HandshakeOptions.
		KeyringOptions:   NewDefaultKeyring(),        // Initializes with default KeyringOptions.
		LogOptions:       sdkoptions.NewDefaultLog(), // Initializes with default LogOptions.
		NodeOptions:      NewDefaultNode(),           // Initializes with default NodeOptions.
		QOSOptions:       NewDefaultQOS(),            // Initializes with default QOSOptions.
		QueryOptions:     NewDefaultQuery(),          // Initializes with default QueryOptions.
		TxOptions:        NewDefaultTx(),             // Initializes with default TxOptions.
	}
}

// WithHandshakeOptions sets the HandshakeOptions for the Options and returns the updated Options.
func (o *Options) WithHandshakeOptions(v *HandshakeOptions) *Options {
	o.HandshakeOptions = v
	return o
}

// WithKeyringOptions sets the KeyringOptions for the Options and returns the updated Options.
func (o *Options) WithKeyringOptions(v *KeyringOptions) *Options {
	o.KeyringOptions = v
	return o
}

// WithLogOptions sets the LogOptions for the Options and returns the updated Options.
func (o *Options) WithLogOptions(v *sdkoptions.LogOptions) *Options {
	o.LogOptions = v
	return o
}

// WithNodeOptions sets the NodeOptions for the Options and returns the updated Options.
func (o *Options) WithNodeOptions(v *NodeOptions) *Options {
	o.NodeOptions = v
	return o
}

// WithQOSOptions sets the QOSOptions for the Options and returns the updated Options.
func (o *Options) WithQOSOptions(v *QOSOptions) *Options {
	o.QOSOptions = v
	return o
}

// WithQueryOptions sets the QueryOptions for the Options and returns the updated Options.
func (o *Options) WithQueryOptions(v *QueryOptions) *Options {
	o.QueryOptions = v
	return o
}

// WithTxOptions sets the TxOptions for the Options and returns the updated Options.
func (o *Options) WithTxOptions(v *TxOptions) *Options {
	o.TxOptions = v
	return o
}

// WithHandshakeOptionsFromCmd updates HandshakeOptions in the Options from the given command's flags.
func (o *Options) WithHandshakeOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewHandshakeOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithHandshakeOptions(opts), nil
}

// WithKeyringOptionsFromCmd updates KeyringOptions in the Options from the given command's flags.
func (o *Options) WithKeyringOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewKeyringOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithKeyringOptions(opts), nil
}

// WithLogOptionsFromCmd updates LogOptions in the Options from the given command's flags.
func (o *Options) WithLogOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := sdkoptions.NewLogOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithLogOptions(opts), nil
}

// WithNodeOptionsFromCmd updates NodeOptions in the Options from the given command's flags.
func (o *Options) WithNodeOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewNodeOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithNodeOptions(opts), nil
}

// WithQOSOptionsFromCmd updates QOSOptions in the Options from the given command's flags.
func (o *Options) WithQOSOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewQOSOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithQOSOptions(opts), nil
}

// WithQueryOptionsFromCmd updates QueryOptions in the Options from the given command's flags.
func (o *Options) WithQueryOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewQueryOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithQueryOptions(opts), nil
}

// WithTxOptionsFromCmd updates TxOptions in the Options from the given command's flags.
func (o *Options) WithTxOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewTxOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithTxOptions(opts), nil
}

// NewOptionsFromCmd creates and returns an Options instance populated with values from the given command's flags.
func NewOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	// Retrieves HandshakeOptions from command flags.
	keyOpts, err := NewHandshakeOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves KeyringOptions from command flags.
	keyringOpts, err := NewKeyringOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves LogOptions from command flags.
	logOpts, err := sdkoptions.NewLogOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves NodeOptions from command flags.
	nodeOpts, err := NewNodeOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves QOSOptions from command flags.
	qosOpts, err := NewQOSOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves QueryOptions from command flags.
	queryOpts, err := NewQueryOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves TxOptions from command flags.
	txOpts, err := NewTxOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new Options instance populated with the retrieved flag values.
	return &Options{
		HandshakeOptions: keyOpts,
		KeyringOptions:   keyringOpts,
		LogOptions:       logOpts,
		NodeOptions:      nodeOpts,
		QOSOptions:       qosOpts,
		QueryOptions:     queryOpts,
		TxOptions:        txOpts,
	}, nil
}
