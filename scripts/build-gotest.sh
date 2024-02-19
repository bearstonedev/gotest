#!/usr/bin/env bash

set -Eeuo pipefail

go clean && go build -o build-test
rm build-test