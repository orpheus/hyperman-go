package main

import (
	"fmt"
	"github.com/orpheus/hyperspace/util"
	"log"
	"os/exec"
)

func main() {
	rootViper := util.SpawnHyperspaceViper(".")

	// For now just grab the default network. Later I'll want to pass
	// this in as a command line argument via commander
	network := rootViper.GetString("defaultNetwork")

	// this needs to become some kind of base path
	// for now assume we're running out of the cmdcenter (cmdscripts dir)
	networkPath := fmt.Sprintf("../networks/%s", network)
	configtxgenPath := fmt.Sprintf("%s/configtxgen", networkPath)

	configtxgenViper := util.SpawnHyperspaceViper(configtxgenPath)
	fmt.Println("Config path for default network: ", fmt.Sprintf("networks/%s", network))

	scriptPath := configtxgenViper.GetString("scriptPath")
	scriptPath = fmt.Sprintf("%s/%s", configtxgenPath, scriptPath)

	configtxPath := configtxgenViper.GetString("configPath")
	configtxPath = fmt.Sprintf("%s/%s", configtxgenPath, configtxPath)

	profile := configtxgenViper.GetString("profile")
	channelID := configtxgenViper.GetString("channelID")

	output := configtxgenViper.GetString("output")
	output = fmt.Sprintf("%s/%s", configtxgenPath, output)

	binary := configtxgenViper.GetString("fabricBinaryName")

	command := exec.Command("/bin/bash",
		scriptPath,
		"-n", network,
		"-b", binary,
		"-c", configtxPath,
		"-p", profile,
		"-ch", channelID,
		"-o", output,
	)

	out, err := command.Output()
	log.Printf("Executed command [%s] %s\nErrorCode = %s\nOutput = %s\n", command.Dir, command.Args, err, out)
}
