# Cryptogen

Creates the cryptographic resources and IDs needed for organizations to exist in the network.  

- in a production network, you would use Certificate Authorities, but cryptogen is fine for test nets


The `hyperspace.yaml` defines the 
 
 `fabricBinary`: name for a fabric `cryptogen` binary  
 `scriptPath`: path to cmdscript (will probably change)  
 `congigs`: a map[string]map[string] == map[organizationName]map[path/output][path] 
    
```
configs:
    organization-name:
        path: the path to the crypto-configurations for the organization (found rn in {networkName}/organizations/cryptogen
        output: where to output the cryptogen material
    org-2-name:
        path: ..
        output: ..
    ..
```