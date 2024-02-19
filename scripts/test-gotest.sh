#!/usr/bin/env bash

set -Eeuo pipefail

go test -coverprofile=coverage.out
totalCoverage=$(go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+')
if (( $(echo "${totalCoverage} 100" | awk '{print ($1 < $2)}') )); then
  echo "Coverage is too low: ${totalCoverage}"
  exit 1
fi