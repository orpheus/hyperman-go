package main

import (
	"fmt"
	"github.com/orpheus/hyperspace/util"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"strings"
)

const NETWORK_PATH_REPLACER = "${NETWORK_PATH}"

func main() {
	rootViper := util.SpawnHyperspaceViper(".")

	// For now just grab the default network. Later I'll want to pass
	// this in as a command line argument via commander
	network := rootViper.Get("defaultNetwork")

	networkPath := fmt.Sprintf("../networks/%s", network)

	// look through network folders for the default network
	networkViper := util.SpawnHyperspaceViper(networkPath)

	// check to make sure nods config is set
	if !networkViper.IsSet("scriptPath") {
		log.Panicf("node configurations not set for network: %s", network)
	}

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
		hyperspaceVipers["orderers"][ordererName] = util.SpawnHyperspaceViper(ordererPath)
	}

	// grab peer configs if set
	if networkViper.IsSet("nodes.peers") {
		for _, peerName := range ordererNodes {
			peerPath := fmt.Sprintf("%s/nodes/peers/%s", networkPath, peerName)
			hyperspaceVipers["orderers"][peerName] = util.SpawnHyperspaceViper(peerPath)
		}
	}

	// combine network path with relative script path
	scriptPath := networkViper.GetString("scriptPath")
	scriptPath = fmt.Sprintf("%s/%s", networkPath, scriptPath)

	// spawn orderers
	ordererNodeConfigs := hyperspaceVipers["orderers"]
	for _, hyperviper := range ordererNodeConfigs { // go routine to spawn
		// nodes?
		// form the cmd line argument for the spawnNode shell script
		args := make([]string, 0) // better way to do this?

		// grab binary name
		binary := hyperviper.GetString("binary")
		args = append(args,"-b", binary)

		// grab env vars
		environment := hyperviper.GetStringSlice("environment")
		for _, env := range environment {
			// need to replace the environment relative paths
			// with a path that the command center can recognize
			env = strings.Replace(env, NETWORK_PATH_REPLACER, networkPath, 1 )
			args = append(args, "-e", string(env))
		}

		// here we set the FABRIC_CFG_PATH to the orderer's directory
		// for the orderer binary to find the orderer.yaml config
		//args = append(args, "-e")
		//pathToOrdererConfig := fmt.Sprintf("%s/nodes/orderers/%s",
		//	networkPath, name)
		//args = append(args, fmt.Sprintf("FABRIC_CFG_PATH=%s", pathToOrdererConfig))

		//fmt.Println(fmt.Sprintf("FABRIC_CFG_PATH=%s", pathToOrdererConfig))
		args = append(args, "-cmd", "start")
		fmt.Println("COMMAND GO", scriptPath)
		cmd := exec.Command(
			scriptPath,
			args...
			)
		// NEEDED TO SEE LOGS IN TERMINAL
		// in the bash script make sure to route
		// stdErr to stdOut to see errors as well ( 2>&1 )
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		// This runs the command waits for it to finish,
		// meanwhile the stdout output is routed to stdout
		// so I can see what's going on in terminal
		// when I start to spawn multiple of them
		// maybe switch to the "Start" cmd, run them in
		// go routines, grab their pids for deactivation later

		// for now this is fine
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Just ran subprocess %d, hanging...\n", cmd.Process.Pid)
	}

	// spawn peers
	//peerNodeConfigs := hyperspaceVipers["peers"]
	//for _, hyperviper := range peerNodeConfigs { // go routine to spawn nodes?
	//	// form the cmd line argument for the spawnNode shell script
	//	args := make([]string, 0) // better way to do this?
	//
	//	// set binary
	//	binary := hyperviper.GetString("binary")
	//	args = append(args,"-b")
	//	args = append(args, binary)
	//
	//	// set cmd
	//	startCmd := hyperviper.GetString("node start")
	//	args = append(args,"-cmd")
	//	args = append(args, startCmd)
	//
	//	// set env vars
	//	environment := hyperviper.GetStringSlice("environment")
	//	for _, env := range environment {
	//		args = append(args, "-e")
	//		args = append(args, env)
	//	}
	//
	//	out, err := exec.Command("./cmdscripts/spawn-node.sh", args...).Output()
	//	log.Printf("ErrorCode = %s\nOutput = %s\n",  err, out)
	//}

}

