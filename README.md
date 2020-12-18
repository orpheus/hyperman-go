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