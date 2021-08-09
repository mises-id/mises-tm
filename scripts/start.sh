#!/bin/sh
misestmd start --pruning nothing --grpc.address 0.0.0.0:9090 --home /misestm/node$ID/misestmd/ > misestmd.log &
tail -f /dev/null