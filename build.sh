#!/bin/sh
#

set -e
set -o noglob

###########################################

export CGO_ENABLED=0
export GO111MODULE=on

export GOOS=windows
export GOARCH=386

####################################################################

RUN_NUMBER=${GITHUB_RUN_NUMBER:-0}

last_tag=`git tag | sort -V | tail -n 1`
prev_tag=`git tag | sort -V | tail -n 2 | head -n 1`
git tag -l $last_tag -n | cut -d' ' -f2- | sed 's/^ *//' > RELEASE.md
git log $prev_tag..$last_tag --pretty=format:"%s" | grep -v "^release" | sed 's/^/- /' | sort >> RELEASE.md

sed -i '/./,$!d' RELEASE.md

####################################################################

echo building for $GOOS/$GOARCH

target=build/wrest.exe
go build -ldflags="-s -w" -o $target main.go

####################################################################

cp README.md build/
cp wcferry/libs/sdk.dll build/
cp wcferry/libs/spy.dll build/

mv build wechat-rest
zip -r wechat-rest.zip wechat-rest/
