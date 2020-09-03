#!/bin/bash

set -eux

DIRECTORIES=$(find . -name *.proto | xargs -I {} dirname {} | sort | uniq)

for DIRECTORY in ${DIRECTORIES};
do
    protoc \
        --go_out=plugins=grpc,paths=source_relative:. \
        --authz_out=paths=source_relative:. \
        -I . \
        ${DIRECTORY}/*.proto
done
