package network

import (
	"fmt"
	"github.com/orpheus/hyperspace/util"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

func createCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a network.",
		Long:  `Create a new network file system`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("expecting 1 argument for network name")
			}
			fmt.Println("Creating network...")
			return makeNetwork(args[0])
		},
	}
}

func makeNetwork (networkName string) error {
	commandCenter := os.Getenv("HYPERSPACE_PATH")
	if len(commandCenter) == 0 {
		commandCenter = "."
	}
	networkRoot := os.Getenv("HYPERSPACE_NETWORK_ROOT")
	if len(networkRoot) == 0 {
		// relative path
		networkRoot = "networks"
	}
	networkRoot = filepath.Join(commandCenter, networkRoot)

	fmt.Println(networkRoot)

	newNetworkPath := filepath.Join(networkRoot, networkName)
	exists, err := util.FileOrDirectoryExists(newNetworkPath)
	if err != nil {
		log.Fatalf("Failed to check if network path exists: %v", err)
	}
	if exists {
		fmt.Println("Network already exists, overwriting...")
		//log.Fatalf("Network path already exists: %s", newNetworkPath)
	}
	err = os.MkdirAll(newNetworkPath, 0755)
	if err != nil {
		log.Fatalf("Failed to make new network directory: %v", err)
	}

	// copy exampleNetwork into newly created directory

	return nil
}

