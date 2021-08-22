#!/bin/sh
misestmd start --pruning nothing --home /misestm/node$ID/misestmd/ > misestmd.log &
tail -f /dev/null