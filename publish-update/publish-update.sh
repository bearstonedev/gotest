#!/usr/bin/env bash

set -Eeuo pipefail

if ((BASH_VERSINFO[0] < 4))
then
  echo "Sorry, you need at least bash-4.0 to run this script."
  echo "Current version: ${BASH_VERSINFO[0]}"
  exit 1
fi


changedFiles=$(git diff --cached --shortstat)
if [[ -n $changedFiles ]]; then
  echo 'Changes detected, will not update version.'
  exit 1
fi

if [[ ${#} -ne 1 ]]; then
    echo 'Please provide the revision type as an argument (MAJOR, MINOR, PATCH).'
    exit 1
fi

currentVersion=()
IFS="." read -r -a currentVersion <<< "$(awk '{print $1}' publish-update/published.version)"
echo "Current version: v${currentVersion[0]}.${currentVersion[1]}.${currentVersion[2]}"

newVersion=("${currentVersion[@]}")
revision=${1^^}
case ${revision} in
  MAJOR|MINOR|PATCH) echo "Revision: ${revision}";;&
  MAJOR)
    (( ++newVersion[0] ))
    newVersion[1]=0
    newVersion[2]=0
    ;;
  MINOR)
    (( ++newVersion[1] ))
    newVersion[2]=0
    ;;
  PATCH) (( ++newVersion[2] ));;
  *)
    echo 'Must specify one of: MAJOR, MINOR, PATCH. Aborting update.'
    exit 1
    ;;
esac

newVersionFull="v${newVersion[0]}.${newVersion[1]}.${newVersion[2]}"
echo "New version: ${newVersionFull}"

echo "Updating published.version with ${newVersionFull} ..."
echo -n "${newVersion[0]}.${newVersion[1]}.${newVersion[2]}" > publish-update/published.version
git commit publish-update/published.version -m "ci: upgrade to ${newVersionFull}" --no-verify && git push

echo "Creating a new tag for ${newVersionFull} ..."
git tag "${newVersionFull}"
git push origin "${newVersionFull}"

echo "Updating Go proxy with ${newVersionFull} ..."
GOPROXY=proxy.golang.org go list -m "github.com/bearstonedev/gotest@${newVersionFull}"

echo "Upgrading examples module to ${newVersionFull} ..."
cd examples
go get "github.com/bearstonedev/gotest@${newVersionFull}"
go mod tidy
git commit go.mod go.sum -m "ci: upgrade examples to ${newVersionFull}" && git push

echo "Successfully published ${newVersionFull}!"