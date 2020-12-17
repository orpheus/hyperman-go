#!/bin/bash

buildGcryptogen() {
  if [ -e "../bin/" ]; then
    rm ../bin/gcryptogen
    echo "Removed old binary before build."
  fi
  go build -o ../bin/gcryptogen ../cmd/cryptogen/main.go
  res=$?
  if [ $res -ne 0 ]; then
    echo "Failed to build gcryptogen"
    exit 1
  else
    echo "Built: gcryptogen"
  fi
}

buildGnetwork() {
  if [ -e "../bin/gnetwork" ]; then
    rm ../bin/gnetwork
    echo "Removed old binary before build."
  fi
  go build -o ../bin/gnetwork ../cmd/network/main.go
  res=$?
  if [ $res -ne 0 ]; then
    echo "Failed to build gnetwork"
    exit 1
  else
    echo "Built: gnetwork"
  fi
}

buildGconfigtxgen() {
  if [ -e "../bin/gconfigtxgen" ]; then
    rm ../bin/gconfigtxgen
    echo "Removed old binary before build."
  fi
  go build -o ../bin/gconfigtxgen ../cmd/configtxgen/main.go
  res=$?
  if [ $res -ne 0 ]; then
    echo "Failed to build gconfigtxgen"
    exit 1
  else
    echo "Built: gconfigtxgen"
  fi
}

# if cmd line arguments were passed, then build from them
if [ -n "$1" ]; then
  while (("$#")); do
    case "$1" in
    configtxgen|ctg)
    buildGcryptogen
    shift
    ;;
    network|n)
    buildGnetwork
    shift
    ;;
    cryptogen|cryp)
    buildGcryptogen
    shift
    ;;
    esac
  done
  exit
fi

buildGnetwork
buildGcryptogen
buildGconfigtxgen
