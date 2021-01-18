package configtxgen

import (
	"fmt"
	"github.com/orpheus/hyperspace/util"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

/**
HyperConfigtxgen
hyperspace + fabric configtxgen
 */
type Configtxgen struct {
	// network name
	network string
	// hyperspace viper
	hv *util.HyperViper
	// Cryptogen fabric binary name. NOT A HYPERSPACE BINARY.
	// Changed only if you generated a custom binary name during build output
	fabricBinary string
	// Path to Hyperspace cmdscript
	scriptPath string
	// Directory path that contains fabric configtx.yaml configuration
	configPath string
	// fabric configtxgen config
	profile string
	// fabric configtxgen config
	channelID string
	// fabric configtxgen config
	output string
}

/*
Read from hyperspace configuration and set properties
*/
func (c *Configtxgen) init (rv *util.RootViper) {
	c.network = rv.Network

	// this will look in the configtxgen directory in the active network
	hv := util.CreateHyperViper(filepath.Join(rv.NetworkPath, "configtxgen"))

	c.fabricBinary = hv.Viper.GetString("fabricBinary")

	scriptPath := hv.Viper.GetString("scriptPath")
	c.scriptPath = filepath.Join(hv.Path, scriptPath)

	// Below is completely unique to configtxgen

	// this gets the configtx.yaml path needed for the
	configtxRelPath := hv.Viper.GetString("configPath")
	// combine the rel. path with the file system path to the
	// hyperspace directory to get the dir containing the configtx.yaml
	c.configPath = filepath.Join(hv.Path, configtxRelPath)

	// read in vars
	c.profile = hv.Viper.GetString("profile")
	c.channelID = hv.Viper.GetString("channelID")

	// combine hyperPath with relative path
	output := hv.Viper.GetString("output")
	c.output = fmt.Sprintf("%s/%s", hv.Path, output)
}

/*
Execute cmdscript
 */
func (c *Configtxgen) Create () {
	command := exec.Command("/bin/bash",
		c.scriptPath,
		"-n", c.network,
		"-b", c.fabricBinary,
		"-c", c.configPath,
		"-p", c.profile,
		"-ch", c.channelID,
		"-o", c.output,
	)

	command.Stdout = os.Stdout

	err := command.Run()
	if err != nil {
		log.Panicf("Cmdscript for configtxgen error: %v", err)
	}
	log.Printf("Configtxgen main script finished with error: %v", err)
}

/**
Initialize a HyperConfigtxgen
 */
func Initialize(rv *util.RootViper)  *Configtxgen {
	ctg := &Configtxgen{}
	ctg.init(rv)
	return ctg
}
