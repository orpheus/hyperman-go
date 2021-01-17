package main

import (
	"github.com/orpheus/hyperspace/core/network"
	"github.com/spf13/cobra"
	"os"
)

// The main command describes the service and
// defaults to printing the help message.
var mainCmd = &cobra.Command{Use: "hyperspace"}


func main () {
	mainCmd.AddCommand(network.Cmd())
	// On failure Cobra prints the usage message and error string, so we only
	// need to exit with a non-0 status
	if mainCmd.Execute() != nil {
		os.Exit(1)
	}
}