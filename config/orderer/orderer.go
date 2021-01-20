package orderer

import (
	"github.com/orpheus/hyperspace/util"
	"os"
	"time"
)

//----------------------------------------------------------------------------------
// NewCoreYaml()
//----------------------------------------------------------------------------------
// Creates an instance of a CoreYaml struct and then loads it with a core.yaml.
//----------------------------------------------------------------------------------
func NewOrdererYaml(filePath string) *OrdererYaml {
	config := &OrdererYaml{}

	// will call os.exit(1) if error so no need to check for error
	util.UnmarshalYaml(filePath, &config, "Unmarshal New OrdererYaml")

	return config
}

//----------------------------------------------------------------------------------
// Write()
//----------------------------------------------------------------------------------
// Writes the config struct to a yaml file given a `filePath` and `perm`
// The yaml file takes on the name of the end of the path given.
//----------------------------------------------------------------------------------
func (c *OrdererYaml) Write(filePath string, perm os.FileMode) {
	util.MarshalAndWriteYaml(c, filePath, perm, "Write OrdererYaml")
}

//----------------------------------------------------------------------------------
// Merge()
//----------------------------------------------------------------------------------
// Merges the called config struct with another OrdererYaml config struct.
// The called struct gets overridden by the values of the struct passed as the arg.
//----------------------------------------------------------------------------------
func (c *OrdererYaml) Merge(config *OrdererYaml) {
	util.MergeStructs(c, config, "Merge OrdererYaml structs")
}

// OrdererYaml directly corresponds to the orderer config YAML.
type OrdererYaml struct {
	General    General    `yaml:"General"`
	FileLedger FileLedger `yaml:"FileLedger"`
	Kafka      Kafka      `yaml:"Kafka"`
	Debug      Debug      `yaml:"Debug"`
	// if the dynamism of Consensus proves to be difficult,
	// change it to be a map[string]string and for the Merge fn
	// loop over the override struct and set the keys and values
	// make sure to remove any default values on the template struct
	// that aren't specified in the override struct.
	Consensus            Consensus            `yaml:"Consensus"`
	Operations           Operations           `yaml:"Operations"`
	Metrics              Metrics              `yaml:"Metrics"`
	ChannelParticipation ChannelParticipation `yaml:"ChannelParticipation"`
	Admin                Admin                `yaml:"Admin"`
}

// General contains config which should be common among all orderer types.
type General struct {
	ListenAddress     string         `yaml:"ListenAddress"`
	ListenPort        uint16         `yaml:"ListenPort"`
	TLS               TLS            `yaml:"TLS"`
	Cluster           Cluster        `yaml:"Cluster"`
	Keepalive         Keepalive      `yaml:"Keepalive"`
	ConnectionTimeout time.Duration  `yaml:"ConnectionTimeout"`
	GenesisMethod     string         `yaml:"GenesisMethod"` // For compatibility only, will be replaced by BootstrapMethod
	GenesisFile       string         `yaml:"GenesisFile"`   // For compatibility only, will be replaced by BootstrapFile
	BootstrapMethod   string         `yaml:"BootstrapMethod"`
	BootstrapFile     string         `yaml:"BootstrapFile"`
	Profile           Profile        `yaml:"Profile"`
	LocalMSPDir       string         `yaml:"LocalMSPDir"`
	LocalMSPID        string         `yaml:"LocalMSPID"`
	BCCSP             *FactoryOpts   `yaml:"BCCSP"`
	Authentication    Authentication `yaml:"Authentication"`
}

// FactoryOpts holds configuration information used to initialize factory implementations
type FactoryOpts struct {
	Default string      `json:"default" yaml:"Default"`
	SW      *SwOpts     `json:"SW,omitempty" yaml:"SW,omitempty"`
	PKCS11  *Pkcs11Opts `yaml:"PKCS11"` // added this myself because orderer.yaml contains it
}

