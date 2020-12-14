#!/bin/bash

# toDo: add directory and file checks

rm -R ../cryptogen/organizations/*
echo "Cleaned cryptogen artifacts..."

rm ../bin/*
echo "Cleaned bin"

mv ../orderer/README.md ../or.md
rm -R ../orderer/*
mv ../or.md ../orderer/README.md
echo "Cleaned orderer" 

rm -R ../configtxgen/system-genesis-block
echo "Cleaned system-genesis-block"
