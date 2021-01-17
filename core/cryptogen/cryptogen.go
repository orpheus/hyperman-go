package cryptogen

import (
	"fmt"
	"github.com/orpheus/hyperspace/core/util"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

/**
HyperCryptogen
hyperspace + fabric cryptogen
*/
type Cryptogen struct {
	// network name
	network string
	// hyperspace viper
	hv *util.HyperViper
	// Cryptogen fabric binary name. NOT A HYPERSPACE BINARY.
	// Changed only if you generated a custom binary name during build output
	fabricBinary string
	// Path to Hyperspace cmdscript
	scriptPath string
}

func (c *Cryptogen) init(rv *util.RootViper) {
	c.network = rv.Network
	// this will look in the configtxgen directory in the active network
	// -- is this comment correct?
	c.hv = util.CreateHyperViper(filepath.Join(rv.NetworkPath, "cryptogen"))

	c.fabricBinary = c.hv.Viper.GetString("fabricBinary")

	scriptPath := c.hv.Viper.GetString("scriptPath")

	// this will need to change if the script path isn't relative to the
	// network, but relative to the HYPERSPACE_ROOT, in which case you
	// would join the scriptName with the HYPERSPACE_ROOT_PATH
	c.scriptPath = filepath.Join(c.hv.Path, scriptPath)
}

func (c *Cryptogen) Make() {
	for org := range c.hv.Viper.GetStringMap("configs") {
		configPath := fmt.Sprintf("configs.%s.path", org)
		outputPath := fmt.Sprintf("configs.%s.output", org)

		configPath = filepath.Join(c.hv.Path, c.hv.Viper.GetString(configPath))
		outputPath = filepath.Join(c.hv.Path, c.hv.Viper.GetString(outputPath))

		cmd := exec.Command("/bin/bash",
			c.scriptPath,
			"-n", c.network,
			"-b", c.fabricBinary,
			"-c", configPath,
			"-o", outputPath,
			"-i", org,
		)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			log.Panicf("Error making cryptogen.\n failed on %s", org)
		}
		log.Printf("Cryptogen main script finished with error: %v", err)
	}
}

/**
Initialize a HyperCryptogen with the RootViper
*/
func Initialize(rv *util.RootViper) *Cryptogen {
	cryp := &Cryptogen{}
	cryp.init(rv)
	return cryp
}
