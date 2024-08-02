package options

import (
	"time"

	"github.com/sentinel-official/sentinel-go-sdk/client/options"
	"github.com/spf13/cobra"
)

// QueryOptions represents options for making queries.
type QueryOptions struct {
	MaxRetries int           `json:"max_retries" toml:"max_retries"` // MaxRetries is the maximum number of retries for the query.
	RPCAddr    string        `json:"rpc_addr" toml:"rpc_addr"`       // RPCAddr is the address of the RPC server.
	Timeout    time.Duration `json:"timeout" toml:"timeout"`         // Timeout is the maximum duration for the query to be executed.
}

// NewDefaultQueryOptions creates a new QueryOptions instance with default values.
func NewDefaultQueryOptions() *QueryOptions {
	return &QueryOptions{
		MaxRetries: options.DefaultQueryMaxRetries,
		RPCAddr:    options.DefaultQueryRPCAddr,
		Timeout:    options.DefaultQueryTimeout,
	}
}

// WithMaxRetries sets the MaxRetries field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithMaxRetries(v int) *QueryOptions {
	q.MaxRetries = v
	return q
}

// WithRPCAddr sets the RPCAddr field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithRPCAddr(v string) *QueryOptions {
	q.RPCAddr = v
	return q
}

// WithTimeout sets the Timeout field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithTimeout(v time.Duration) *QueryOptions {
	q.Timeout = v
	return q
}

// AddQueryFlagsToCmd adds query-related flags to the given cobra command.
func AddQueryFlagsToCmd(cmd *cobra.Command) {
	options.SetFlagQueryMaxRetries(cmd)
	options.SetFlagQueryRPCAddr(cmd)
	options.SetFlagQueryTimeout(cmd)
}

// NewQueryOptionsFromCmd creates and returns QueryOptions from the given cobra command's flags.
func NewQueryOptionsFromCmd(cmd *cobra.Command) (*QueryOptions, error) {
	maxRetries, err := options.GetQueryMaxRetriesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	rpcAddr, err := options.GetQueryRPCAddrFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	timeout, err := options.GetQueryTimeoutFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new QueryOptions instance populated with the retrieved flag values.
	return &QueryOptions{
		MaxRetries: maxRetries,
		RPCAddr:    rpcAddr,
		Timeout:    timeout,
	}, nil
}
