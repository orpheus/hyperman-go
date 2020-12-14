#!/bin/bash

./clean-artifacts.sh

../scripts/build.sh

../bin/gcryptogen
 
./create-consortium.sh
 
./spawn-orderer.sh
 