// SwOpts contains options for the SWFactory
type SwOpts struct {
	// Default algorithms when not specified (Deprecated?)
	Security     int               `json:"security" yaml:"Security"`
	Hash         string            `json:"hash" yaml:"Hash"`
	FileKeystore *FileKeystoreOpts `json:"filekeystore,omitempty" yaml:"FileKeyStore,omitempty"`
}

// Pkcs11Opts added this myself because orderer.yaml has it but the fabric struct doesn't
type Pkcs11Opts struct {
	Library      string            `yaml:"Library"`
	Label        string            `yaml:"Label"`
	Pin          string            `yaml:"Pin"`
	Hash         string            `yaml:"Hash"`
	Security     int               `yaml:"Security"`
	FileKeystore *FileKeystoreOpts `json:"filekeystore,omitempty" yaml:"FileKeyStore,omitempty"`
}

// Pluggable Keystores, could add JKS, P12, etc..
type FileKeystoreOpts struct {
	KeyStorePath string `yaml:"KeyStore"`
}

type Cluster struct {
	ListenAddress                        string        `yaml:"ListenAddress"`
	ListenPort                           uint16        `yaml:"ListenPort"`
	ServerCertificate                    string        `yaml:"ServerCertificate"`
	ServerPrivateKey                     string        `yaml:"ServerPrivateKey"`
	ClientCertificate                    string        `yaml:"ClientCertificate"`
	ClientPrivateKey                     string        `yaml:"ClientPrivateKey"`
	RootCAs                              []string      `yaml:"RootCAs"`
	DialTimeout                          time.Duration `yaml:"DialTimeout"`
	RPCTimeout                           time.Duration `yaml:"RPCTimeout"`
	ReplicationBufferSize                int           `yaml:"ReplicationBufferSize"`
	ReplicationPullTimeout               time.Duration `yaml:"ReplicationPullTimeout"`
	ReplicationRetryTimeout              time.Duration `yaml:"ReplicationRetryTimeout"`
	ReplicationBackgroundRefreshInterval time.Duration `yaml:"ReplicationBackgroundRefreshInterval"`
	ReplicationMaxRetries                int           `yaml:"ReplicationMaxRetries"`
	SendBufferSize                       int           `yaml:"SendBufferSize"`
	CertExpirationWarningThreshold       time.Duration `yaml:"CertExpirationWarningThreshold"`
	TLSHandshakeTimeShift                time.Duration `yaml:"TLSHandshakeTimeShift"`
}

// Keepalive contains configuration for gRPC servers.
type Keepalive struct {
	ServerMinInterval time.Duration `yaml:"ServerMinInterval"`
	ServerInterval    time.Duration `yaml:"ServerInterval"`
	ServerTimeout     time.Duration `yaml:"ServerTimeout"`
}

// TLS contains configuration for TLS connections.
type TLS struct {
	Enabled               bool          `yaml:"Enabled"`
	PrivateKey            string        `yaml:"PrivateKey"`
	Certificate           string        `yaml:"Certificate"`
	RootCAs               []string      `yaml:"RootCAs"`
	ClientAuthRequired    bool          `yaml:"ClientAuthRequired"`
	ClientRootCAs         []string      `yaml:"ClientRootCAs"`
	TLSHandshakeTimeShift time.Duration `yaml:"TLSHandshakeTimeShift"`
}

