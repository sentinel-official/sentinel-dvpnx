package options

import (
	"github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
	"github.com/spf13/cobra"
)

// AddLogFlagsToCmd attaches logging-related flags to the provided cobra command.
func AddLogFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagLogFormat(cmd)
	flags.SetFlagLogLevel(cmd)
}
