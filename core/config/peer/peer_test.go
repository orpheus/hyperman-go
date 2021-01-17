package config

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func dur (t string) time.Duration {
	d, _ := time.ParseDuration(t)
	return d
}

func TestLoadCoreYaml (t *testing.T) {
	coreYaml, err := NewCoreYaml("core.yaml")
	if err != nil {
		t.Errorf("Failed to create CoreYaml: %v", err)
	}

	require.Equal(t, coreYaml.Peer.ID, "jdoe")
	require.Equal(t, coreYaml.Peer.NetworkId, "dev")
	require.Equal(t, coreYaml.Peer.ListenAddress, "0.0.0.0:7051")
	require.Equal(t, coreYaml.Peer.Address, "0.0.0.0:7051")
	require.Equal(t, coreYaml.Peer.ChaincodeAddress, "")
	require.Empty(t, coreYaml.Peer.ChaincodeListenAddress)
	require.Empty(t, coreYaml.Peer.ChaincodeListenAddress)
	require.False(t, coreYaml.Peer.AddressAutoDetect)

	require.Equal(t, coreYaml.Peer.Keepalive.Interval, dur("7200s"))
	require.Equal(t, coreYaml.Peer.Keepalive.Timeout, dur("20s"))
	require.Equal(t, coreYaml.Peer.Keepalive.MinInterval, dur("60s"))
	require.Equal(t, coreYaml.Peer.Keepalive.Client.Interval, dur("60s"))
	require.Equal(t, coreYaml.Peer.Keepalive.Client.Timeout, dur("20s"))
	require.Equal(t, coreYaml.Peer.Keepalive.DeliveryClient.Interval, dur("60s"))
	require.Equal(t, coreYaml.Peer.Keepalive.DeliveryClient.Timeout, dur("20s"))

	require.Equal(t, coreYaml.Peer.Gossip.Boostrap, "127.0.0.1:7051")
	require.False(t, coreYaml.Peer.Gossip.UseLeaderElection)
	require.True(t, coreYaml.Peer.Gossip.OrgLeader)
	require.Equal(t, coreYaml.Peer.Gossip.MembershipTrackerInterval, dur("5s"))
	require.Empty(t, coreYaml.Peer.Gossip.Endpoint)
	require.Equal(t, coreYaml.Peer.Gossip.MaxBlockCountToStore, 10)
	require.Equal(t, coreYaml.Peer.Gossip.MaxPropagationBurstLatency, dur("10ms"))
	require.Equal(t, coreYaml.Peer.Gossip.MaxPropagationBurstSize, 10)
	require.Equal(t, coreYaml.Peer.Gossip.PropagateIterations, 1)
	require.Equal(t, coreYaml.Peer.Gossip.PropagatePeerNum, 3)
	require.Equal(t, coreYaml.Peer.Gossip.PullInterval, dur("4s"))
	require.Equal(t, coreYaml.Peer.Gossip.PullPeerNum, 3)
	require.Equal(t, coreYaml.Peer.Gossip.RequestStateInfoInterval, dur("4s"))
	require.Equal(t, coreYaml.Peer.Gossip.PublishStateInfoInterval, dur("4s"))
	require.Empty(t, coreYaml.Peer.Gossip.StateInfoRetentionInterval)
	require.Equal(t, coreYaml.Peer.Gossip.PublishCertPeriod, dur("10s"))
	require.False(t, coreYaml.Peer.Gossip.SkipBlockVerification)
	require.Equal(t, coreYaml.Peer.Gossip.DialTimeout, dur("3s"))
	require.Equal(t, coreYaml.Peer.Gossip.ConnTimeout, dur("2s"))
	require.Equal(t, coreYaml.Peer.Gossip.RecvBuffSize, 20)
	require.Equal(t, coreYaml.Peer.Gossip.SendBuffSize, 200)
	require.Equal(t, coreYaml.Peer.Gossip.DigestWaitTime, dur("1s"))
	require.Equal(t, coreYaml.Peer.Gossip.RequestWaitTime, dur("1500ms"))
	require.Equal(t, coreYaml.Peer.Gossip.ResponseWaitTime, dur("2s"))
	require.Equal(t, coreYaml.Peer.Gossip.AliveTimeInterval, dur("5s"))
	require.Equal(t, coreYaml.Peer.Gossip.AliveExpirationTimeout, dur("25s"))
	require.Equal(t, coreYaml.Peer.Gossip.ReconnectInterval, dur("25s"))
	require.Equal(t, coreYaml.Peer.Gossip.MaxConnectionAttempts, 120)
	require.Equal(t, coreYaml.Peer.Gossip.MsgExpirationFactor, 20)
	require.Empty(t, coreYaml.Peer.Gossip.ExternalEndpoint)
	require.Equal(t, coreYaml.Peer.Gossip.Election.StartupGracePeriod, dur("15s"))
	require.Equal(t, coreYaml.Peer.Gossip.Election.MembershipSampleInterval, dur("1s"))
	require.Equal(t, coreYaml.Peer.Gossip.Election.LeaderAliveThreshold, dur("10s"))
	require.Equal(t, coreYaml.Peer.Gossip.Election.LeaderElectionDuration, dur("5s"))
	require.Equal(t, coreYaml.Peer.Gossip.PvtData.PullRetryThreshold, dur("60s"))
	require.Equal(t, coreYaml.Peer.Gossip.PvtData.TransientstoreMaxBlockRetention, 1000)
	require.Equal(t, coreYaml.Peer.Gossip.PvtData.PushAckTimeout, dur("3s"))
	require.Equal(t, coreYaml.Peer.Gossip.PvtData.BtlPullMargin, 10)
	require.Equal(t, coreYaml.Peer.Gossip.PvtData.ReconcileBatchSize, 10)
	require.Equal(t, coreYaml.Peer.Gossip.PvtData.ReconcileSleepInterval, dur("1m"))
	require.True(t, coreYaml.Peer.Gossip.PvtData.ReconciliationEnabled)
	require.False(t, coreYaml.Peer.Gossip.PvtData.SkipPullingInvalidTransactionsDuringCommit)
	require.Equal(t, coreYaml.Peer.Gossip.PvtData.ImplicitCollectionDisseminationPolicy.RequiredPeerCount, 0)
	require.Equal(t, coreYaml.Peer.Gossip.PvtData.ImplicitCollectionDisseminationPolicy.MaxPeerCount, 1)
	require.False(t, coreYaml.Peer.Gossip.State.Enabled)
	require.Equal(t, coreYaml.Peer.Gossip.State.CheckInterval, dur("10s"))
	require.Equal(t, coreYaml.Peer.Gossip.State.ResponseTimeout, dur("3s"))
	require.Equal(t, coreYaml.Peer.Gossip.State.BatchSize, 10)
	require.Equal(t, coreYaml.Peer.Gossip.State.BlockBufferSize, 20)
	require.Equal(t, coreYaml.Peer.Gossip.State.MaxRetries, 3)

	require.False(t, coreYaml.Peer.Tls.Enabled)
	require.False(t, coreYaml.Peer.Tls.ClientAuthRequired)
	require.Equal(t, coreYaml.Peer.Tls.Cert.File, "tls/server.crt")
	require.Equal(t, coreYaml.Peer.Tls.Key.File, "tls/server.key")
	require.Equal(t, coreYaml.Peer.Tls.Rootcert.File, "tls/ca.crt")
	require.Equal(t, coreYaml.Peer.Tls.ClientRootCAs.Files[0], "tls/ca.crt")
	require.Empty(t, coreYaml.Peer.Tls.ClientKey.File)
	require.Empty(t, coreYaml.Peer.Tls.ClientCert.File)

	require.Equal(t, coreYaml.Peer.Authentication.Timewindow, dur("15m"))
	require.Equal(t, coreYaml.Peer.FileSystemPath, "/var/hyperledger/production")

	require.Equal(t, coreYaml.Peer.BCCSP.Default, "SW")
	require.Equal(t, coreYaml.Peer.BCCSP.SW.Hash, "SHA2")
	require.Equal(t, coreYaml.Peer.BCCSP.SW.Security, "256")
	require.Empty(t, coreYaml.Peer.BCCSP.SW.FileKeyStore.KeyStore)
	require.Empty(t, coreYaml.Peer.BCCSP.PKCS11.Library)
	require.Empty(t, coreYaml.Peer.BCCSP.PKCS11.Label)
	require.Empty(t, coreYaml.Peer.BCCSP.PKCS11.Pin)
	require.Empty(t, coreYaml.Peer.BCCSP.PKCS11.Hash)
	require.Empty(t, coreYaml.Peer.BCCSP.PKCS11.Security)

	require.Equal(t, coreYaml.Peer.MspConfigPath, "msp")
	require.Equal(t, coreYaml.Peer.LocalMspId, "SampleOrg")
	require.Equal(t, coreYaml.Peer.Client.ConnTimeout, dur("3s"))

	require.Equal(t, coreYaml.Peer.Deliveryclient.ReConnectBackoffThreshold, dur("3600s"))
	require.Equal(t, coreYaml.Peer.Deliveryclient.ReconnectTotalTimeThreshold, dur("3600s"))
	require.Equal(t, coreYaml.Peer.Deliveryclient.ConnTimeout, dur("3s"))
	require.Empty(t, coreYaml.Peer.Deliveryclient.AddressOverrides)

	require.Equal(t, coreYaml.Peer.LocalMspType, "bccsp")

	require.False(t, coreYaml.Peer.Profile.Enabled)
	require.Equal(t, coreYaml.Peer.Profile.ListenAddress, "0.0.0.0:6060")

	require.Equal(t, coreYaml.Peer.Handlers.AuthFilters[0].Name, "DefaultAuth")
	require.Equal(t, coreYaml.Peer.Handlers.AuthFilters[1].Name, "ExpirationCheck")
	require.Equal(t, coreYaml.Peer.Handlers.Decorators[0].Name, "DefaultDecorator")
	require.Equal(t, coreYaml.Peer.Handlers.Endorsers.Escc.Name, "DefaultEndorsement")
	require.Empty(t, coreYaml.Peer.Handlers.Endorsers.Escc.Library)
	require.Equal(t, coreYaml.Peer.Handlers.Validators.Vscc.Name, "DefaultValidation")
	require.Empty(t, coreYaml.Peer.Handlers.Validators.Vscc.Library)

	require.Empty(t, coreYaml.Peer.ValidatorPoolSize)

	require.True(t, coreYaml.Peer.Discovery.Enabled)
	require.True(t, coreYaml.Peer.Discovery.AuthCacheEnabled)
	require.Equal(t, coreYaml.Peer.Discovery.AuthCacheMaxSize, 1000)
	require.Equal(t, coreYaml.Peer.Discovery.AuthCachePurgeRetentionRatio, 0.75)
	require.False(t, coreYaml.Peer.Discovery.OrgMembersAllowedAccess)

	require.Equal(t, coreYaml.Peer.Limits.Concurrency.EndorserService, 2500)
	require.Equal(t, coreYaml.Peer.Limits.Concurrency.DeliverService, 2500)

	require.Equal(t, coreYaml.Vm.Endpoint, "unix:///var/run/docker.sock")
	require.False(t, coreYaml.Vm.Docker.Tls.Enabled)
	require.Equal(t, coreYaml.Vm.Docker.Tls.Ca.File, "docker/ca.crt")
	require.Equal(t, coreYaml.Vm.Docker.Tls.Cert.File, "docker/tls.crt")
	require.Equal(t, coreYaml.Vm.Docker.Tls.Key.File, "docker/tls.key")
	require.False(t, coreYaml.Vm.Docker.AttachStdout)
	require.Equal(t, coreYaml.Vm.Docker.HostConfig.NetworkMode, "host")
	require.Empty(t, coreYaml.Vm.Docker.HostConfig.Dns)
	require.Equal(t, coreYaml.Vm.Docker.HostConfig.LogConfig.Type, "json-file")
	require.Equal(t, coreYaml.Vm.Docker.HostConfig.LogConfig.Config.Maxsize, "50m")
	require.Equal(t, coreYaml.Vm.Docker.HostConfig.LogConfig.Config.Maxfile, "5")
	require.Equal(t, coreYaml.Vm.Docker.HostConfig.Memory, 2147483648)

	require.Empty(t, coreYaml.Chaincode.Id.Path)
	require.Empty(t, coreYaml.Chaincode.Id.Name)
	require.Equal(t, coreYaml.Chaincode.Builder, "$(DOCKER_NS)/fabric-ccenv:$(TWO_DIGIT_VERSION)")
	require.False(t, coreYaml.Chaincode.Pull)
	require.Equal(t, coreYaml.Chaincode.Golang.Runtime, "$(DOCKER_NS)/fabric-baseos:$(TWO_DIGIT_VERSION)")
	require.False(t, coreYaml.Chaincode.Golang.DynamicLink)
	require.Equal(t, coreYaml.Chaincode.Java.Runtime, "$(DOCKER_NS)/fabric-javaenv:$(TWO_DIGIT_VERSION)")
	require.Equal(t, coreYaml.Chaincode.Node.Runtime, "$(DOCKER_NS)/fabric-javaenv:$(TWO_DIGIT_VERSION)")
	require.Empty(t, coreYaml.Chaincode.ExternalBuilders)
	require.Equal(t, coreYaml.Chaincode.InstallTimeout, dur("300s"))
	require.Equal(t, coreYaml.Chaincode.Startuptimeout, dur("300s"))
	require.Equal(t, coreYaml.Chaincode.Executetimeout, dur("30s"))
	require.Equal(t, coreYaml.Chaincode.Mode, "net")
	require.Equal(t, coreYaml.Chaincode.Keepalive, 0)
	require.Equal(t, coreYaml.Chaincode.System.Lifecycle, "enable")
	require.Equal(t, coreYaml.Chaincode.System.Cscc, "enable")
	require.Equal(t, coreYaml.Chaincode.System.Lscc, "enable")
	require.Equal(t, coreYaml.Chaincode.System.Qscc, "enable")
	require.Equal(t, coreYaml.Chaincode.Logging.Level, "info")
	require.Equal(t, coreYaml.Chaincode.Logging.Shim, "warning")
	require.Equal(t, coreYaml.Chaincode.Logging.Format, "%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

	require.Empty(t, coreYaml.Ledger.Blockchain)
	require.Equal(t, coreYaml.Ledger.State.StateDatabase, "goleveldb")
	require.Equal(t, coreYaml.Ledger.State.TotalQueryLimit, 100000)
	require.Equal(t, coreYaml.Ledger.State.CouchDBConfig.CouchDBAddress, "127.0.0.1:5984")
	require.Empty(t, coreYaml.Ledger.State.CouchDBConfig.Username)
	require.Empty(t, coreYaml.Ledger.State.CouchDBConfig.Password)
	require.Equal(t, coreYaml.Ledger.State.CouchDBConfig.MaxRetries, 3)
	require.Equal(t, coreYaml.Ledger.State.CouchDBConfig.MaxRetriesOnStartup, 10)
	require.Equal(t, coreYaml.Ledger.State.CouchDBConfig.RequestTimeout, dur("35s"))
	require.Equal(t, coreYaml.Ledger.State.CouchDBConfig.InternalQueryLimit, 1000)
	require.Equal(t, coreYaml.Ledger.State.CouchDBConfig.MaxBatchUpdateSize, 1000)
	require.Equal(t, coreYaml.Ledger.State.CouchDBConfig.WarmIndexesAfterNBlocks, 1)
	require.False(t, coreYaml.Ledger.State.CouchDBConfig.CreateGlobalChangesDB)
	require.Equal(t, coreYaml.Ledger.State.CouchDBConfig.CacheSize, 64)
	require.True(t, coreYaml.Ledger.History.EnableHistoryDatabase)
	require.Equal(t, coreYaml.Ledger.PvtdataStore.CollElgProcMaxDbBatchSize, 5000)
	require.Equal(t, coreYaml.Ledger.PvtdataStore.CollElgProcDbBatchesInterval, 1000)
	require.Equal(t, coreYaml.Ledger.PvtdataStore.DeprioritizedDataReconcilerInterval, dur("60m"))
	require.Equal(t, coreYaml.Ledger.Snapshots.RootDir, "/var/hyperledger/production/snapshots")

	require.Equal(t, coreYaml.Operations.ListenAddress, "127.0.0.1:9443")
	require.False(t, coreYaml.Operations.Tls.Enabled)
	require.Empty(t, coreYaml.Operations.Tls.Cert.File)
	require.Empty(t, coreYaml.Operations.Tls.Key.File)
	require.False(t, coreYaml.Operations.Tls.ClientAuthRequired)
	require.Empty(t, coreYaml.Operations.Tls.ClientRootCAs.Files)

	require.Equal(t, coreYaml.Metrics.Provider, "disabled")
	require.Equal(t, coreYaml.Metrics.Statsd.Network, "udp")
	require.Equal(t, coreYaml.Metrics.Statsd.Address, "127.0.0.1:8125")
	require.Equal(t, coreYaml.Metrics.Statsd.WriteInterval, dur("10s"))
	require.Empty(t, coreYaml.Metrics.Statsd.Prefix)
}

func TestWriteCoreYaml (t *testing.T) {
	coreYaml, err := NewCoreYaml("core.yaml")
	if err != nil {
		t.Errorf("Failed to create CoreYaml: %v", err)
	}
	coreYaml.Write("test_core.yaml", 0755)

	coreYaml, err = NewCoreYaml("test_core.yaml")
	if err != nil {
		t.Errorf("Failed to load generated CoreYaml: %v", err)
	}

	t.Cleanup(func () {
		os.Remove("test_core.yaml")
	})
}
