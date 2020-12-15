#!/bin/bash

# currently not much of a gate
source network-gate.sh
source scriptUtils.sh

# remove.. adding here for convenience during testing
source clean-artifacts.sh
source ../scripts/build.sh

# run gcryptogen main script to pull from config
../bin/gcryptogen

# create consortium and genesis block
source create-consortium.sh

../bin/gnetwork
