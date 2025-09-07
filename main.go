package main

import (
	"context"
	"os"

	"github.com/sentinel-official/sentinel-go-sdk/app"

	"github.com/sentinel-official/sentinel-dvpnx/cmd"
)

func main() {
	exitCode := app.Run(context.Background(), cmd.NewRootCmd)
	os.Exit(exitCode)
}
