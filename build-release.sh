#!/bin/bash

mkdir -p bin

TARGETS=(darwin/amd64 linux/amd64)

for target in ${TARGETS[@]}; do
  export GOOS=${target%/*}
  export GOARCH=${target##*/}
  go build
  mv md2slides bin/md2slides_"$GOOS"_"$GOARCH"
done
