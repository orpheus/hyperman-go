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
	// get path to project
	commandCenter := os.Getenv("HYPERSPACE_PATH")
	if len(commandCenter) == 0 {
		commandCenter = "."
	}
	// get desired network root location. the network root
	// is the path to the`networks` directory that contains
	// all networks. in the current version, 0.1.0, it is at
	// the top level of the project, so
	// networkRoot = "$HYPERLEDGER_ROOT/networks"
	networkRoot := os.Getenv("HYPERSPACE_NETROOT")
	if len(networkRoot) == 0 {
		// relative path
		networkRoot = "networks"
	}
	networkRoot = filepath.Join(commandCenter, networkRoot)

	fmt.Printf("Network root: %s", networkRoot)

	// make the path for the new network location
	newNetworkPath := filepath.Join(networkRoot, networkName)
	// check if the network already exists
	exists, err := util.FileOrDirectoryExists(newNetworkPath)
	if err != nil {
		log.Fatalf("Failed to check if network path exists: %v", err)
	}
	if exists {
		fmt.Println("Network already exists, overwriting...")
		//log.Fatalf("Network path already exists: %s", newNetworkPath)
	}
	// create directory for new network
	err = os.MkdirAll(newNetworkPath, 0755)
	if err != nil {
		log.Fatalf("Failed to make new network directory: %v", err)
	}

	// create directories for [config, configtxgen, cryptogen, nodes, organizations
	// USE NETWORK_HYPERSPACE.YAML TO CREATE THE NETWORK
	// 1. Create a hyperspace network config (later, interactive console)
	// 2. Call `hyperspace network create "newNetwork" path/to/network-hyperspace.yaml`
	// 3. Code, 1. Reads config 2. Creates node system for each node listed

	return nil
}

