Create networks using

1. the current path of where the hyperspace binary was called
2. path argument (--config)
3. NETWORK_ROOT env var

path argument should override all, NETWORK_ROOT overrides current path, current path
used if neither of the others were provided

Create a fn to generate the templates-fabric so if people want to get a template with
all the comments really easily they can just call a hyperspace function vs opening 
fabric or fabric-samples and doing a copy/paste.
- this should just shell out to a bash script to do a copy

