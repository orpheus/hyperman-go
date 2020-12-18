package main

import (
	"context"
	"fmt"
	"github.com/orpheus/hyperspace/core"
	"github.com/orpheus/hyperspace/core/configtxgen"
	"github.com/orpheus/hyperspace/core/cryptogen"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)


	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		osignal := <-c
		log.Printf("Received shutdown signal: %v", osignal)
		cancel()
	}()

	// Create a viper at the root of the application
	// to read in the active network and any other
	// configuration that may be applicable to their
	// entire star system.
	rv := core.CreateRootViper()
	// Read from the hyperspace config and make the crypto
	// by executing a commander (cmd/shell/script)
	cryptogen.Initialize(rv).Make()
	// Read from the hyperspace config for configtxgen
	// create the genesis block and consortiums.
	configtxgen.Initialize(rv).Create()


	// check to make sure nods config is set
	if !rv.NetworkViper.IsSet("scriptPath") {
		log.Panicf("node configurations not set for network: %s", rv.Network)
	}

	/**
	The following code checks the hyperspace.yaml at the root of a network.
	From it, it will grab the orderers and peers it will need to spawn by their listed name.
	It knows to look for `nodes.orderers` and `nodes.`peers` and then creates paths to the nodes
	using their name like so: `{network}/nodes/orderers/{orderer_name}`. Replace "orderers" with "peers" and vice versa
	*/
	// check to make sure nods config is set
	if !rv.NetworkViper.IsSet("nodes") {
		log.Panicf("node configurations not set for network: %s", rv.Network)
	}
	// make sure an orderer is defined
	if !rv.NetworkViper.IsSet("nodes.orderers") {
		log.Panicf("Need to specify at least one orderer node in a hyperspace network configuration: %s", rv.Network)
	}

	// combine network path with relative script path
	scriptPath := rv.NetworkViper.GetString("scriptPath")
	scriptPath = filepath.Join(rv.NetworkPath, scriptPath)

	//processes := make(chan *os.Process)

	fmt.Println("Spawning orderers")
	spawnOrderers(rv, scriptPath, ctx) // goroutine?
	fmt.Println("Spawning peers")
	spawnPeers(rv, scriptPath, ctx) // goroutine?

	<-ctx.Done()

	// stall os.exit so the KillProcess logs show
	// there has to be a better way to do this
	// I could send in a channel to each spawn function
	// and wait for that channel to fill up maybe?
	log.Println("Hyperspace closing...")
	time.Sleep(time.Second * 3)

	// channel wait example
	// wait for modules to be initialized
	// process ids needs to be the length of the total amount of processes
	//for range processIds {
	//	pid := <-processIds
	//	log.Info("Link-Module Initialized", logger.Attrs{"ID": id})
	//}
}

