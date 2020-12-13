#!/bin/bash

echo 
echo "##############################################"
echo "#                                            #"
echo "#        GENERATING CRYPTO IDENTITIES        #"
echo "#                                            #"
echo "##############################################"
echo

pwd

source scriptUtils.sh

checkCryptogen () {
  infoln "Checking for cryptogen binary..."
  which $1
  if [ "$?" -ne 0 ]; then
    fatalln "Cryptogen tool not found: $1... exiting..."
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
      echo "FOUND identity $2"
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
eval set -- $PARAMS

# remove the following check to let defaults be created
if [ -z $CONFIG ]; then
  fatalln "No config specified. Exiting..."
fi

makeCrypto
