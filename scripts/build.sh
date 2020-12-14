#!/bin/bash

rm -R ../bin/*
go build -o ../bin/gcryptogen ../cmd/cryptogen/main.go
res=$?
if [ $res -ne 0 ]; then
  echo "Failed to build gryptogen"
  exit 1
else
  echo "Built: gcryptogen"
fi


