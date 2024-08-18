package flags

import (
	sdkflags "github.com/sentinel-official/sentinel-go-sdk/flags"
	"github.com/spf13/cobra"
)

// AddKeyringFlags adds keyring-related flags to the given cobra command.
func AddKeyringFlags(cmd *cobra.Command) {
	sdkflags.SetFlagKeyringBackend(cmd)
}
