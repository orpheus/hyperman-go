# Hyperspace

## Go far, go fast, explore deep.

### Hyperledger Command Center

Hyperspace is a hyperledger command center that allows you to dynamically
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

Hyperspace needs access to the fabric and fabric-ca code directly so it
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
hyperspace is currently under `/cryptogen/organization/{my_orgs}`

**!:** make sure that the orgs and orderers found in the `configtx`
match the number of `orgs` in peers and orderers found under
`/cryptogen/organizations/{my_orgs}/**`

> CURRENTLY THE SYSTEM GENESIS BLOCK IS GENERATED RELATIVE TO WHERE YOU
CALL THE BASH SCRIPT SO CALL IT IN THE /CMDSCRIPTS/ DIR SO IT GETS
OUTPUT IN THE CORRECT PLACE

### SPAWNING AN ORDERER NODE
  - very important to note that the `orderer` does not let you set a
    config path with a flag. it looks for `orderer.yaml` in the default
    config paths which are `.` and `/env/hyperledger/fabric/` or via the
    env car `FABRIC_CFG_PATH`
  - also note: I copied the `configtx.yaml` to create the genesis block
    from `fabric-samples/test-network/configtx/configtx.yaml` and NOT
    from `fabric/sampleConfigs/configtx.yaml` OR `/fabric-samples/config/configtx.yaml`
    - `cp fabric-samples/test-network/configtx/configtx.yaml
      hyperspace/configtxgen/configtx.yaml`
  - NOTE: that `fabric/sampleConfigs` mirrors `fabric-samples/config`
    - verify that they're the same
  
  - so to spawn an `orderer` node I need to make sure the `orderer.yaml`
    is in the directory I called the binary from or add it to `FABRIC_CFG_PATH`
    - toDo: update orderer to accept custom config path
  - what I just did to get it to work was to copy `fabric-samples/config` right into `hyperspace`
    - then I moved just the `orderer.yaml` into the cmdscripts so the
      config could find it

  - it ran but then crashed because it was trying to access a directory
    that didn't exist
    - in the `orderer.yaml` I changed the `/var/**` dirs to `../orderer`
      (a local dir I temporarily created)
      - This WORKED

  - in `cmdscripts/` run `./spawn-orderer.sh`

### Spawning an orderer node from scratch via a static config
  - from now on, just make sure to call all cmdscripts from the
    cmdscript directory, think of it as your cmd-center

  - running `network.sh` will clean all the artifacts, generate the
    crypto, create a consortium, and spawn a node
    - eventually I'll have to check for existing material and only
      manually clean before spawn`
