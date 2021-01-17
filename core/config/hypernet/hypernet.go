package config

import (
	"fmt"
	config "github.com/orpheus/hyperspace/core/config/peer"
	"github.com/orpheus/hyperspace/core/util"
	"gopkg.in/yaml.v2"
)

type uniqueName string

type Hypernet struct {
	Peers map[uniqueName]struct{
		binaryName string `yaml:"binaryName"`
		config     config.CoreYaml `yaml:"config"`
	} `yaml:"peers"`
	Orderers map[uniqueName]struct{

	} `yaml:"orderers"`
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
		//Config     FabricConfigtxConfig
	}
	// todo: fill out
	Cryptogen struct {

	}
}

func NewHypernet (filePath string) (*Hypernet, error){
	config := &Hypernet{}

	yamlBytes, err := util.ReadInYamlData(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create and load a HypernetYaml from: %v", err)
	}

	err = yaml.Unmarshal(yamlBytes, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling HypernetYaml: %v", err)
	}

	return config, nil
}
