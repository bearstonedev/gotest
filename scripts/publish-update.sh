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
IFS="." read -r -a currentVersion <<< "$(awk '{print $1}' scripts/published.version)"
echo "Current version: v${currentVersion[0]}.${currentVersion[1]}.${currentVersion[2]}"

newVersion=("${currentVersion[@]}")
revision=${1^^}
case ${revision} in
  MAJOR|MINOR|PATCH) echo "Revision: ${revision}";;&
  MAJOR) (( ++newVersion[0] ));;
  MINOR) (( ++newVersion[1] ));;
  PATCH) (( ++newVersion[2] ));;
  *)
    echo 'Must specify one of: MAJOR, MINOR, PATCH. Aborting update.'
    exit 1
    ;;
esac

newVersionFull="v${newVersion[0]}.${newVersion[1]}.${newVersion[2]}"
echo "New version: ${newVersionFull}"

echo "Creating a new tag for ${newVersionFull} ..."
git tag "${newVersionFull}"
git push origin "${newVersionFull}"

echo "Updating Go proxy with ${newVersionFull} ..."
GOPROXY=proxy.golang.org go list -m "github.com/bearstonedev/gotest@${newVersionFull}"

echo "Updating published.version ..."
echo -n "${newVersion[0]}.${newVersion[1]}.${newVersion[2]}" > scripts/published.version
echo "Successfully published ${newVersionFull}!"