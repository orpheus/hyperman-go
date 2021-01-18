package config

import (
	"fmt"
	"github.com/orpheus/hyperspace/core/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//----------------------------------------------------------------------------------
// NewCoreYaml()
//----------------------------------------------------------------------------------
// Creates an instance of a CoreYaml struct and then loads it with a core.yaml.
//----------------------------------------------------------------------------------
func NewCoreYaml (filePath string) (*CoreYaml, error) {
	config := &CoreYaml{}

	yamlBytes, err := util.ReadInYamlData(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create and load a CoreYaml from: %v", err)
	}

	err = yaml.Unmarshal(yamlBytes, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling CoreYaml: %v", err)
	}

	return config, nil
}

//----------------------------------------------------------------------------------
// Write()
//----------------------------------------------------------------------------------
// Writes the config struct to a yaml file given a `filePath` and `perm`
// The yaml file takes on the name of the end of the path given.
//----------------------------------------------------------------------------------
func (c *CoreYaml) Write (filePath string, perm os.FileMode) error {
	d, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("error marshing CoreYaml to bytes: %v", err)
	}

	err = ioutil.WriteFile(filePath,d, perm)
	if err != nil {
		log.Fatalf("error writing CoreYaml to OS: %v", err)

	}
	return nil
}

//----------------------------------------------------------------------------------
// Merge()
//----------------------------------------------------------------------------------
// Merges the called config struct with another CoreYaml config struct.
// The called struct gets overridden by the values of the struct passed as the arg.
//----------------------------------------------------------------------------------
func (c *CoreYaml) Merge (config *CoreYaml) {
	util.MergeStructs(c, config, "Merge CoreYaml configs")
}

//----------------------------------------------------------------------------------
// &CoreYaml{}
//----------------------------------------------------------------------------------
// Yaml skeleton for core.yaml. Used to store fabric peer configurations.
//----------------------------------------------------------------------------------
type CoreYaml struct {
	Peer Peer `yaml:"peer"`
	Vm Vm `yaml:"vm"`
	Chaincode Chaincode `yaml:"chaincode"`
	Ledger Ledger `yaml:"ledger"`
	Operations Operations `yaml:"operations"`
	Metrics Metrics `yaml:"metrics"`
}

