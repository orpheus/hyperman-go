#!/bin/bash

source cmdscripts/util/scriptUtils.sh

buildHyperspace () {
  if [ -a "./bin/hyperspace" ]; then
    infoln "Found existing hyperspace binary, cleaning..."
    rm ./bin/hyperspace
  fi
  go build -o ./bin/hyperspace ./cmd/hyperspace/main.go
  infoln "Built hyperspace binary"
}

buildHyperspace