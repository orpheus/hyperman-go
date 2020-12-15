package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os/exec"
)

func main() {
	rootViper := spawnHyperspaceViper(".")

	// For now just grab the default network. Later I'll want to pass
	// this in as a command line argument via commander
	network := rootViper.Get("defaultNetwork")
	fmt.Println("default network found: ", network)

	networkPath := fmt.Sprintf("networks/%s", network)

	// look through network folders for the default network
	networkViper := spawnHyperspaceViper(networkPath)
	fmt.Println("config path for default network: ", fmt.Sprintf("networks/%s", network))

	// check to make sure nods config is set
	if !networkViper.IsSet("nodes") {
		log.Panicf("node configurations not set for network: %s", network)
	}

	// make sure an orderer is defined
	if !networkViper.IsSet("nodes.orderer") {
		log.Panicf("Need to specify at least one orderer node in a hyperspace network configuration: %s", network)
	}

	// get names of the orderers
	ordererNodes := networkViper.GetStringSlice("nodes.orderer")

	// create a hyperspace vipers map
	hyperspaceVipers := make(map[string]map[string]*viper.Viper)
	// create a keys for nodes"
	hyperspaceVipers["orderers"] = make(map[string]*viper.Viper)
	hyperspaceVipers["peers"] = make(map[string]*viper.Viper)

	// loop through each orderer config and spawn a hyperspace viper
	for _, ordererName := range ordererNodes {
		ordererPath := fmt.Sprintf("%s/nodes/orderers/%s", networkPath, ordererName)
		hyperspaceVipers["orderers"][ordererName] = spawnHyperspaceViper(ordererPath)
	}

	// grab peer configs if set
	if networkViper.IsSet("nodes.peers") {
		for _, peerName := range ordererNodes {
			peerPath := fmt.Sprintf("%s/nodes/peers/%s", networkPath, peerName)
			hyperspaceVipers["orderers"][peerName] = spawnHyperspaceViper(peerPath)
		}
	}

	fmt.Println(hyperspaceVipers)
	// spawn orderers
	ordererNodeConfigs := hyperspaceVipers["orderers"]
	for _, hyperviper := range ordererNodeConfigs { // go routine to spawn nodes?
		// form the cmd line argument for the spawnNode shell script
		args := make([]string, 0) // better way to do this?

		// grab binary name
		binary := hyperviper.GetString("binary")
		args = append(args,"-b")
		args = append(args, binary)

		// grab env vars
		environment := hyperviper.GetStringSlice("environment")
		for _, env := range environment {
			args = append(args, "-e")
			args = append(args, env)
		}

		out, err := exec.Command("./cmdscripts/spawn-node.sh", args...).Output()
		log.Printf("ErrorCode = %s\nOutput = %s\n",  err, out)
	}

	// spawn peers
	peerNodeConfigs := hyperspaceVipers["peers"]
	for _, hyperviper := range peerNodeConfigs { // go routine to spawn nodes?
		// form the cmd line argument for the spawnNode shell script
		args := make([]string, 0) // better way to do this?

		// set binary
		binary := hyperviper.GetString("binary")
		args = append(args,"-b")
		args = append(args, binary)

		// set cmd
		startCmd := hyperviper.GetString("node start")
		args = append(args,"-cmd")
		args = append(args, startCmd)

		// set env vars
		environment := hyperviper.GetStringSlice("environment")
		for _, env := range environment {
			args = append(args, "-e")
			args = append(args, env)
		}

		out, err := exec.Command("./cmdscripts/spawn-node.sh", args...).Output()
		log.Printf("ErrorCode = %s\nOutput = %s\n",  err, out)
	}

}

func spawnHyperspaceViper (paths ...string) *viper.Viper {
	v := viper.New()
	v.SetConfigName("hyperspace")
	v.SetConfigType("yaml")
	for _, path := range paths {
		v.AddConfigPath(path)
		// adding this so we can run from the cmdscripts dir
		v.AddConfigPath(fmt.Sprintf("../%s", path))
	}
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}

