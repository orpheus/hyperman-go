#!/bin/bash

# Clean cryptogen material for organizations
if [ -d "../cryptogen/organizations/*" ]; then
    rm -R ../cryptogen/organizations/*
    res=$?
    if [ $res -ne 0 ]; then
        echo "Failed to remove cryptogen"
        exit 1
    else
        echo "Cleaned cryptogen artifacts..."
    fi
else
    echo "Cryptogen clean."
fi

# Clean binaries
if [ -d "../bin/*" ]; then
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

# Clean orderer material
mv ../orderer/README.md ../or.md

if [ -d "../orderer/*" ]; then
    rm -R ../orderer/*
    res=$?
    if [ $res -ne 0 ]; then
        echo "Failed to remove orderer/*"
        exit 1
    else
        echo "Removed orderer/*"
    fi
else
    echo "Orderer clean."
fi

mv ../or.md ../orderer/README.md

# Clean system-genesis-block
if [ -d "../configtxgen/system-genesis-block" ]; then
    rm -R ../configtxgen/system-genesis-block
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