// SASLPlain contains configuration for SASL/PLAIN authentication
type SASLPlain struct {
	Enabled  bool   `yaml:"Enabled"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

// Authentication contains configuration parameters related to authenticating
// client messages.
type Authentication struct {
	TimeWindow         time.Duration `yaml:"TimeWindow"`
	NoExpirationChecks bool          `yaml:"NoExpirationChecks"`
}

// Profile contains configuration for Go pprof profiling.
type Profile struct {
	Enabled bool   `yaml:"Enabled"`
	Address string `yaml:"Address"`
}

// FileLedger contains configuration for the file-based ledger.
type FileLedger struct {
	Location string `yaml:"Location"`
	Prefix   string `yaml:"Prefix"` // For compatibility only. This setting is no longer supported.
}

// Kafka contains configuration for the Kafka-based orderer.
type Kafka struct {
	Retry     Retry     `yaml:"Retry"`
	Verbose   bool      `yaml:"Verbose"`
	Version   string    `yaml:"Version"` // TODO Move this to global config
	TLS       TLS       `yaml:"TLS"`
	SASLPlain SASLPlain `yaml:"SASLPlain"`
	Topic     Topic     `yaml:"Topic"`
}

// Retry contains configuration related to retries and timeouts when the
// connection to the Kafka cluster cannot be established, or when Metadata
// requests needs to be repeated (because the cluster is in the middle of a
// leader election).
type Retry struct {
	ShortInterval   time.Duration   `yaml:"ShortInterval"`
	ShortTotal      time.Duration   `yaml:"ShortTotal"`
	LongInterval    time.Duration   `yaml:"LongInterval"`
	LongTotal       time.Duration   `yaml:"LongTotal"`
	NetworkTimeouts NetworkTimeouts `yaml:"NetworkTimeouts"`
	Metadata        Metadata        `yaml:"Metadata"`
	Producer        Producer        `yaml:"Producer"`
	Consumer        Consumer        `yaml:"Consumer"`
}

// NetworkTimeouts contains the socket timeouts for network requests to the
// Kafka cluster.
type NetworkTimeouts struct {
	DialTimeout  time.Duration `yaml:"DialTimeout"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

// Metadata contains configuration for the metadata requests to the Kafka
// cluster.
type Metadata struct {
	RetryMax     int           `yaml:"RetryMax"`
	RetryBackoff time.Duration `yaml:"RetryBackoff"`
}

// Producer contains configuration for the producer's retries when failing to
// post a message to a Kafka partition.
type Producer struct {
	RetryMax     int           `yaml:"RetryMax"`
	RetryBackoff time.Duration `yaml:"RetryBackoff"`
}

// Consumer contains configuration for the consumer's retries when failing to
// read from a Kafa partition.
type Consumer struct {
	RetryBackoff time.Duration `yaml:"RetryBackoff"`
}

// Topic contains the settings to use when creating Kafka topics
type Topic struct {
	ReplicationFactor int16 `yaml:"ReplicationFactor"`
}

// Debug contains configuration for the orderer's debug parameters.
type Debug struct {
	BroadcastTraceDir string `yaml:"BroadcastTraceDir"`
	DeliverTraceDir   string `yaml:"DeliverTraceDir"`
}

// The allowed key-value pairs here depend on consensus plugin. Even
// though this struct is dynamic (interface{}), Fabric has to look
// somewhere for specific values, so just add all possible Consensus
// key-pairs here so we can read and write without hassle.
// todo: update with all possible Consensus fields
type Consensus struct {
	WALDir  string `yaml:"WALDir"`
	SnapDir string `yaml:"SnapDir"`
}

// Operations configures the operations endpoint for the orderer.
type Operations struct {
	ListenAddress string `yaml:"ListenAddress"`
	TLS           TLS    `yaml:"TLS"`
}

// Metrics configures the metrics provider for the orderer.
type Metrics struct {
	Provider string `yaml:"Provider"`
	Statsd   Statsd `yaml:"Statsd"`
}

// Statsd provides the configuration required to emit statsd metrics from the orderer.
type Statsd struct {
	Network       string        `yaml:"Network"`
	Address       string        `yaml:"Address"`
	WriteInterval time.Duration `yaml:"WriteInterval"`
	Prefix        string        `yaml:"Prefix"`
}

// Admin configures the admin endpoint for the orderer.
type Admin struct {
	ListenAddress string `yaml:"ListenAddress"`
	TLS           TLS    `yaml:"TLS"`
}

// ChannelParticipation provides the channel participation API configuration for the orderer.
// Channel participation uses the same ListenAddress and TLS settings of the Operations service.
type ChannelParticipation struct {
	Enabled            bool   `yaml:"Enabled"`
	MaxRequestBodySize string `yaml:"MaxRequestBodySize"` // changed from uint32 so it's easier to read
}
