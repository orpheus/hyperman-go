package config

import (
	"github.com/orpheus/hyperspace/core/util"
	"os"
	"time"
)

//----------------------------------------------------------------------------------
// NewCoreYaml()
//----------------------------------------------------------------------------------
// Creates an instance of a CoreYaml struct and then loads it with a core.yaml.
//----------------------------------------------------------------------------------
func NewConfigtxYaml(filePath string) *ConfigtxYaml {
	config := &ConfigtxYaml{}

	// will call os.exit(1) if error so no need to check for error
	util.UnmarshalYaml(filePath, &config, "Unmarshal New ConfigtxYaml")

	return config
}

//----------------------------------------------------------------------------------
// Write()
//----------------------------------------------------------------------------------
// Writes the config struct to a yaml file given a `filePath` and `perm`
// The yaml file takes on the name of the end of the path given.
//----------------------------------------------------------------------------------
func (c *ConfigtxYaml) Write(filePath string, perm os.FileMode) {
	util.MarshalAndWriteYaml(c, filePath, perm, "Write ConfigtxYaml")
}

//----------------------------------------------------------------------------------
// Merge()
//----------------------------------------------------------------------------------
// Merges the called config struct with another OrdererYaml config struct.
// The called struct gets overridden by the values of the struct passed as the arg.
//----------------------------------------------------------------------------------
func (c *ConfigtxYaml) Merge(config *ConfigtxYaml) {
	util.MergeStructs(c, config, "Merge ConfigtxYaml structs")
}

// ConfigtxYaml consists of the structs used by the configtxgen tool.
type ConfigtxYaml struct {
	Profiles      map[string]*Profile        `yaml:"Profiles"`
	Organizations []*Organization            `yaml:"Organizations"`
	Channel       *Profile                   `yaml:"Channel"`
	Application   *Application               `yaml:"Application"`
	Orderer       *Orderer                   `yaml:"Orderer"`
	Capabilities  map[string]map[string]bool `yaml:"Capabilities"`
}

// Profile encodes orderer/application configuration combinations for the
// configtxgen tool.
type Profile struct {
	Consortium   string                 `yaml:"Consortium"`
	Application  *Application           `yaml:"Application"`
	Orderer      *Orderer               `yaml:"Orderer"`
	Consortiums  map[string]*Consortium `yaml:"Consortiums"`
	Capabilities map[string]bool        `yaml:"Capabilities"`
	Policies     map[string]*Policy     `yaml:"Policies"`
}

// Policy encodes a channel config policy
type Policy struct {
	Type string `yaml:"Type"`
	Rule string `yaml:"Rule"`
}

// Consortium represents a group of organizations which may create channels
// with each other
type Consortium struct {
	Organizations []*Organization `yaml:"Organizations"`
}

// Application encodes the application-level configuration needed in config
// transactions.
type Application struct {
	Organizations []*Organization    `yaml:"Organizations"`
	Capabilities  map[string]bool    `yaml:"Capabilities"`
	Policies      map[string]*Policy `yaml:"Policies"`
	ACLs          map[string]string  `yaml:"ACLs"`
}

// Organization encodes the organization-level configuration needed in
// config transactions.
type Organization struct {
	Name     string             `yaml:"Name"`
	ID       string             `yaml:"ID"`
	MSPDir   string             `yaml:"MSPDir"`
	MSPType  string             `yaml:"MSPType"`
	Policies map[string]*Policy `yaml:"Policies"`

	// Note: Viper deserialization does not seem to care for
	// embedding of types, so we use one organization struct
	// for both orderers and applications.
	AnchorPeers      []*AnchorPeer `yaml:"AnchorPeers"`
	OrdererEndpoints []string      `yaml:"OrdererEndpoints"`

	// AdminPrincipal is deprecated and may be removed in a future release
	// it was used for modifying the default policy generation, but policies
	// may now be specified explicitly so it is redundant and unnecessary
	AdminPrincipal string `yaml:"AdminPrincipal"`

	// SkipAsForeign indicates that this org definition is actually unknown to this
	// instance of the tool, so, parsing of this org's parameters should be ignored.
	SkipAsForeign bool `yaml:"SkipAsForeign"`
}

// AnchorPeer encodes the necessary fields to identify an anchor peer.
type AnchorPeer struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

// Orderer contains configuration associated to a channel.
type Orderer struct {
	OrdererType   string             `yaml:"OrdererType"`
	Addresses     []string           `yaml:"Addresses"`
	BatchTimeout  time.Duration      `yaml:"BatchTimeout"`
	BatchSize     BatchSize          `yaml:"BatchSize"`
	Kafka         Kafka              `yaml:"Kafka"`
	EtcdRaft      *ConfigMetadata    `yaml:"EtcdRaft"`
	Organizations []*Organization    `yaml:"Organizations"`
	MaxChannels   uint64             `yaml:"MaxChannels"`
	Capabilities  map[string]bool    `yaml:"Capabilities"`
	Policies      map[string]*Policy `yaml:"Policies"`
}

// BatchSize contains configuration affecting the size of batches.
type BatchSize struct {
	MaxMessageCount   string `yaml:"MaxMessageCount"`
	AbsoluteMaxBytes  string `yaml:"AbsoluteMaxBytes"`
	PreferredMaxBytes string `yaml:"PreferredMaxBytes"`
}

// Kafka contains configuration for the Kafka-based orderer.
type Kafka struct {
	Brokers []string `yaml:"Brokers"`
}

// ConfigMetadata is serialized and set as the value of ConsensusType.Metadata in
// a channel configuration when the ConsensusType.Type is set "etcdraft".
type ConfigMetadata struct {
	Consenters           []*Consenter `yaml:"Consenters"`
	Options              *Options     `yaml:"Options"`
	XXX_NoUnkeyedLiteral struct{}     `yaml:"-"`
	XXX_unrecognized     []byte       `yaml:"-"`
	XXX_sizecache        int32        `yaml:"-"`
}

// Consenter represents a consenting node (i.e. replica).
type Consenter struct {
	Host                 string   `yaml:"Host"`
	Port                 uint32   `yaml:"Port"`
	ClientTlsCert        string   `yaml:"ClientTLSCert"`
	ServerTlsCert        string   `yaml:"ServerTLSCert"`
	XXX_NoUnkeyedLiteral struct{} `yaml:"XXX_NoUnkeyedLiteral,omitempty"`
	XXX_unrecognized     []byte   `yaml:"XXX_unrecognized,omitempty"`
	XXX_sizecache        int32    `yaml:"XXX_sizecache,omitempty"`
}

// Options to be specified for all the etcd/raft nodes. These can be modified on a
// per-channel basis.
type Options struct {
	TickInterval      string `yaml:"TickInterval"`
	ElectionTick      uint32 `yaml:"ElectionTick"`
	HeartbeatTick     uint32 `yaml:"HeartbeatTick"`
	MaxInflightBlocks uint32 `yaml:"MaxInflightBlocks"`
	// Take snapshot when cumulative data exceeds certain size in bytes.
	SnapshotIntervalSize uint32   `yaml:"SnapshotIntervalSize"`
	XXX_NoUnkeyedLiteral struct{} `yaml:"XXX_NoUnkeyedLiteral,omitempty"`
	XXX_unrecognized     []byte   `yaml:"XXX_unrecognized,omitempty"`
	XXX_sizecache        int32    `yaml:"XXX_sizecache,omitempty"`
}
