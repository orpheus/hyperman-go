#!/bin/bash

if [ -z "$1" ]; then
  echo "error: need to specify network as first argument"
  exit 1
fi

NETWORK="${1}"

if [ -d "../networks/${NETWORK}" ]; then
  echo "Found network: ${NETWORK}."
else
  echo "Could not find network: ${NETWORK}."
  exit 1
fi
