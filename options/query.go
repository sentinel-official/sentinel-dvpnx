package options

import (
	"errors"
	"time"

	sdkflags "github.com/sentinel-official/sentinel-go-sdk/flags"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/dvpn-node/flags"
)

// Query represents options for making queries.
type Query struct {
	MaxRetries int      `json:"max_retries" toml:"max_retries"` // MaxRetries is the maximum number of retries for the query.
	RPCAddrs   []string `json:"rpc_addrs" toml:"rpc_addrs"`     // RPCAddrs is a slice of addresses of the RPC servers.
	Timeout    string   `json:"timeout" toml:"timeout"`         // Timeout is the maximum duration for the query to be executed.
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
	q.Timeout = v.String()
	return q
}

// GetMaxRetries returns the maximum number of retries.
func (q *Query) GetMaxRetries() int {
	return q.MaxRetries
}

// GetRPCAddrs returns the RPC addresses.
func (q *Query) GetRPCAddrs() []string {
	return q.RPCAddrs
}

// GetTimeout returns the timeout duration.
func (q *Query) GetTimeout() time.Duration {
	timeout, err := time.ParseDuration(q.Timeout)
	if err != nil {
		panic(err)
	}

	return timeout
}

// ValidateQueryMaxRetries checks if the MaxRetries field is valid.
func ValidateQueryMaxRetries(maxRetries int) error {
	if maxRetries < 0 {
		return errors.New("max retries must be non-negative")
	}

	return nil
}

// ValidateQueryRPCAddrs checks if the RPCAddrs field is valid.
func ValidateQueryRPCAddrs(rpcAddrs []string) error {
	if len(rpcAddrs) == 0 {
		return errors.New("at least one RPC address must be provided")
	}

	return nil
}

// ValidateQueryTimeout checks if the Timeout field is valid.
func ValidateQueryTimeout(timeout string) error {
	if _, err := time.ParseDuration(timeout); err != nil {
		return errors.New("invalid timeout duration")
	}

	return nil
}

// Validate validates all fields of the Query struct.
func (q *Query) Validate() error {
	if err := ValidateQueryMaxRetries(q.MaxRetries); err != nil {
		return err
	}
	if err := ValidateQueryRPCAddrs(q.RPCAddrs); err != nil {
		return err
	}
	if err := ValidateQueryTimeout(q.Timeout); err != nil {
		return err
	}

	return nil
}

// NewQueryFromCmd creates and returns Query from the given cobra command's sdkflags.
func NewQueryFromCmd(cmd *cobra.Command) (*Query, error) {
	// Retrieve the max retries flag value from the command.
	maxRetries, err := sdkflags.GetQueryMaxRetries(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the RPC addresses flag value from the command.
	rpcAddrs, err := flags.GetQueryRPCAddrs(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the timeout flag value from the command.
	timeout, err := sdkflags.GetQueryTimeout(cmd)
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
