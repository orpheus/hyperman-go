package hypernet

import (
	"github.com/orpheus/hyperspace/config/configtx"
	"github.com/orpheus/hyperspace/config/orderer"
	"github.com/orpheus/hyperspace/config/peer"
)

// todo: generate any values that can be that aren't explicitly set
// todo: generate data that needs to be generated...
// todo: scriptPaths: default to command-script name
type Hypernet struct {
	Name        string      `yaml:"Name"`
	Netroot     string      `yaml:"Netroot"`
	Peers       []*Peer     `yaml:"peers"`
	Orderers    []*Orderer  `yaml:"orderers"`
	Configtxgen Configtxgen `yaml:"Configtxgen"`
	Cryptogen   Cryptogen   `yaml:"Cryptogen"`
}

type Peer struct {
	BinaryName  string        `yaml:"BinaryName"`
	Config      peer.CoreYaml `yaml:"Config"`
	Environment []string      `yaml:"Environment"`
}

type Orderer struct {
	BinaryName  string              `yaml:"BinaryName"`
	Config      orderer.OrdererYaml `yaml:"Config"`
	Environment []string            `yaml:"Environment"`
}

type Configtxgen struct {
	BinaryName string                `yaml:"BinaryName"`
	Config     configtx.ConfigtxYaml `yaml:"Config"`
	ScriptPath string                `yaml:"ScriptPath"` // path or name of shell script to create configtxgen
	// the following are passed to the binary as flag arguments
	ConfigPath string `yaml:"ConfigPath"`
	Profile    string `yaml:"Profile"`
	ChannelID  string `yaml:"ChannelID"`
	Output     string `yaml:"Output"`
}

type Cryptogen struct {
	BinaryName string            `yaml:"BinaryName"`
	ScriptPath string            `yaml:"ScriptPath"`
	Configs    []*NamePathOutput `yaml:"Configs"`
}

type NamePathOutput struct {
	Name   string `yaml:"Name"`
	Path   string `yaml:"Path"`
	Output string `yaml:"Output"`
}
