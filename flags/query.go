package flags

import (
	sdkflags "github.com/sentinel-official/sentinel-go-sdk/flags"
	"github.com/spf13/cobra"
)

var (
	DefaultQueryRPCAddrs = []string{sdkflags.DefaultQueryRPCAddr}
)

// GetQueryRPCAddrs retrieves the "query.rpc-addrs" flag value from the command.
func GetQueryRPCAddrs(cmd *cobra.Command) ([]string, error) {
	return cmd.Flags().GetStringSlice("query.rpc-addrs")
}

// SetFlagQueryRPCAddrs adds the "query.rpc-addrs" flag to the command.
func SetFlagQueryRPCAddrs(cmd *cobra.Command) {
	cmd.Flags().StringSlice("query.rpc-addrs", DefaultQueryRPCAddrs, "Comma-separated list of addresses of the RPC servers.")
}

// AddQueryFlags adds query-related flags to the given cobra command.
func AddQueryFlags(cmd *cobra.Command) {
	sdkflags.SetFlagQueryMaxRetries(cmd)
	SetFlagQueryRPCAddrs(cmd)
	sdkflags.SetFlagQueryTimeout(cmd)
}
