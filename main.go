package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-dvpnx/cmd"
)

func main() {
	// Create a context that listens for SIGINT and SIGTERM signals
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Retrieve the user's home directory.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: Unable to determine user home directory: %v\n", err)
		os.Exit(1)
	}

	// Enable Cobra's feature to traverse and execute hooks for commands.
	cobra.EnableTraverseRunHooks = true

	// Initialize the root command for the application.
	rootCmd := cmd.NewRootCmd(homeDir)

	// Execute the root command.
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
