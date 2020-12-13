#!/bin/bash

rm -R ./bin/*
go build -o bin/gcryptogen ./cmd/cryptogen/main.go
echo "Built: gcryptogen"

