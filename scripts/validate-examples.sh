#!/usr/bin/env bash

set -Eeuo pipefail

cd examples
go clean && go build -o build-test-examples && go test
rm build-test-examples