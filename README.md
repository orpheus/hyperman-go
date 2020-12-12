# Hyperman

### Hyperledger Command Center

Hyperman is a hyperledger command center that allows you to dynamically
spawn, create, destroy, and alter hyperledger nodes. It maintains a
memorybank of configurations and allows you to create, read, edit, and
delete, store, and version any configuration required to spawn a production ready network.

## Prerequisites

Hyperman needs access to the fabric and fabric-ca code directly so it
can dynamically build the cmd binaries.

For running locally, pull down `fabric` and `fabric-ca` and place the
following scrips in the respective paths.


1. add `buildCmd.sh` and `buildNodes.sh` to `fabric/scripts/`

2. add `fabric/scripts` and `/fabric/bin` to $PATH 

