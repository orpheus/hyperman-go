package util

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

/**
A Hyperspace Viper is a spf13/viper instantiated with
a hyperpsace.yaml config
It contains the path to the directory housing the hyperspace.yaml
and a living Viper to control.
 */

/**
Living config files with relevant network information and paths.
HyperVipers don't care about information they hold (the actual
data in the hyperspace.yaml. They just act as living data stores
for import network information relative to their station.
 */
type HyperViper struct {
	// Path to the directory that contains the hyperspace.yaml
	// Could this be removed in vapor of just calling filepath.Dir(viper.ConfigFileUsed())?
	Path string
	// Viper instance for hyperspace config
	// named act so you can access and call like
	// hyperViper.cmd.GetString()
	Viper *viper.Viper
}

func CreateViperYaml (configName string, paths ...string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(configName)
	v.SetConfigType("yaml")

	for _, path := range paths {
		v.AddConfigPath(path)
	}

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}

/**
Spawn a HyperSpaceViper given a path to a hyperspace.yaml
 */
func SpawnHyperSpaceViper (paths ...string) *viper.Viper {
	v := viper.New()
	v.SetConfigName("hyperspace")
	v.SetConfigType("yaml")

	for _, path := range paths {
		v.AddConfigPath(path)
		// adding this so we can run from the cmdscripts dir
		// toDo: remove these relative paths and assume control at all times is based off the control center
		v.AddConfigPath(filepath.Join("../", path))
	}

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}

/**
Creates a HyperViper for a particular squadron/fleet
*/
func CreateHyperViper (path string) *HyperViper {
	hv := &HyperViper{}

	hv.Path = path
	hv.Viper = SpawnHyperSpaceViper(hv.Path)

	return hv
}

type RootViper struct {
	Viper *viper.Viper
	Network string
	NetworkPath string
	NetworkViper *viper.Viper
}

/**
Creates two Vipers, one at Hyperspace Root (the control center)
and another at the Network Root
 */
func CreateRootViper (network string) *RootViper {
	viper := SpawnHyperSpaceViper(".")
	if network == "" {
		network = viper.GetString("defaultNetwork")
	}
	networkPath := filepath.Join("networks", network)
	networkViper := SpawnHyperSpaceViper(networkPath)

	return &RootViper{
		viper,
		network,
		networkPath,
		networkViper,
	}
}