type Peer struct {
	ID string `yaml:"id"`
	NetworkId string `yaml:"networkId"`
	ListenAddress string `yaml:"listenAddress"`
	ChaincodeListenAddress string `yaml:"chaincodeListenAddress"`
	ChaincodeAddress string `yaml:"chaincodeAddress"`
	Address string `yaml:"address"`
	AddressAutoDetect bool `yaml:"addressAutoDetect"`
	Keepalive struct {
		Interval time.Duration `yaml:"interval"`
		Timeout time.Duration `yaml:"timeout"`
		MinInterval time.Duration `yaml:"minInterval"`
		Client struct {
			Interval time.Duration `yaml:"interval"`
			Timeout time.Duration `yaml:"timeout"`
		} `yaml:"client"`
		DeliveryClient struct {
			Interval time.Duration `yaml:"interval"`
			Timeout time.Duration `yaml:"timeout"`
		} `yaml:"deliveryClient"`
	} `yaml:"keepalive"`
	Gossip struct {
		Boostrap string `yaml:"bootstrap"`
		UseLeaderElection bool `yaml:"useLeaderElection"`
		OrgLeader bool `yaml:"orgLeader"`
		MembershipTrackerInterval time.Duration `yaml:"membershipTrackerInterval"`
		Endpoint string `yaml:"endpoint"`
		MaxBlockCountToStore int `yaml:"maxBlockCountToStore"`
		MaxPropagationBurstLatency time.Duration `yaml:"maxPropagationBurstLatency"`
		MaxPropagationBurstSize int `yaml:"maxPropagationBurstSize"`
		PropagateIterations int `yaml:"propagateIterations"`
		PropagatePeerNum int `yaml:"propagatePeerNum"`
		PullInterval time.Duration `yaml:"pullInterval"`
		PullPeerNum int `yaml:"pullPeerNum"`
		RequestStateInfoInterval time.Duration `yaml:"requestStateInfoInterval"`
		PublishStateInfoInterval time.Duration `yaml:"publishStateInfoInterval"`
		StateInfoRetentionInterval time.Duration `yaml:"stateInfoRetentionInterval"`
		PublishCertPeriod time.Duration `yaml:"publishCertPeriod"`
		SkipBlockVerification bool `yaml:"skipBlockVerification"`
		DialTimeout time.Duration `yaml:"dialTimeout"`
		ConnTimeout time.Duration `yaml:"connTimeout"`
		RecvBuffSize int `yaml:"recvBuffSize"`
		SendBuffSize int `yaml:"sendBuffSize"`
		DigestWaitTime time.Duration `yaml:"digestWaitTime"`
		RequestWaitTime time.Duration `yaml:"requestWaitTime"`
		ResponseWaitTime time.Duration `yaml:"responseWaitTime"`
		AliveTimeInterval time.Duration `yaml:"aliveTimeInterval"`
		AliveExpirationTimeout time.Duration `yaml:"aliveExpirationTimeout"`
		ReconnectInterval time.Duration `yaml:"reconnectInterval"`
		MaxConnectionAttempts int `yaml:"maxConnectionAttempts"`
		MsgExpirationFactor int `yaml:"msgExpirationFactor"`
		ExternalEndpoint string `yaml:"externalEndpoint"`
		Election struct {
			StartupGracePeriod time.Duration `yaml:"startupGracePeriod"`
			MembershipSampleInterval time.Duration `yaml:"membershipSampleInterval"`
			LeaderAliveThreshold time.Duration `yaml:"leaderAliveThreshold"`
			LeaderElectionDuration time.Duration `yaml:"leaderElectionDuration"`
		} `yaml:"election"`
		PvtData struct {
			PullRetryThreshold time.Duration `yaml:"pullRetryThreshold"`
			TransientstoreMaxBlockRetention int `yaml:"transientstoreMaxBlockRetention"`
			PushAckTimeout time.Duration `yaml:"pushAckTimeout"`
			BtlPullMargin int `yaml:"btlPullMargin"`
			ReconcileBatchSize int `yaml:"reconcileBatchSize"`
			ReconcileSleepInterval time.Duration `yaml:"reconcileSleepInterval"`
			ReconciliationEnabled bool `yaml:"reconciliationEnabled"`
			SkipPullingInvalidTransactionsDuringCommit bool `yaml:"skipPullingInvalidTransactionsDuringCommit"`
			ImplicitCollectionDisseminationPolicy struct {
				RequiredPeerCount int `yaml:"requiredPeerCount"`
				MaxPeerCount int `yaml:"maxPeerCount"`
			} `yaml:"implicitCollectionDisseminationPolicy"`
		} `yaml:"pvtData"`
		State struct {
			Enabled bool `yaml:"enabled"`
			CheckInterval time.Duration `yaml:"checkInterval"`
			ResponseTimeout time.Duration `yaml:"responseTimeout"`
			BatchSize int `yaml:"batchSize"`
			BlockBufferSize int `yaml:"blockBufferSize"`
			MaxRetries int `yaml:"maxRetries"`
		} `yaml:"state"`
	} `yaml:"gossip"`
	Tls struct {
		Enabled bool `yaml:"enabled"`
		ClientAuthRequired bool `yaml:"clientAuthRequired"`
		Cert struct {
			File string `yaml:"file"`
		} `yaml:"cert"`
		Key struct {
			File string `yaml:"file"`
		} `yaml:"key"`
		Rootcert struct {
			File string `yaml:"file"`
		} `yaml:"rootcert"`
		ClientRootCAs struct {
			Files []string `yaml:"files"`
		} `yaml:"clientRootCAs"`
		ClientKey struct {
			File string `yaml:"file"`
		} `yaml:"clientKey"`
		ClientCert struct {
			File string `yaml:"file"`
		} `yaml:"clientCert"`
	} `yaml:"tls"`
	Authentication struct {
		Timewindow time.Duration `yaml:"timewindow"`
	} `yaml:"authentication"`
	FileSystemPath string `yaml:"fileSystemPath"`
	BCCSP struct {
		Default string `yaml:"Default"`
		SW struct {
			Hash string `yaml:"Hash"`
			Security string `yaml:"Security"` // should this be an int ?
			FileKeyStore struct {
				KeyStore string `yaml:"KeyStore"`
			} `yaml:"FileKeyStore"`
		} `yaml:"SW"`
		PKCS11 struct {
			Library string `yaml:"Library"`
			Label string `yaml:"Label"`
			Pin string `yaml:"Pin"`
			Hash string `yaml:"Hash"`
			Security string `yaml:"Security"`
		} `yaml:"PKCS11"`
	} `yaml:"BCCSP"`
	MspConfigPath string `yaml:"mspConfigPath"`
	LocalMspId string `yaml:"localMspId"`
	Client struct {
		ConnTimeout time.Duration `yaml:"connTimeout"`
	} `yaml:"client"`
	Deliveryclient struct {
		ReconnectTotalTimeThreshold time.Duration `yaml:"reconnectTotalTimeThreshold"`
		ConnTimeout time.Duration `yaml:"connTimeout"`
		ReConnectBackoffThreshold time.Duration `yaml:"reConnectBackoffThreshold"`
		AddressOverrides []struct{
			From string `yaml:"from"`
			To string `yaml:"to"`
			CaCertsFile string `yaml:"caCertsFile"`
		} `yaml:"addressOverrides"`
	} `yaml:"deliveryclient"`
	LocalMspType string `yaml:"localMspType"`
	Profile struct {
		Enabled bool `yaml:"enabled"`
		ListenAddress string `yaml:"listenAddress"`
	} `yaml:"profile"`
	Handlers struct {
		AuthFilters []struct{
			Name string `yaml:"name"`
		} `yaml:"authFilters"`
		Decorators []struct{
			Name string `yaml:"name"`
		} `yaml:"decorators"`
		Endorsers struct {
			Escc struct {
				Name string `yaml:"name"`
				Library string `yaml:"library"`
			} `yaml:"escc"`
		} `yaml:"endorsers"`
		Validators struct {
			Vscc struct {
				Name string `yaml:"name"`
				Library string `yaml:"library"`
			} `yaml:"vscc"`
		} `yaml:"validators"`
		Library string `yaml:"library"`
	} `yaml:"handlers"`
	ValidatorPoolSize string `yaml:"validatorPoolSize"`
	Discovery struct {
		Enabled bool `yaml:"enabled"`
		AuthCacheEnabled bool `yaml:"authCacheEnabled"`
		AuthCacheMaxSize int `yaml:"authCacheMaxSize"`
		AuthCachePurgeRetentionRatio float64 `yaml:"authCachePurgeRetentionRatio"`
		OrgMembersAllowedAccess bool `yaml:"orgMembersAllowedAccess"`
	} `yaml:"discovery"`
	Limits struct {
		Concurrency struct {
			EndorserService int `yaml:"endorserService"`
			DeliverService int `yaml:"deliverService"`
		} `yaml:"concurrency"`
	} `yaml:"limits"`
}

