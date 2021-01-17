# Configtxgen

- contains the `configtx.yaml` needed by the `configtxgen` binary to create:
     - the genesis block + system channel
     - toDo: research what else this binary does
     
The `hyperspace.yaml` defines the 
 
 `fabricBinary`: name for a fabric `configtxgen` binary  
 `scriptPath`: path to cmdscript (will probably change)  
 `configPath`: path to the directory housing the `configtx.yaml`  
 `profile`: Idk yet, but needs to match a profile in the `configtx.yaml`  
 `channelID`: Idk yet, an ID for a channel (system or application?)  
 `output`: output path for the system genesis block  