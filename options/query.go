package options

import (
	"time"

	sdkflags "github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/dvpn-node/cmd/flags"
)

// Query represents options for making queries.
type Query struct {
	MaxRetries int           `json:"max_retries" toml:"max_retries"` // MaxRetries is the maximum number of retries for the query.
	RPCAddrs   []string      `json:"rpc_addrs" toml:"rpc_addrs"`     // RPCAddrs is a slice of addresses of the RPC servers.
	Timeout    time.Duration `json:"timeout" toml:"timeout"`         // Timeout is the maximum duration for the query to be executed.
}

// NewQuery creates a new Query instance with default values.
func NewQuery() *Query {
	return &Query{
		MaxRetries: sdkflags.DefaultQueryMaxRetries,
		RPCAddrs:   flags.DefaultQueryRPCAddrs,
		Timeout:    sdkflags.DefaultQueryTimeout,
	}
}

// WithMaxRetries sets the MaxRetries field and returns the modified Query instance.
func (q *Query) WithMaxRetries(v int) *Query {
	q.MaxRetries = v
	return q
}

// WithRPCAddrs sets the RPCAddrs field and returns the modified Query instance.
func (q *Query) WithRPCAddrs(v ...string) *Query {
	q.RPCAddrs = v
	return q
}

// WithTimeout sets the Timeout field and returns the modified Query instance.
func (q *Query) WithTimeout(v time.Duration) *Query {
	q.Timeout = v
	return q
}

// AddQueryFlagsToCmd adds query-related flags to the given cobra command.
func AddQueryFlagsToCmd(cmd *cobra.Command) {
	sdkflags.SetFlagQueryMaxRetries(cmd)
	flags.SetFlagQueryRPCAddrs(cmd)
	sdkflags.SetFlagQueryTimeout(cmd)
}

// NewQueryFromCmd creates and returns Query from the given cobra command's sdkflags.
func NewQueryFromCmd(cmd *cobra.Command) (*Query, error) {
	// Retrieve the max retries flag value from the command.
	maxRetries, err := sdkflags.GetQueryMaxRetriesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the RPC addresses flag value from the command.
	rpcAddrs, err := flags.GetQueryRPCAddrsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the timeout flag value from the command.
	timeout, err := sdkflags.GetQueryTimeoutFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new Query instance populated with the retrieved flag values.
	return &Query{
		MaxRetries: maxRetries,
		RPCAddrs:   rpcAddrs,
		Timeout:    timeout,
	}, nil
}
