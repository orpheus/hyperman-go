#!/bin/bash

source scriptUtils.sh

checkCryptogen () {
  echo "Checking for cryptogen binary"
  which $1
  if [ "$?" -ne 0 ]; then
    fatalln "Cryptogen tool not found: $1... exiting..."
  fi
}

makeCrypto () {
  if [ -z $BINARY ]; then
    echo "Overwriting binary: $BINARY"
    BINARY="cryptogen"
  fi

  checkCryptogen $BINARY

  if [ -z $IDENTIY ]; then
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
  # $BINARY generate --config=$CONFIG --output=$OUTPUT
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

makeCrypto

