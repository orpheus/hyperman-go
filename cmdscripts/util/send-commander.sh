# Send a cmdscript to it's network location

# To use hyperspace binary anywhere you just need to
# set the HYPERSPACE_LOCATION or HYPERSPACE_PATH in your
# bash profile. On calling `hyperspace` it will take this cfg
# path and append it to all of it's paths as an abs path
# allowing hyperledger to operate anywhere in the os' filesystem.

# This function will take a path and cd  the operating
# cmdscript into it so it may perform it's functions there
# this allows the path passed in to be relative or abs
# depending on your preference

source util/scriptUtils.sh

found_cmd_location=false
COMMAND_CENTER=""

for var in "$@"
do
  if [ "$found_cmd_location" = true ]; then
    if [ -n "$var" ]; then
      infoln "Command center location: $var"
      COMMAND_CENTER="$var"
    else
       fatalln "missing --command-center flag; must specify location"
    fi
  fi

  if [[ $var == --command-center ]]; then
    found_cmd_location=true
  fi
done

if [ -z "$COMMAND_CENTER" ]; then
     fatalln "missing --command-center flag; must specify location"
fi

set -x
cd "$COMMAND_CENTER"
res=$?
{ set +x ; } 2>/dev/null
if [ "$res" -ne 0 ]; then
    fatalln "Failed to send commander to: $COMMAND_CENTER"
fi

infoln "Commander sent to: $COMMAND_CENTER"