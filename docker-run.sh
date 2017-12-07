#!/bin/sh

docker run \
    -v /mnt/:/mnt/ \
    -d \
    --name commango \
    tkitsunai/commango