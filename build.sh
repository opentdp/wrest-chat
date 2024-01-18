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
git log $prev_tag..$last_tag --pretty=format:"%s" | grep -v "^release" | sed 's/^/- /' | sort > RELEASE.md

####################################################################

if [ -f webview/public/browser/index.html ]; then
    ls -al webview
    ls -al webview/public
    ls -al webview/public/browser
    cp -av webview/public/browser/. public/
fi

echo building for $GOOS/$GOARCH

target=build/wrest.exe
go build -ldflags="-s -w" -o $target main.go

####################################################################

cp README.md build/

mv build wechat-rest
zip -r wechat-rest.zip wechat-rest/
