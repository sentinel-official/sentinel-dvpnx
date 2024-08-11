package options

import (
	"github.com/sentinel-official/sentinel-go-sdk/client/options"
	"github.com/spf13/cobra"
)

// Options aggregates all the individual option structs.
type Options struct {
	*Handshake   `json:"handshake" toml:"handshake"` // Options related to key creation.
	*Keyring     `json:"keyring" toml:"keyring"`     // Options related to keyring configuration.
	*options.Log `json:"log" toml:"log"`             // Options related to logging.
	*Node        `json:"node" toml:"node"`           // Options related to node.
	*QOS         `json:"qos" toml:"qos"`             // Options related to quality of service (QOS).
	*Query       `json:"query" toml:"query"`         // Options related to querying.
	*Tx          `json:"tx" toml:"tx"`               // Options related to transactions.
}

// New creates and returns a new instance of Options with default values.
func New() *Options {
	return &Options{
		Handshake: NewHandshake(),   // Initializes with default Handshake.
		Keyring:   NewKeyring(),     // Initializes with default Keyring.
		Log:       options.NewLog(), // Initializes with default LogOptions.
		Node:      NewNode(),        // Initializes with default Node.
		QOS:       NewQOS(),         // Initializes with default QOS.
		Query:     NewQuery(),       // Initializes with default Query.
		Tx:        NewTx(),          // Initializes with default Tx.
	}
}

// WithHandshake sets the Handshake for the Options and returns the updated Options.
func (o *Options) WithHandshake(v *Handshake) *Options {
	o.Handshake = v
	return o
}

// WithKeyring sets the Keyring for the Options and returns the updated Options.
func (o *Options) WithKeyring(v *Keyring) *Options {
	o.Keyring = v
	return o
}

// WithLog sets the Log for the Options and returns the updated Options.
func (o *Options) WithLog(v *options.Log) *Options {
	o.Log = v
	return o
}

// WithNode sets the Node for the Options and returns the updated Options.
func (o *Options) WithNode(v *Node) *Options {
	o.Node = v
	return o
}

// WithQOSOptions sets the QOS for the Options and returns the updated Options.
func (o *Options) WithQOSOptions(v *QOS) *Options {
	o.QOS = v
	return o
}

// WithQueryOptions sets the Query for the Options and returns the updated Options.
func (o *Options) WithQueryOptions(v *Query) *Options {
	o.Query = v
	return o
}

// WithTxOptions sets the Tx for the Options and returns the updated Options.
func (o *Options) WithTxOptions(v *Tx) *Options {
	o.Tx = v
	return o
}

// WithHandshakeFromCmd updates Handshake in the Options from the given command's flags.
func (o *Options) WithHandshakeFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewHandshakeFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithHandshake(opts), nil
}

// WithKeyringFromCmd updates Keyring in the Options from the given command's flags.
func (o *Options) WithKeyringFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewKeyringFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithKeyring(opts), nil
}

// WithLogFromCmd updates LogOptions in the Options from the given command's flags.
func (o *Options) WithLogFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := options.NewLogFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithLog(opts), nil
}

// WithNodeFromCmd updates Node in the Options from the given command's flags.
func (o *Options) WithNodeFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewNodeFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithNode(opts), nil
}

// WithQOSFromCmd updates QOS in the Options from the given command's flags.
func (o *Options) WithQOSFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewQOSFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithQOSOptions(opts), nil
}

// WithQueryFromCmd updates Query in the Options from the given command's flags.
func (o *Options) WithQueryFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewQueryFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithQueryOptions(opts), nil
}

// WithTxFromCmd updates Tx in the Options from the given command's flags.
func (o *Options) WithTxFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewTxFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithTxOptions(opts), nil
}

// NewFromCmd creates and returns an Options instance populated with values from the given command's flags.
func NewFromCmd(cmd *cobra.Command) (*Options, error) {
	// Retrieves Handshake from command flags.
	key, err := NewHandshakeFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Keyring from command flags.
	keyring, err := NewKeyringFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves LogOptions from command flags.
	log, err := options.NewLogFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Node from command flags.
	node, err := NewNodeFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves QOS from command flags.
	qos, err := NewQOSFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Query from command flags.
	query, err := NewQueryFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Tx from command flags.
	tx, err := NewTxFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new Options instance populated with the retrieved flag values.
	return &Options{
		Handshake: key,
		Keyring:   keyring,
		Log:       log,
		Node:      node,
		QOS:       qos,
		Query:     query,
		Tx:        tx,
	}, nil
}