type Vm struct {
	Endpoint string `yaml:"endpoint"`
	Docker struct {
		Tls struct {
			Enabled bool `yaml:"enabled"`
			Ca struct {
				File string `yaml:"file"`
			} `yaml:"ca"`
			Cert struct {
				File string `yaml:"file"`
			} `yaml:"cert"`
			Key struct {
				File string `yaml:"file"`
			} `yaml:"key"`
		} `yaml:"tls"`
		AttachStdout bool `yaml:"attachStdout"`
		HostConfig struct {
			NetworkMode string `yaml:"NetworkMode"`
			Dns []string `yaml:"Dns"`
			LogConfig	struct {
				Type string `yaml:"Type"`
				Config struct {
					Maxsize string `yaml:"max-size"`
					Maxfile string `yaml:"max-file"`
				} `yaml:"Config"`
			} `yaml:"LogConfig"`
			Memory int `yaml:"Memory"`
		} `yaml:"hostConfig"`
	} `yaml:"docker"`
}

type Chaincode struct {
	Id struct {
		Path string `yaml:"path"`
		Name string `yaml:"name"`
	} `yaml:"id"`
	Builder string `yaml:"builder"`
	Pull bool `yaml:"pull"`
	Golang struct {
		Runtime string `yaml:"runtime"`
		DynamicLink bool `yaml:"dynamicLink"`
	} `yaml:"golang"`
	Java struct {
		Runtime string `yaml:"runtime"`
	} `yaml:"java"`
	Node struct {
		Runtime string `yaml:"runtime"`
	} `yaml:"node"`
	ExternalBuilders []struct{
		Path string `yaml:"path"`
		Name string `yaml:"name"`
		PropagateEnvironment []string `yaml:"propagateEnvironment"`
	} `yaml:"externalBuilders"`
	InstallTimeout time.Duration `yaml:"installTimeout"`
	Startuptimeout time.Duration `yaml:"startuptimeout"`
	Executetimeout time.Duration `yaml:"executetimeout"`
	Mode string `yaml:"mode"`
	Keepalive int `yaml:"keepalive"`
	System struct {
		Lifecycle string `yaml:"_lifecycle"`
		Cscc string `yaml:"cscc"`
		Lscc string `yaml:"lscc"`
		Qscc string `yaml:"qscc"`
	} `yaml:"system"`
	Logging struct {
		Level string `yaml:"level"`
		Shim string `yaml:"shim"`
		Format string `yaml:"format"`
	} `yaml:"logging"`
}