/**
SPAWN ORDERERS:
This code loops through the hyperspace configurations gathered above for the orderers,
and using the specified env variables, binary name, and startCmd, will call the
spawn-node.sh script to spawn an orderer.
Note: no "scriptPath" is defined in this configuration. This could be added later
to maximize flexibility and allow others to hack it.
Note: If a "scriptPath" were allowed here, the path would be relative to where the
HYPERSPACE_CONTROLLER (aka GOD) is, meaning that GOD would have to join the relative
path with the absolute path of the CONTROL_CENTER (the hyperspace directory root)
Note: need to name the HYPERSPACE_CONTROLLER, not GOD. Who or what controls the Hyperspace?
...think more on this later
*/
func spawnOrderers (rv *core.RootViper, scriptPath string, ctx context.Context) {
	// create a key:map for orderers
	ordererVipers := make(map[string]*viper.Viper)

	// get the orderers
	ordererNodes := rv.NetworkViper.GetStringSlice("nodes.orderers")
	// loop through each orderer config and spawn a hyperspace viper
	for _, ordererName := range ordererNodes {
		ordererPath := filepath.Join(rv.NetworkPath, "/nodes/orderers/", ordererName)
		// toDo: create an orderer/peer struct with a HyperViper
		ordererVipers[ordererName] = core.SpawnHyperSpaceViper(ordererPath)
	}

	for ordererName, hyperviper := range ordererVipers { // go routine?
		// form the cmd line argument for the spawnNode shell script
		args := make([]string, 0) // better way to do this?

		// grab binary name
		binary := hyperviper.GetString("binary")
		// set the command_center for the cmdscript to the node's directory
		commandCenter := filepath.Dir(hyperviper.ConfigFileUsed())
		args = append(args,
			"-b", binary,
			"-cmd", "start",
			// toDo: THIS MAY NO LONGER BE NEEDED
			// instead you COULD use cmd.Dir and run the go exec cmd in the command-center directly
			// vs sending the bash scripts there manually (may be safer?)...
			"--command-center", commandCenter,
		)

		// grab and set env vars
		environment := hyperviper.GetStringSlice("environment")
		for _, env := range environment {
			args = append(args, "-e", string(env))
		}

		// just want to make sure I'm closing over these variables just in case
		// because I don't know golang that well yet
		go func (orderName, scriptPath string, args []string, ctx context.Context) {
			cmd := exec.Command(
				scriptPath,
				args...
			)
			// NEEDED TO SEE LOGS IN TERMINAL
			// in the bash script make sure to route
			// stdErr to stdOut to see errors as well ( 2>&1 )
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			// the stdout output is routed to stdout
			// so I can see what's going on in terminal
			// when I start to spawn multiple of them
			// maybe switch to the "Start" cmd, run them in
			// go routines, grab their pids for deactivation later
			// for now this is fine, when you change it, change peer as well
			err := cmd.Start()
			//if err != nil {
			//	log.Fatal(err)
			//}
			//log.Printf("Just ran orderer subprocess %d\n", cmd.Process.Pid)

			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Ran peer subprocess %s:%d...\n", ordererName, cmd.Process.Pid)

			// Shutdown processes on os signals
			<-ctx.Done()
			KillProcess(cmd, "orderer")
		}(ordererName, scriptPath, args, ctx)
	}
}

func KillProcess (cmd *exec.Cmd, name string) {
	if err := cmd.Process.Kill(); err != nil {
		log.Fatalf("Failed to kill process: %s\n%v", name, err)
	}
	log.Printf("Succesfully killed process: %s", name)
}

/**
SPAWN PEERS:
This code loops through the hyperspace configurations gathered above for the peers,
and using the specified env variables, binary name, and startCmd, will call the
spawn-node.sh script to spawn an orderer.
Note: no "scriptPath" is defined in this configuration. This could be added later
to maximize flexibility and allow others to hack it.
Note: If a "scriptPath" were allowed here, the path would be relative to where the
HYPERSPACE_CONTROLLER is, meaning that would have to join the relative
path with the absolute path of the CONTROL_CENTER (the hyperspace directory root)
Note: need to name the HYPERSPACE_CONTROLLER, not GOD. Who or what controls the Hyperspace?
...think more on this later
*/
func spawnPeers (rv *core.RootViper, scriptPath string, ctx context.Context) {
	// create a key:map for peers
	peerVipers := make(map[string]*viper.Viper)
	// get the peers
	peerNodes := rv.NetworkViper.GetStringSlice("nodes.peers")
	// loop through each peer config and spawn a hyperspace viper
	if rv.NetworkViper.IsSet("nodes.peers") {
		for _, peerName := range peerNodes {
			peerPath := filepath.Join(rv.NetworkPath, "/nodes/peers/", peerName)
			peerVipers[peerName] = core.SpawnHyperSpaceViper(peerPath)
		}
	}

	for peerName, hyperviper := range peerVipers { // go routine to spawn nodes?
		// form the cmd line argument for the spawnNode shell script
		args := make([]string, 0) // better way to do this?

		// set binary and startCmd
		binary := hyperviper.GetString("binary")

		startCmd := hyperviper.GetString("startCmd")
		// set the command_center for the cmdscript to the node's directory
		commandCenter := filepath.Dir(hyperviper.ConfigFileUsed())
		args = append(
			args,
			"-b", binary,
			"-cmd", startCmd,
			"--command-center", commandCenter,
		)

		// set env vars
		environment := hyperviper.GetStringSlice("environment")
		for _, env := range environment {
			args = append(args, "-e")
			args = append(args, env)
		}

		go func (peerName, scriptPath string, args []string, ctx context.Context) {
			cmd := exec.Command(
				scriptPath,
				args...
			)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Ran peer subprocess %s:%d...\n", peerName, cmd.Process.Pid)

			// Shutdown processes on os signals
			<-ctx.Done()
			KillProcess(cmd, peerName)
		}(peerName, scriptPath, args, ctx)
	}
}

