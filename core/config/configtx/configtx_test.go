package config

import (
	"github.com/orpheus/hyperspace/core/util"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

//----------------------------------------------------------------------------------
// TestLoadNew()
//----------------------------------------------------------------------------------
// Loads an configtx.yaml file into an ConfigtxYaml config struct and then checks
// each value of the struct to make sure the fields got set correctly according
// to the configtx.yaml file read in.
//----------------------------------------------------------------------------------
func TestLoadNew(t *testing.T) {
	c := NewConfigtxYaml("configtx.yaml")

	// -------------------------------------------------------------
	// PROFILE 1, "TwoOrgsOrdererGenesis"
	// -------------------------------------------------------------
	profile1 := c.Profiles["TwoOrgsOrdererGenesis"]

	// test *ChannelDefaults
	require.Equal(t, profile1.Policies["Readers"].Type, "ImplicitMeta")
	require.Equal(t, profile1.Policies["Readers"].Rule, "ANY Readers")
	require.Equal(t, profile1.Policies["Writers"].Type, "ImplicitMeta")
	require.Equal(t, profile1.Policies["Writers"].Rule, "ANY Writers")
	require.Equal(t, profile1.Policies["Admins"].Type, "ImplicitMeta")
	require.Equal(t, profile1.Policies["Admins"].Rule, "MAJORITY Admins")
	require.True(t, profile1.Capabilities["V2_0"])
	// test Orderer/*OrdererDefaults
	require.Equal(t, profile1.Orderer.OrdererType, "etcdraft")
	require.Equal(t, profile1.Orderer.Addresses[0], "orderer.example.com:7050")
	require.Equal(t, profile1.Orderer.EtcdRaft.Consenters[0].Host, "orderer.example.com")
	require.Equal(t, profile1.Orderer.EtcdRaft.Consenters[0].Port, uint32(7050))
	require.Equal(t, profile1.Orderer.EtcdRaft.Consenters[0].ClientTlsCert, "../organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.crt")
	require.Equal(t, profile1.Orderer.EtcdRaft.Consenters[0].ServerTlsCert, "../organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.crt")
	require.Equal(t, profile1.Orderer.BatchTimeout, util.Dur("2s"))
	require.Equal(t, profile1.Orderer.BatchSize.MaxMessageCount, "10")
	require.Equal(t, profile1.Orderer.BatchSize.AbsoluteMaxBytes, "99 MB")
	require.Equal(t, profile1.Orderer.BatchSize.PreferredMaxBytes, "512 KB")
	// test Orderer/Organizations[*OrdererOrg]
	require.Equal(t, profile1.Orderer.Organizations[0].Name, "OrdererOrg")
	require.Equal(t, profile1.Orderer.Organizations[0].ID, "OrdererMSP")
	require.Equal(t, profile1.Orderer.Organizations[0].MSPDir, "../organizations/ordererOrganizations/example.com/msp")
	require.Equal(t, profile1.Orderer.Organizations[0].Policies["Readers"].Type, "Signature")
	require.Equal(t, profile1.Orderer.Organizations[0].Policies["Readers"].Rule, "OR('OrdererMSP.member')")
	require.Equal(t, profile1.Orderer.Organizations[0].Policies["Writers"].Type, "Signature")
	require.Equal(t, profile1.Orderer.Organizations[0].Policies["Writers"].Rule, "OR('OrdererMSP.member')")
	require.Equal(t, profile1.Orderer.Organizations[0].Policies["Admins"].Type, "Signature")
	require.Equal(t, profile1.Orderer.Organizations[0].Policies["Admins"].Rule, "OR('OrdererMSP.admin')")
	require.Equal(t, profile1.Orderer.Organizations[0].OrdererEndpoints[0], "orderer.example.com:7050")
	// test Orderer/Capabilities/*OrdererCapabilities
	require.True(t, profile1.Orderer.Capabilities["V2_0"])
	consort1 := profile1.Consortiums["SampleConsortium"]
	// test only Org1, assume if that passes, Org2 would pass as it has the same info
	require.Equal(t, consort1.Organizations[0].Name, "Org1MSP")
	require.Equal(t, consort1.Organizations[0].ID, "Org1MSP")
	require.Equal(t, consort1.Organizations[0].MSPDir, "../organizations/peerOrganizations/org1.example.com/msp")
	require.Equal(t, consort1.Organizations[0].Policies["Readers"].Type, "Signature")
	require.Equal(t, consort1.Organizations[0].Policies["Readers"].Rule, "OR('Org1MSP.admin', 'Org1MSP.peer', 'Org1MSP.client')")
	require.Equal(t, consort1.Organizations[0].Policies["Writers"].Type, "Signature")
	require.Equal(t, consort1.Organizations[0].Policies["Writers"].Rule, "OR('Org1MSP.admin', 'Org1MSP.client')")
	require.Equal(t, consort1.Organizations[0].Policies["Admins"].Type, "Signature")
	require.Equal(t, consort1.Organizations[0].Policies["Admins"].Rule, "OR('Org1MSP.admin')")
	require.Equal(t, consort1.Organizations[0].Policies["Endorsement"].Type, "Signature")
	require.Equal(t, consort1.Organizations[0].Policies["Endorsement"].Rule, "OR('Org1MSP.peer')")
	require.Equal(t, consort1.Organizations[0].AnchorPeers[0].Host, "peer0.org1.example.com")
	require.Equal(t, consort1.Organizations[0].AnchorPeers[0].Port, 7051)

	// -------------------------------------------------------------
	// PROFILE 2, "TwoOrgsChannel"
	// -------------------------------------------------------------
	profile2 := c.Profiles["TwoOrgsChannel"]
	require.Equal(t, profile2.Consortium, "SampleConsortium")
	// *ChannelDefaults
	require.True(t, profile2.Capabilities["V2_0"])
	require.Equal(t, profile2.Policies["Readers"].Type, "ImplicitMeta")
	require.Equal(t, profile2.Policies["Readers"].Rule, "ANY Readers")
	require.Equal(t, profile2.Policies["Writers"].Type, "ImplicitMeta")
	require.Equal(t, profile2.Policies["Writers"].Rule, "ANY Writers")
	require.Equal(t, profile2.Policies["Admins"].Type, "ImplicitMeta")
	require.Equal(t, profile2.Policies["Admins"].Rule, "MAJORITY Admins")
	// test Application/*ApplicationDefaults
	require.True(t, profile2.Application.Capabilities["V2_0"])
	require.Equal(t, profile2.Application.Policies["Readers"].Type, "ImplicitMeta")
	require.Equal(t, profile2.Application.Policies["Readers"].Rule, "ANY Readers")
	require.Equal(t, profile2.Application.Policies["Writers"].Type, "ImplicitMeta")
	require.Equal(t, profile2.Application.Policies["Writers"].Rule, "ANY Writers")
	require.Equal(t, profile2.Application.Policies["Admins"].Type, "ImplicitMeta")
	require.Equal(t, profile2.Application.Policies["Admins"].Rule, "MAJORITY Admins")
	require.Equal(t, profile2.Application.Policies["LifecycleEndorsement"].Type, "ImplicitMeta")
	require.Equal(t, profile2.Application.Policies["LifecycleEndorsement"].Rule, "MAJORITY Endorsement")
	require.Equal(t, profile2.Application.Policies["Endorsement"].Type, "ImplicitMeta")
	require.Equal(t, profile2.Application.Policies["Endorsement"].Rule, "MAJORITY Endorsement")
	// test Application/Organizations[*Org1]
	org1 := profile2.Application.Organizations[0]
	require.Equal(t, org1.Name, "Org1MSP")
	require.Equal(t, org1.ID, "Org1MSP")
	require.Equal(t, org1.MSPDir, "../organizations/peerOrganizations/org1.example.com/msp")
	require.Equal(t, org1.Policies["Readers"].Type, "Signature")
	require.Equal(t, org1.Policies["Readers"].Rule, "OR('Org1MSP.admin', 'Org1MSP.peer', 'Org1MSP.client')")
	require.Equal(t, org1.Policies["Writers"].Type, "Signature")
	require.Equal(t, org1.Policies["Writers"].Rule, "OR('Org1MSP.admin', 'Org1MSP.client')")
	require.Equal(t, org1.Policies["Admins"].Type, "Signature")
	require.Equal(t, org1.Policies["Admins"].Rule, "OR('Org1MSP.admin')")
	require.Equal(t, org1.Policies["Endorsement"].Type, "Signature")
	require.Equal(t, org1.Policies["Endorsement"].Rule, "OR('Org1MSP.peer')")
	require.Equal(t, org1.AnchorPeers[0].Host, "peer0.org1.example.com")
	require.Equal(t, org1.AnchorPeers[0].Port, 7051)
}

//----------------------------------------------------------------------------------
// TestReadAndWrite()
//----------------------------------------------------------------------------------
// Reads in an configtx.yaml, writes it the file system under another name, creating
// a new configtx.yaml, then reads the generated yaml in and expects no errors.
//----------------------------------------------------------------------------------
func TestReadAndWrite(t *testing.T) {
	generated := "configtx2.yaml"

	coreYaml := NewConfigtxYaml("configtx.yaml")

	coreYaml.Write(generated, 0755)

	coreYaml = NewConfigtxYaml(generated)

	t.Cleanup(func() {
		_ = os.Remove(generated)
	})
}

//----------------------------------------------------------------------------------
// TestMerge()
//----------------------------------------------------------------------------------
// Reads in two different instances of configtx.yaml, changes some of the values of
// the second one read in, then merges the two, expected the changed values to be
// present in the merged config struct.
//----------------------------------------------------------------------------------
func TestMerge(t *testing.T) {
	c := NewConfigtxYaml("configtx.yaml")

	// read in another instance of a core config
	// simulates reading in a user's core config
	c2 := NewConfigtxYaml("configtx.yaml")

	// change some values
	c2.Profiles["TwoOrgsChannel"].Consortium = "TestConsortium"
	c2.Profiles["TwoOrgsOrdererGenesis"].Consortiums["SampleConsortium"].Organizations[0].Name = "Test_Org1MSP"
	c2.Profiles["TwoOrgsOrdererGenesis"].Orderer.Addresses[0] = "test.example.org:9999"

	c.Merge(c2)

	// expect the changes values to exist on the original struct
	require.Equal(t, c.Profiles["TwoOrgsChannel"].Consortium, "TestConsortium")
	require.Equal(t, c.Profiles["TwoOrgsOrdererGenesis"].Consortiums["SampleConsortium"].Organizations[0].Name, "Test_Org1MSP")
	require.Equal(t, c.Profiles["TwoOrgsOrdererGenesis"].Orderer.Addresses[0], "test.example.org:9999")
}
