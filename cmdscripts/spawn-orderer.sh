#!/bin/bash

ORDERER_PATH="../cryptogen/organizations/testOrgs/ordererOrganizations/example.com/orderers/orderer.example.com"
GENESIS_BLOCK_PATH="../configtxgen/system-genesis-block/genesis.block"


BINARY="-b orderer"
ENV="-e FABRIC_LOGGING_SPEC=INFO\
    -e ORDERER_GENERAL_LISTENADDRESS=0.0.0.0\
    -e ORDERER_GENERAL_LISTENPORT=7050\
    -e ORDERER_GENERAL_GENESISMETHOD=file\
    -e ORDERER_GENERAL_GENESISFILE=${GENESIS_BLOCK_PATH}\
    -e ORDERER_GENERAL_LOCALMSPID=OrdererMSP\
    -e ORDERER_GENERAL_LOCALMSPDIR=${ORDERER_PATH}/msp\
    -e ORDERER_GENERAL_TLS_ENABLED=true\
    -e ORDERER_GENERAL_TLS_PRIVATEKEY=${ORDERER_PATH}/tls/server.key\
    -e ORDERER_GENERAL_TLS_CERTIFICATE=${ORDERER_PATH}/tls/server.crt\
    -e ORDERER_GENERAL_TLS_ROOTCAS=[${ORDERER_PATH}/tls/ca.crt]\
    -e ORDERER_KAFKA_TOPIC_REPLICATIONFACTOR=1\
    -e ORDERER_KAFKA_VERBOSE=true\
    -e ORDERER_GENERAL_CLUSTER_CLIENTCERTIFICATE=${ORDERER_PATH}/tls/server.crt\
    -e ORDERER_GENERAL_CLUSTER_CLIENTPRIVATEKEY=${ORDERER_PATH}/tls/server.key\
    -e ORDERER_GENERAL_CLUSTER_ROOTCAS=[${ORDERER_PATH}/tls/ca.crt]"


ARGS="${BINARY} ${ENV}"

echo "Spawn orderer..."

set -x
./spawn-node.sh $ARGS
{ set +x; } 
