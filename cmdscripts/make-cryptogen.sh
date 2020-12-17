#!/bin/bash

echo
echo "##############################################"
echo "#                                             "
echo "#        GENERATING CRYPTO IDENTITY           "
echo "#                                             "
echo "##############################################"
echo

# Before you can bring up a network, each organization needs to generate the crypto
# material that will define that organization on the network. Because Hyperledger
# Fabric is a permissioned blockchain, each node and user on the network needs to
# use certificates and keys to sign and verify its actions. In addition, each user
# needs to belong to an organization that is recognized as a member of the network.
# You can use the Cryptogen tool or Fabric CAs to generate the organization crypto
# material.

# By default, the sample network uses cryptogen. Cryptogen is a tool that is
# meant for development and testing that can quickly create the certificates and keys
# that can be consumed by a Fabric network. The cryptogen tool consumes a series
# of configuration files for each organization in the "organizations/cryptogen"
# directory. Cryptogen uses the files to generate the crypto  material for each
# org in the "organizations" directory.

# You can also Fabric CAs to generate the crypto material. CAs sign the certificates
# and keys that they generate to create a valid root of trust for each organization.
# The script uses Docker Compose to bring up three CAs, one for each peer organization
# and the ordering organization. The configuration file for creating the Fabric CA
# servers are in the "organizations/fabric-ca" directory. Within the same directory,
# the "registerEnroll.sh" script uses the Fabric CA client to create the identities,
# certificates, and MSP folders that are needed to create the test network in the
# "organizations/ordererOrganizations" directory.

source util/scriptUtils.sh

checkCryptogen () {
  infoln "Checking for cryptogen binary..."
  which $1
  if [ "$?" -ne 0 ]; then
    if [ -e "../bin/$1" ]; then
      echo "Found $1"
    else
      fatalln "Cryptogen tool not found: $1... exiting..."
    fi
  fi
}

makeCrypto () {
  if [ -z $BINARY ]; then
    BINARY="cryptogen"
  fi

  checkCryptogen $BINARY

  if [ -z $IDENTITY ]; then
    IDENTITY="cryptogen"
  fi

  echo "CONFIG=$CONFIG" &&
  echo "OUTPUT=$OUTPUT" &&
  echo "IDENTITY=$IDENTITY" &&
  echo "BINARY=$BINARY" &&

  cmd="$BINARY generate"

  if [ -n "$CONFIG" ]; then
    cmd="$cmd --config=$CONFIG"
  fi

  if [ -n "$OUTPUT" ]; then
    cmd="$cmd --output=$OUTPUT"
  fi

  infoln "Creating crypto for $IDENTITY"
  set -x
  $cmd
  res=$?
  { set +x; } 2>/dev/null
  if [ $res -ne 0 ]; then
    fatalln "Failed to generate certificates..."
  fi

  infoln "Created crypto for $IDENTITY\n"
}


PARAMS=""

while (( "$#" )); do
  case "$1" in
    -n|--network)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
      NETWORK=$2
      shift 2
    else
      echo "Error: Argument for $1 is missing" >&2
      exit 1
    fi
    ;;
    -c|--config)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
      CONFIG=$2
      shift 2
    else
      echo "Error: Argument for $1 is missing" >&2
      exit 1
    fi
    ;;
    -o|--output)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
      OUTPUT=$2
      shift 2
    else
      echo "Error: Argument for $1 is missing" >&2
      exit 1
    fi
    ;;
    -b|--binary)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
      BINARY=$2
      shift 2
    else
      echo "Error: Argument for $1 is missing" >&2
      exit 1
    fi
    ;;
    -i|--identity)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
      IDENTITY=$2
      shift 2
    else
      echo "Error: Argument for $1 is missing" >&2
      exit 1
    fi
    ;;
    *) # preserve positional arguments
    PARAMS="$PARAMS $1"
    shift
    ;;
esac
done

# set positional arguments in their proper place
eval set -- "${PARAMS}"

# comment out the following check to let defaults be created
if [ -z "${CONFIG}" ]; then
  fatalln "No config specified. Exiting..."
fi

if [ -z "${NETWORK}" ]; then
  fataln "Network no specified. Exiting..."
fi

makeCrypto

