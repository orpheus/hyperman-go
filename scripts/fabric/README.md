## Scripts to build Fabric binaries

copy and past these into FABRIC_PROJECT_PATH/scripts and use them to generate binaries dynamically

## buildNodes.sh
Build peers dynamically naming them by an index
```sh
# Build two peers and one orderer
./buildNodes -p=2 -o=1

## outputs a `peer-0` `peer-1` and `orderer-1` binary
```

## buildCmd.sh

Builds individual peers with customer name or all peers if `ALL` is passed to `-t`
```shell script
# build all fabric binaries
./buildCmd.sh -t=ALL

# build fabric binary with specified name
./buildCmd.sh  -t=peer -n=my-peer
```
