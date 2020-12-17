#!/bin/bash

###################################################
# HOW DO YOU CLEAN FILES UNDER DIRECTORIES WITH * #
###################################################

#source network-gate.sh
NETWORK="testNetwork"

# Clean binaries
if [ -a "../bin/gcryptogen" ]; then
    rm ../bin/*
    res=$?
    if [ $res -ne 0 ]; then
        echo "Failed to clean bin"
        exit 1
    else
        echo "Cleaned bin"
    fi
else
    echo "Bin clean."
fi

# Clean cryptogen material for organizations
if [ -d "../networks/${NETWORK}/cryptogen/organizations" ]; then
    # same problem here as with trying to remove the ledger/* data, see below...
    rm -R "../networks/${NETWORK}/cryptogen/organizations"
    res=$?
    if [ $res -ne 0 ]; then
        echo "Failed to remove cryptogen"
        exit 1
    else
        mkdir "../networks/${NETWORK}/cryptogen/organizations"
        echo "Cleaned cryptogen artifacts..."
    fi
else
    echo "Cryptogen clean."
fi

# Clean orderer material (this is the ledger artifacts the get created)
# should this clean the ledger for each orderer? probably
# for now just static hard code it to orderer-01 for testNetwork
# mv ../orderer/README.md ../or.md

# !important, if there are no files under ledger/, then ledger/* will
# return an error because it cannot find any files to delete
if [ -d "../networks/${NETWORK}/nodes/orderers/orderer-01/ledger" ]; then
    # hack.. idk how to rm ledger data with path/ledger/*
    # when I sue the /* bash throws an error saying there's no file
    # or directory even though I can run the same cmd from the terminal
    # and it deletes fine - so just remove the ledger folder for now
    # and make it after deletion
      rm -r "../networks/${NETWORK}/nodes/orderers/orderer-01/ledger"
    res=$?
    if [ $res -ne 0 ]; then
        echo "Failed to remove ledger data for orderer-01"
        exit 1
    else
        mkdir "../networks/${NETWORK}/nodes/orderers/orderer-01/ledger"
        echo "Cleaned ledger data for orderer-01"
    fi
else
    echo "Ledger data for orderer-01 clean"
fi

# mv ../or.md ../orderer/README.md

# Clean system-genesis-block
if [ -d "../networks/${NETWORK}/configtxgen/system-genesis-block" ]; then
    rm -R "../networks/${NETWORK}/configtxgen/system-genesis-block"
    res=$?
    if [ $res -ne 0 ]; then
        echo "Failed to clean system-genesis-block"
        exit 1
    else
        echo "Cleaned system-genesis-block"
    fi
else
    echo "System genesis block clean."
fi
