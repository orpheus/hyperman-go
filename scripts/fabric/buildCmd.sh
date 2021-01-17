#!/bin/bash

# accepts a type (-t=type) and name (-n=name) where type is the fabric binary name
# ./buildCmd.sh -t=ALL // to build all nodes
# ./buildCmd.sh -t=peer -n=my-peer

for i in "$@"
do
    case $i in
        -t=*|--type=*)
        TYPE="${i#*=}"
        shift
        ;;
        -n=*|--name=*)
        NAME="${i#*=}"
        shift
        ;;
        --default)
        DEFAULT=TRUE
        shift
        ;;
    *)
        ;;
esac
done

if [ -z $TYPE ]; then
    echo "No type specified. Building peer"
    TYPE="peer"
fi

if [ -z $NAME ]; then
    NAME=$TYPE
fi

function logBuilding() {
    echo
    echo "building: type=$TYPE:name=$NAME"
}

function logBuilt() {
    echo "built: type=$TYPE:name=$NAME"
    echo
}

if [ $TYPE == "ALL" ]; then
    echo "Building ALL binaries..."
    go build -o ./bin ./cmd/...
    echo "Built ALL binaries."
    exit;
# osnadmin errors when trying to build like the others
# it needs the main.go specified in the path
# else you will get an error like:
# "cannot write multiple packages to non-directory"
elif [ $TYPE == "osnadmin" ]; then
    logBuilding
    go build -o ./bin/$NAME ./cmd/$TYPE/main.go
    logBuilt
    exit
fi

logBuilding

go build -o ./bin/$NAME ./cmd/$TYPE/...
 
logBuilt
