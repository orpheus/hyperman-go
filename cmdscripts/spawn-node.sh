#!/bin/bash

# Spawn hyperledger nodes given a binary, environment variables, and
# start script. This allows me to build multiple binaries of the same
# type with different names and spawn them all through a single script

BINARY=""
ENVIRONMENT_VARS=""
START_CMD=""
PARAMS=""

while (( "$#" )); do 
    case "$1" in
    -b|--binary)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
        BINARY=$2
        shift 2
    else 
        echo "Argument for $1 is missing" >&2
        exit 1
    fi
    ;;
    -e|-env|--env-var)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
        # split env key=value up by "="
        key="${2%%=*}" # get prefix
        value="${2#*=}" # get suffix
        if [ -n "${key}" ]; then
            if [ -n "${value}" ]; then
                echo "export ${2}"
                export "${2}"
            else
                echo "ERROR: missing env value for $key"
                echo "exiting..."
                exit 1
            fi
        else
            echo "Error: missing env var :${2}"
            echo "exiting..."
            exit 1
        fi
        shift 2
    else 
        echo "Argument for $1 is missing" >&2
        exit 1
    fi
    ;;
    -cmd|--start-cmd|--command|--start)
    if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
        START_CMD=$2
        shift 2
    else 
        echo "Argument for $1 is missing" >&2
        exit 1
    fi
    ;;
    *)
    PARANS="$PARAMS $1"
    shift
    ;;
esac
done

set -x
$BINARY $START_CMD
res=$?
{ set +x; } 2>/dev/null
if [ $res -ne 0 ]; then
    echo "error: failed to start binary"
    echo "exiting..."
    exit 1
fi

echo "${BINARY} started..."





    
