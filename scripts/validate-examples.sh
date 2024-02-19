#!/usr/bin/env bash

set -Eeuo pipefail

cd examples
go test
go clean && go build -o build-test-examples
rm build-test-examples
