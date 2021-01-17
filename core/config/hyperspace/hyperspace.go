package config

import config "github.com/orpheus/hyperspace/core/config/peer"

// todo: fill out fabric config structs
type FabricOrdererConfig struct {}
type FabricConfigtxConfig struct {}

type uniqueName string

type HyperspaceNetworkConfig struct {
	Peers map[uniqueName]struct{
		binaryName string
		config     config.CoreYaml
	}
	Orderers map[uniqueName]struct{

	}
	// todo: generate any values that can be that aren't explicitly set
	Configtxgen struct {
		// binary name, defaults to `configtxgen`
		// todo: default to command-script name
		FabricBinary string
		// path to command-script to create genesis/consortiums
		// todo: default to command-script name
		ScriptPath string
		// The following are the allowed flags passed to the binary
		ConfigPath string
		Profile    string
		ChannelID  string
		Output     string
		Config     FabricConfigtxConfig
	}
	// todo: fill out
	Cryptogen struct {

	}
}

var HysNetConfig = &HyperspaceNetworkConfig{}
// Loads the hys_net_config.yaml into a struct
func (c *HyperspaceNetworkConfig) Load () error {
	// read in yaml file to struct
	return nil
}

// Loads the default fabric `orderer.yaml` config into a new struct
// and returns that struct. Used as a template to override.
func LoadOrdererTemplate () *FabricOrdererConfig {
	config := &FabricOrdererConfig{}
	return config
}

// Loads the default fabric `configtx.yaml` config into a new struct
// and returns that struct. Used as a template to override.
func LoadConfigtxTemplate () *FabricConfigtxConfig {
	config := &FabricConfigtxConfig{}
	return config
}