type Ledger struct {
	Blockchain string `yaml:"blockchain"` // {}Interface?
	State struct {
		StateDatabase string `yaml:"stateDatabase"`
		TotalQueryLimit int `yaml:"totalQueryLimit"`
		CouchDBConfig struct {
			CouchDBAddress string `yaml:"couchDBAddress"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			MaxRetries int `yaml:"maxRetries"`
			MaxRetriesOnStartup int `yaml:"maxRetriesOnStartup"`
			RequestTimeout time.Duration `yaml:"requestTimeout"`
			InternalQueryLimit int `yaml:"internalQueryLimit"`
			MaxBatchUpdateSize int `yaml:"maxBatchUpdateSize"`
			WarmIndexesAfterNBlocks int `yaml:"warmIndexesAfterNBlocks"`
			CreateGlobalChangesDB bool `yaml:"createGlobalChangesDB"`
			CacheSize int `yaml:"cacheSize"`
		} `yaml:"couchDBConfig"`
	} `yaml:"state"`
	History struct {
		EnableHistoryDatabase bool `yaml:"enableHistoryDatabase"`
	} `yaml:"history"`
	PvtdataStore struct {
		CollElgProcMaxDbBatchSize int `yaml:"collElgProcMaxDbBatchSize"`
		CollElgProcDbBatchesInterval int `yaml:"collElgProcDbBatchesInterval"`
		DeprioritizedDataReconcilerInterval time.Duration `yaml:"deprioritizedDataReconcilerInterval"`
	} `yaml:"pvtdataStore"`
	Snapshots struct {
		RootDir string `yaml:"rootDir"`
	} `yaml:"snapshots"`
}

type Operations struct {
	ListenAddress string `yaml:"listenAddress"`
	Tls struct {
		Enabled bool `yaml:"enabled"`
		Cert struct {
			File string `yaml:"file"`
		} `yaml:"cert"`
		Key struct {
			File string `yaml:"file"`
		} `yaml:"key"`
		ClientAuthRequired bool `yaml:"clientAuthRequired"`
		ClientRootCAs struct {
			Files []string `yaml:"files"`
		}  `yaml:"clientRootCAs"`
	} `yaml:"tls"`
}

type Metrics struct {
	Provider string `yaml:"provider"`
	Statsd struct {
		Network string `yaml:"network"`
		Address string `yaml:"address"`
		WriteInterval time.Duration `yaml:"writeInterval"`
		Prefix string `yaml:"prefix"`
	} `yaml:"statsd"`
}