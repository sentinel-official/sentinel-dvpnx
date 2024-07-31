package main

import (
	"github.com/cosmos/cosmos-sdk/version"
	sdkcmd "github.com/sentinel-official/sentinel-go-sdk/cmd"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:          "sentinelnode",
		SilenceUsage: true,
	}

	root.AddCommand(
		sdkcmd.KeysCmd(),
		version.NewVersionCommand(),
	)

	_ = root.Execute()
}
