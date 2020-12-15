#!/bin/bash

if [ -e "../bin/gcryptogen" ]; then
  rm ../bin/gcryptogen
  echo "Removed old binary before build."
fi
go build -o ../bin/gcryptogen ../cmd/cryptogen/main.go
res=$?
if [ $res -ne 0 ]; then
  echo "Failed to build gryptogen"
  exit 1
else
  echo "Built: gcryptogen"
fi

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


