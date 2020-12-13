# Hyperman

### Hyperledger Command Center

Hyperman is a hyperledger command center that allows you to dynamically
spawn, create, destroy, and alter hyperledger nodes. It maintains a
memorybank of configurations and allows you to create, read, edit, and
delete, store, and version any configuration required to spawn a production ready network.

## Ideas
- create multiple network directories and have the configs and binaries
  isolated to each netdir
- dynamically create base configurations needed to start up a network

## Prerequisites

Hyperman needs access to the fabric and fabric-ca code directly so it
can dynamically build the cmd binaries.

For running locally, pull down `fabric` and `fabric-ca` and place the
following scrips in the respective paths.


1. add `buildCmd.sh` and `buildNodes.sh` to `fabric/scripts/`

2. add `fabric/scripts` and `/fabric/bin` to `$PATH`

## Guide

### Make Crypto material for organizations
To create the crypto material, you need organization configurations. See
`/organizations/*` for examples. 

1. under `/cryptogen/`, configure the `cryptogen.yaml` to point to
   configuration paths for organization nodes (peers and orderers). By
   default they point to the config files found under `/organizations`.
   Generate the crypto material by running `go run
   cmd/cryptogen/main.go` or just build the binary and call it.

### Create the Genesis Block

To create the Genesis Block we create a consortium. This is what the
`configtxgen` binary does for us. You can see the config flags this
commander takes under `fabric/cmd/configtxgen` on the first line of the
`main` function.

Other than needing a `profile` `channelID` and `outputBlock`, the
configtxgen will look for a `configtx.yaml` configuration file. You can
specify this path via `configPath` flag or it'll look at the
`FABRIC_CFG_PATH` env variable. So you can either pass in

1, `configtxgen ..... -configPath "../path/to/cfg"`
2. `FABRIC_CFG_PATH="../path/to/cfg" configtxgen ,,,,`
3. add `export FABRIC_CFG_PATH="../path/to/cfg"` to your bash profile
4. run `export FABRIC_CFG_PATH="../path/to/cfg"` in current shell



