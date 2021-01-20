package cryptogen

import (
	"fmt"
	"github.com/orpheus/hyperspace/util"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	cmdName = "cryptogen"
	cmdDesc = "Generate cryptogen"
)

var cmd = &cobra.Command{
	Use:   cmdName,
	Short: fmt.Sprint(cmdDesc),
	Long:  fmt.Sprint(cmdDesc),
}

// Cmd returns the cobra command for Cryptogen
func Cmd() *cobra.Command {
	//cmd.AddCommand(generateCmd())
	return cmd
}

//func generateCmd() *cobra.Command {
//	return &cobra.Command{
//		Use:   "generate",
//		Short: "Generates cryptogen.",
//		Long:  `Generates cryptogen material for organizations.`,
//		RunE: func(cmd *cobra.Command, args []string) error {
//			if len(args) > 1 {
//				return fmt.Errorf("trailing args detected")
//			}
//			var network string
//			if len(args) == 1 {
//				network = args[0]
//
//				netpath := filepath.Join("networks", network)
//				exists, err := util.FileOrDirectoryExists(netpath)
//				if err != nil {
//					log.Panic(err)
//				}
//				if !exists {
//					log.Fatalf("Network: %v, not found. Path at %s does not exist", network, netpath)
//				}
//			}
//			if network == "" {
//				fmt.Println("Network not specified. Using default network specified in root hyperspace.yaml")
//			}
//			Generate(network)
//			return nil
//		},
//	}
//}
//
//func Generate (network string, config *hypernet.Cryptogen) error {
//	// given a network,
//	// use the NETROOT/network path as the base path
//	// read in hypernet.yaml
//	// check for HYPERSPACE_PATH
//
//	for _, c := range config.Configs {
//		configPath := fmt.Sprintf("configs.%s.path", c.Name)
//		outputPath := fmt.Sprintf("configs.%s.output", c.Name)
//
//		// create paths relative to NETROOT/network
//		configPath = filepath.Join(c.hv.Path, c.hv.Viper.GetString(configPath))
//		outputPath = filepath.Join(c.hv.Path, c.hv.Viper.GetString(outputPath))
//
//		cmd := exec.Command("/bin/bash",
//			config.ScriptPath,
//			"-n", network,
//			"-b", config.BinaryName,
//			"-c", configPath,
//			"-o", outputPath,
//			"-i", c.Name,
//		)
//
//		cmd.Stdout = os.Stdout
//		cmd.Stderr = os.Stderr
//
//		err := cmd.Run()
//		if err != nil {
//			log.Panicf("Error making cryptogen.\n failed on %s", c.Name)
//		}
//		log.Printf("Cryptogen: `generate` finished with error: %v", err)
//	}
//
//	return nil
//}

// -----------------------------------------------
// HyperCryptogen
// -----------------------------------------------
// hyperspace + fabric cryptogen
// -----------------------------------------------
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
		log.Printf("Cryptogen: `generate` finished with error: %v", err)
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
