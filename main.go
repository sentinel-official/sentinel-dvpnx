package main

import (
	"github.com/sentinel-official/sentinel-go-sdk/app"

	"github.com/sentinel-official/sentinel-dvpnx/cmd"
)

func main() {
	app.Run(cmd.NewRootCmd)
}
