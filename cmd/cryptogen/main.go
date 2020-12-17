package main

import (
	"fmt"
	"github.com/orpheus/hyperspace/util"
	"log"
	"os"
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
	cryptogenPath := fmt.Sprintf("%s/cryptogen", networkPath)

	cryptogenViper := util.SpawnHyperspaceViper(cryptogenPath)
	fmt.Println("Config path for default network: ", fmt.Sprintf("networks/%s", network))


	scriptPath := cryptogenViper.GetString("scriptPath")
	scriptPath = fmt.Sprintf("%s/%s", cryptogenPath, scriptPath)
	configs := cryptogenViper.GetStringMap("configs")

	for org := range configs {
		configPath := fmt.Sprintf("configs.%s.path", org)
		outputPath := fmt.Sprintf("configs.%s.output", org)

		configPath = fmt.Sprintf("%s/%s", cryptogenPath, cryptogenViper.GetString(configPath))
		outputPath = fmt.Sprintf("%s/%s", cryptogenPath, cryptogenViper.GetString(outputPath))

		cmd := exec.Command("/bin/bash",
			scriptPath,
			"-n", network,
			"-b", cryptogenViper.GetString("fabricBinaryName"),
			"-c", configPath,
			"-o", outputPath,
			"-i", org,
		)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		log.Printf("Cryptogen main script finished with error: %v", err)
	}
}
