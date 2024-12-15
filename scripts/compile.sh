#!/bin/bash

set -eo pipefail

goreleaser build --clean --snapshot

cd dist

rm artifacts.json config.yaml metadata.json

rhash -r --sha256 . -o checksums.txt
