#!/bin/bash

echo
echo "##############################################"
echo "#                                             "
echo "#             Creating Consortium             "
echo "#                                             "
echo "##############################################"
echo

source cmdscripts/util/scriptUtils.sh

# Once you create the organization crypto material, you need to create the
# genesis block of the orderer system channel. This block is required to bring
# up any orderer nodes and create any application channels.

# The configtxgen tool is used to create the genesis block. Configtxgen consumes a
# "configtx.yaml" file that contains the definitions for the sample network. The
# genesis block is defined using the "TwoOrgsOrdererGenesis" profile at the bottom
# of the file. This profile defines a sample consortium, "SampleConsortium",
# consisting of our two Peer Orgs. This consortium defines which organizations are
# recognized as members of the network. The peer and ordering organizations are defined
# in the "Profiles" section at the top of the file. As part of each organization
# profile, the file points to a the location of the MSP directory for each member.
# This MSP is used to create the channel MSP that defines the root of trust for
# each organization. In essence, the channel MSP allows the nodes and users to be
# recognized as network members. The file also specifies the anchor peers for each
# peer org. In future steps, this same file is used to create the channel creation
# transaction and the anchor peer updates.
#
#
# If you receive the following warning, it can be safely ignored:
#
# [bccsp] GetDefault -> WARN 001 Before using BCCSP, please call InitFactories(). Falling back to bootBCCSP.
#
# You can ignore the logs regarding intermediate certs, we are not using them in
# this crypto implementation.

# Generate orderer system channel genesis block.
function createConsortium() {

  which "$BINARY"
  if [ "$?" -ne 0 ]; then
    fatalln "configtxgen tool not found."
  fi

  infoln "Generating Orderer Genesis block"

  # Note: For some unknown reason (at least for now) the block file can't be
  # named orderer.genesis.block or the orderer will fail to launch!
  set -x
  $BINARY -profile "$PROFILE" -channelID "$CHANNEL_ID" -outputBlock "$OUTPUT" -configPath "$CONFIG"
  res=$?
  { set +x; } 2>/dev/null
  if [ $res -ne 0 ]; then
    fatalln "Failed to generate orderer genesis block..."
  fi

  infoln "Generate CCP files for Org1 and Org2"
  # pass the path along to the shell file so it can reference it's
  # needed files relatively
  "networks/${NETWORK}/organizations/ccp-generate.sh" "networks/${NETWORK}/organizations/"
  res=$?
  { set +x; } 2>/dev/null
  if [ $res -ne 0 ]; then
    fatalln "Failed to generate CCP files for Org1 and Org2..."
  fi
}

PARAMS=""

# COMMENT OUT DEFAULTS FOR NOW
#BINARY="configtxgen"
#PROFILE="TwoOrgsOrdererGenesis"
#CHANNEL_ID="system-channel"
#OUTPUT="../networks/${NETWORK}/configtxgen/system-genesis-block/genesis.block"
## contains the configtx.yaml that is needed for the configtxgen binary
#CONFIG_PATH="../networks/${NETWORK}/configtxgen/"

while (( "$#" )); do
  case "$1" in
  -n | --network)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
      NETWORK=$2
      shift 2
    else
      fataln "Error: Argument for $1 is missing"
    fi
    ;;
  -p | --profile)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
      PROFILE=$2
      shift 2
    else
      fataln "Error: Argument for $1 is missing"
    fi
    ;;
  -c | --config)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
      CONFIG=$2
      shift 2
    else
      fataln "Error: Argument for $1 is missing"
    fi
    ;;
  -ch | --channel-id)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
      CHANNEL_ID=$2
      shift 2
    else
      fataln "Error: Argument for $1 is missing"
    fi
    ;;
  -o | --output)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
      OUTPUT=$2
      shift 2
    else
      fataln "Error: Argument for $1 is missing"
    fi
    ;;
    -b | --binary)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
      BINARY=$2
      shift 2
    else
      fataln "Error: Argument for $1 is missing"
    fi
    ;;
  *)
    PARAMS="$PARAMS $1"
    shift
    ;;
  esac
done

# comment out the following check to let defaults be created
if [ -z "${CONFIG}" ]; then
  fatalln "No config specified. Exiting..."
fi

if [ -z "${NETWORK}" ]; then
  fataln "Network no specified. Exiting..."
fi

createConsortium