package orderer

import (
	"github.com/orpheus/hyperspace/util"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var testYaml = "orderer_test.yaml"

//----------------------------------------------------------------------------------
// TestLoadNew()
//----------------------------------------------------------------------------------
// Loads an orderer.yaml file into an OrdererYaml config struct and then checks
// each value of the struct to make sure the fields got set correctly according
// to the orderer.yaml file read in.
//----------------------------------------------------------------------------------
func TestLoadNew(t *testing.T) {
	o := NewOrdererYaml(testYaml)

	require.Equal(t, o.General.ListenAddress, "127.0.0.1")
	require.Equal(t, o.General.ListenPort, uint16(7050))
	require.False(t, o.General.TLS.Enabled)
	require.Equal(t, o.General.TLS.PrivateKey, "tls/server.key")
	require.Equal(t, o.General.TLS.Certificate, "tls/server.crt")
	require.Equal(t, o.General.TLS.RootCAs[0], "tls/ca.crt")
	require.False(t, o.General.TLS.ClientAuthRequired)
	require.Empty(t, o.General.TLS.ClientRootCAs)
	require.Equal(t, o.General.Keepalive.ServerMinInterval, util.Dur("60s"))
	require.Equal(t, o.General.Keepalive.ServerInterval, util.Dur("7200s"))
	require.Equal(t, o.General.Keepalive.ServerTimeout, util.Dur("20s"))
	require.Equal(t, o.General.Cluster.SendBufferSize, 10)
	require.Empty(t, o.General.Cluster.ClientCertificate)
	require.Empty(t, o.General.Cluster.ClientPrivateKey)
	require.Empty(t, o.General.Cluster.ListenPort)
	require.Empty(t, o.General.Cluster.ListenAddress)
	require.Empty(t, o.General.Cluster.ServerCertificate)
	require.Empty(t, o.General.Cluster.ServerPrivateKey)
	require.Equal(t, o.General.BootstrapMethod, "file")
	require.Empty(t, o.General.BootstrapFile)
	require.Equal(t, o.General.LocalMSPDir, "msp")
	require.Equal(t, o.General.LocalMSPID, "SampleOrg")
	require.False(t, o.General.Profile.Enabled)
	require.Equal(t, o.General.Profile.Address, "0.0.0.0:6060")
	require.Equal(t, o.General.BCCSP.Default, "SW")
	require.Equal(t, o.General.BCCSP.SW.Hash, "SHA2")
	require.Equal(t, o.General.BCCSP.SW.Security, 256)
	require.Empty(t, o.General.BCCSP.SW.FileKeystore.KeyStorePath)
	require.Empty(t, o.General.BCCSP.PKCS11.Library)
	require.Empty(t, o.General.BCCSP.PKCS11.Label)
	require.Empty(t, o.General.BCCSP.PKCS11.Pin)
	require.Empty(t, o.General.BCCSP.PKCS11.Hash)
	require.Empty(t, o.General.BCCSP.PKCS11.Security)
	require.Empty(t, o.General.BCCSP.PKCS11.FileKeystore.KeyStorePath)
	require.Equal(t, o.General.Authentication.TimeWindow, util.Dur("15m"))

	require.Equal(t, o.FileLedger.Location, "/var/hyperledger/production/orderer")

	require.Equal(t, o.Kafka.Retry.ShortInterval, util.Dur("5s"))
	require.Equal(t, o.Kafka.Retry.ShortTotal, util.Dur("10m"))
	require.Equal(t, o.Kafka.Retry.LongInterval, util.Dur("5m"))
	require.Equal(t, o.Kafka.Retry.LongTotal, util.Dur("12h"))
	require.Equal(t, o.Kafka.Retry.NetworkTimeouts.DialTimeout, util.Dur("10s"))
	require.Equal(t, o.Kafka.Retry.NetworkTimeouts.ReadTimeout, util.Dur("10s"))
	require.Equal(t, o.Kafka.Retry.NetworkTimeouts.WriteTimeout, util.Dur("10s"))
	require.Equal(t, o.Kafka.Retry.Metadata.RetryBackoff, util.Dur("250ms"))
	require.Equal(t, o.Kafka.Retry.Metadata.RetryMax, 3)
	require.Equal(t, o.Kafka.Retry.Producer.RetryBackoff, util.Dur("100ms"))
	require.Equal(t, o.Kafka.Retry.Producer.RetryMax, 3)
	require.Equal(t, o.Kafka.Retry.Consumer.RetryBackoff, util.Dur("2s"))
	require.Equal(t, o.Kafka.Topic.ReplicationFactor, int16(3))
	require.False(t, o.Kafka.Verbose)
	require.False(t, o.Kafka.TLS.Enabled)
	require.Empty(t, o.Kafka.TLS.PrivateKey)
	require.Empty(t, o.Kafka.TLS.Certificate)
	require.Empty(t, o.Kafka.TLS.RootCAs)
	require.False(t, o.Kafka.SASLPlain.Enabled)
	require.Empty(t, o.Kafka.SASLPlain.User)
	require.Empty(t, o.Kafka.SASLPlain.Password)
	require.Empty(t, o.Kafka.Version)

	require.Empty(t, o.Debug.BroadcastTraceDir)
	require.Empty(t, o.Debug.DeliverTraceDir)

	require.Equal(t, o.Operations.ListenAddress, "127.0.0.1:8443")
	require.False(t, o.Operations.TLS.Enabled)
	require.Empty(t, o.Operations.TLS.Certificate)
	require.Empty(t, o.Operations.TLS.PrivateKey)
	require.False(t, o.Operations.TLS.ClientAuthRequired)
	require.Empty(t, o.Operations.TLS.ClientRootCAs)

	require.Equal(t, o.Metrics.Provider, "disabled")
	require.Equal(t, o.Metrics.Statsd.Network, "udp")
	require.Equal(t, o.Metrics.Statsd.Address, "127.0.0.1:8125")
	require.Equal(t, o.Metrics.Statsd.WriteInterval, util.Dur("30s"))
	require.Empty(t, o.Metrics.Statsd.Prefix)

	require.Equal(t, o.Admin.ListenAddress, "127.0.0.1:9443")
	require.False(t, o.Admin.TLS.Enabled)
	require.Empty(t, o.Admin.TLS.Certificate)
	require.Empty(t, o.Admin.TLS.PrivateKey)
	require.True(t, o.Admin.TLS.ClientAuthRequired)
	require.Empty(t, o.Admin.TLS.ClientRootCAs)

	require.False(t, o.ChannelParticipation.Enabled)
	require.Equal(t, o.ChannelParticipation.MaxRequestBodySize, "1 MB")

	require.Equal(t, o.Consensus.WALDir, "/var/hyperledger/production/orderer/etcdraft/wal")
	require.Equal(t, o.Consensus.SnapDir, "/var/hyperledger/production/orderer/etcdraft/snapshot")
}

//----------------------------------------------------------------------------------
// TestWrite()
//----------------------------------------------------------------------------------
// Reads in an orderer.yaml, writes it the file system under another name, creating
// a new orderer.yaml, then reads the generated yaml in and expects no errors.
//----------------------------------------------------------------------------------
func TestWrite(t *testing.T) {
	generated := "test_orderer.yaml"

	OrdererYaml := NewOrdererYaml(testYaml)
	OrdererYaml.Write(generated, 0755)

	OrdererYaml = NewOrdererYaml(generated)

	t.Cleanup(func() {
		_ = os.Remove(generated)
	})
}

//----------------------------------------------------------------------------------
// TestMerge()
//----------------------------------------------------------------------------------
// Reads in two different instances of orderer.yaml, changes some of the values of
// the second one read in, then merges the two, expected the changed values to be
// present in the merged config struct.
//----------------------------------------------------------------------------------
func TestMerge(t *testing.T) {
	baseOrdererYaml := NewOrdererYaml(testYaml)

	// read in another instance of a Orderer config
	// simulates reading in a user's Orderer config
	userOrdererYaml := NewOrdererYaml(testYaml)

	// change some values
	userOrdererYaml.General.ListenAddress = "999.9.9.9"
	userOrdererYaml.General.ListenPort = 0
	userOrdererYaml.Metrics.Statsd.Network = "tcp"

	baseOrdererYaml.Merge(userOrdererYaml)

	// expect the changes values to exist on the original struct
	require.Equal(t, baseOrdererYaml.General.ListenAddress, "999.9.9.9")
	require.Equal(t, baseOrdererYaml.General.ListenPort, uint16(0))
	require.Equal(t, baseOrdererYaml.Metrics.Statsd.Network, "tcp")
}
