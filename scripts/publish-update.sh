#!/bin/bash

changedFiles=$(git diff --cached --shortstat)
if [[ -n $changedFiles ]]; then
  echo 'Changes detected, will not update version.'
  exit 1
fi

if [[ $# -ne 1 ]]; then
    echo 'Version number is expected in the form: #.#.#'
    echo 'The preceding v will be appended by the script.'
    exit 1
fi

versionNumber="v${1}"
git tag "$versionNumber"
git push origin "$versionNumber"
GOPROXY=proxy.golang.org go list -m "github.com/bearstonedev/gotest@${versionNumber}"
