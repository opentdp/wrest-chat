#!/bin/sh
#

set -e
set -o noglob

####################################################################

if [ ! -d node_modules ] ; then
    echo Installing node modules...
    npm install
fi

npm run build
