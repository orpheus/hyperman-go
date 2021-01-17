# Hyperspace 
## Go far, move fast.


## 0.1.0
- requires custom named fabric binaries
    - go into fabric and build two peers, peer-01 & peer-02, and one orderer (default name)
        - make sure these binaries are in your PATH
        - alternatively, you can make the binary names whatever you'd like and then just update the corresponding hyperpsace.yaml associated with that node to point to that binary
    
- Golang ^1.15 (is what I'm building it with)


#### Build

`./scripts/build/.sh` - build the binary

##### Run 

`hyperspace` - run the binary

##### Kill 

`control-c` - will gracefully shut down the processes 

##### Clean

`./cmdscripts/clean-artifacts.sh`

- run this if you experience errors starting the network about not being able to find the correct
system channel, genesis block, or certificates. It will clean the generated resources.

## Terms

`control_center` - The root of the hyperspace project. This is where the code operates from and controls various aspects of the network. 

`command_center` - The `cmdScripts` directory. This is where the bash scripts (commanders) live.
They go out into various parts of the network and issue commands to binaries.

I think of it like this: There is one main CONTROLLER who controls the COMMANDERS which issue commands to the binaries or SOLDIERS.
The CONTROLLER is the Go code and lives at the root of the application and controls the COMMANDERS which are the bash scripts that take their posts in 
different networks and file paths to issue commands to the SOLDIERS which are the fabric binaries. 
The Controller issues commands to the Commanders which in turn issues commands to the Soldiers.

## Environment

`HYPERSPACE_PATH=path/to/hyperspace`

Hyperspace project path. Needed to generate absolute paths relative to
creating directories and calling command scripts.

`HYPERSPACE_NETROOT=path/to/networks`

This is the path to where you want the generated networks to live.
 
Defaults to `HYPERSPACE_PATH/networks`




