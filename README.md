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

## Command Scripts Path Problem
- currently my paths are all over the place just to get things to
  work...
- `createCortium.sh` sets the working directory so I can run it
  anywhere, and `organizations/cpp-generate.sh` does the same thing so
  that it's path's can correctly reference the cpp-templates 
- I need to make it so I can run these commands anywhere without having
  to hack the `pwd` 
- in `cmd/cryptogen/main.go` I use an absolute path on my own directory
  structure to get it to correctly reference my config files and their
  relative paths... this needs to change as well

## Prerequisites

Hyperman needs access to the fabric and fabric-ca code directly so it
can dynamically build the cmd binaries.

For running locally, pull down `fabric` and `fabric-ca` and place the
following scrips in the respective paths.


1. add `buildCmd.sh` and `buildNodes.sh` to `fabric/scripts/`

2. add `fabric/scripts` and `/fabric/bin` to `$PATH`

3. run `buildCmd.sh -t=ALL`

**!: ** currently for these build scripts to work, you need to call them
in the `fabric` root dir... toDo: add env vars

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

**!:** the `configtx.yaml` file has to have a Profile that matches the
`-profile` you passed in to the `createConsortium.sh` script

**!:** then inside the `configtx.yaml` you need to point the
organization's msp directories to the correct cryptogen path which for
hyperman is currently under `/cryptogen/organization/{my_orgs}`

**!:** make sure that the orgs and orderers found in the `configtx`
match the number of `orgs` in peers and orderers found under
`/cryptogen/organizations/{my_orgs}/**

> CURRENTLY THE SYSTEM GENESIS BLOCK IS GENERATED RELATIVE TO WHERE YOU
CALL THE BASH SCRIPT SO CALL IT IN THE /CMDSCRIPTS/ DIR SO IT GETS
OUTPUT IN THE CORRECT PLACE


