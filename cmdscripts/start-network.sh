#!/bin/bash

source util/scriptUtils.sh

# remove.. adding here for convenience during testing
source clean-artifacts.sh "testNetwork"
source ../scripts/build.sh

# run gcryptogen main script to pull from config
../bin/gcryptogen

# create consortium and genesis block
../bin/gconfigtxgen

#../bin/gnetwork
../bin/gnetwork
