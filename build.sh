#!/bin/sh
#

set -e
set -o noglob

###########################################

export CGO_ENABLED=0
export GO111MODULE=on

build() {
    echo building for $1/$2
    target=build/wrest-$1-$2
    if [ x"$1" = x"windows" ]; then
        target="${target}.exe"
    fi
    GOOS=$1 GOARCH=$2 go build -ldflags="-s -w" -o $target main.go
}

####################################################################

build windows 386

####################################################################

cp README.md build/
cp wcf-sdk/libs/sdk.dll build/
cp wcf-sdk/libs/spy.dll build/

mv build wechat-rest
zip -r wechat-rest.zip wechat-rest/
