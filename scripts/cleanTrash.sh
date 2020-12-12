#!/bin/bash

source scriptUtils.sh

rm -R ~/code/github/orpheus/go/hyperman-go/trash/*

res=$?
if [ $res -ne 0 ]; then
  fatalln "Faield to clean trash..."
fi

println "Trash Cleaned"

