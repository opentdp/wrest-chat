#!/bin/sh
#

set -e
set -o noglob

###########################################

export CGO_ENABLED=0
export GO111MODULE=on

export GOOS=windows
export GOARCH=amd64

####################################################################

RUN_NUMBER=${GITHUB_RUN_NUMBER:-0}

last_tag=`git tag | sort -V | tail -n 1`
prev_tag=`git tag | sort -V | tail -n 2 | head -n 1`
git log $prev_tag..$last_tag --pretty=format:"%s" | grep -v "^release" | sed 's/^/- /' | sort > RELEASE.md

version=`echo $last_tag | sed 's/^v//'`
sed -i "s/^const Version = \".*\"/const Version = \"$version\"/" args/build.go

build_version=$((`grep -oP 'BuildVersion = "\K\d+' args/build.go` + $RUN_NUMBER))
sed -i "s/^const BuildVersion = \".*\"/const BuildVersion = \"$build_version\"/" args/build.go

echo "build info - tag: $last_tag, version: $version, build: $build_version"

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

mkdir build/wcferry

cp README.md build/
cp config.yml build/
cp wcferry/bin/sdk.dll build/wcferry/
cp wcferry/bin/spy.dll build/wcferry/
cp wcferry/bin/wcf.exe build/wcferry/

sed -i 's#](./#](https://github.com/opentdp/wechat-rest/blob/master/#g' build/README.md

mv build wechat-rest
zip -r wechat-rest.zip wechat-rest/
