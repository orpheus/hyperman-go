#!/bin/bash

# script to build a dynamic amount of hypderledger cmd binaries
# must have /buildCmd.sh/ set to $PATH
#
# To build 3 peers, 1 orderer, and 1 cryptogen
# buildNodes.sh -p=3 -o=1 -c=1 
#
#

for i in "$@"
do
    case $i in
        -ctg=*|--configtxgen=*)
        CONFIGTXGEN=configtxgen
        CONFIGTXGEN_AMT="${i#*=}"
        shift
        ;;
        -ctl=*|--configtxlator=*)
        CONFIGTXLATOR=configtxlator
        CONFIGTXLATOR_AMT="${i#*=}"
        shift;
        ;;
        -c=*|--cryptogen=*)
        CRYPTOGEN=cryptogen
        CRYPTOGEN_AMT="${i#*=}"
        shift;
        ;;
        -d=*|--discover=*)
        DISCOVER=discover
        DISCOVER_AMT="${i#*=}"
        shift;
        ;;
        -i=*|--idemixgen=*)
        IDEMIXGEN=idemixgen
        IDEMIXGEN_AMT="${i#*=}"
        shift;
        ;;
        -o=*|--orderer=*)
        ORDERER=orderer
        ORDERER_AMT="${i#*=}"
        shift
        ;;
        -osn=*|-a=*|--admin|--osnadmin=*)
        OSNADMIN=osnadmin
        OSNADMIN_AMT="${i#*=}"
        shift
        ;;
        -p=*|--peer=*)
        PEER=peer
        PEER_AMT="${i#*=}"
        shift
        ;;
        --default)
        DEFAULT=YES
        shift
        ;;
        
     *)

    ;;
esac
done

echo "##############################################"
echo "#                                            #"
echo "#        BUILDING HYPERLEDGER BINARIES       #"
echo "#                                            #"
echo "##############################################"
echo "Building the following:"
echo
echo "CONFIGTXGEN:${CONFIGTXGEN}:${CONFIGTXGEN_AMT}"
echo "CONFIGTXLATOR:${CONFIGTXLATOR}:${CONFIGTXLATOR_AMT}"
echo "CRYPTOGEN:${CRYPTOGEN}:${CRYPTOGEN_AMT}"
echo "DISCOVER:${DISCOVER}:${DISCOVER_AMT}"
echo "IDEMIXGEN:${IDEMIXGEN}:${IDEMIXGEN_AMT}"
echo "ORDERER:${ORDERER}:${ORDERER_AMT}"
echo "OSNADMIN:${OSNAMIN}:${OSNADMIN_AMT}"
echo "PEER:${PEER}:$PEER_AMT"
echo

buildCmd () {
    buildCmd.sh --type=$1 --name=$2
}

build () {
    _type=$1
    _amt=$2
    
    if [ ! -z $_type ]; then
        echo "building: $_amt $_type"
        for (( i=0; i < $_amt; i++ )); do
            echo "$_type-$i"
            name="$_type-$i"
            buildCmd $_type "$_type-$i"
        done
    fi
}

if [ ! -z $CONFIGTXGEN ]; then
    build $CONFIGTXGEN $CONFIGTXGEN_AMT
fi

if [ ! -z $CONFIGTXLATOR ]; then
    build $CONFIGTXLATOR $CONFIGTXLATOR_AMT
fi

if [ ! -z $CRYPTOGEN ]; then
    build $CRYPTOGEN $CRYPTOGEN_AMT
fi

if [ ! -z $DISCOVER ]; then
    build $DISCOVER $DISCOVER_AMT
fi

if [ ! -z $IDEMIXGEN ]; then
    build $IDEMIXGEN $IDEMIXGEN_AMT
fi

if [ ! -z $ORDERER ]; then
    build $ORDERER $ORDERER_AMT
fi

if [ ! -z $OSNADMIN ]; then
    build $OSNADMIN $OSNADMIN_AMT
fi

if [ ! -z $PEER ]; then
    build $PEER $PEER_AMT
fi

