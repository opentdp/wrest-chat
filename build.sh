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

cp README.md build/
cp config.yml build/
sed -i 's#](./#](https://github.com/opentdp/wrest-chat/blob/master/#g' build/README.md

mkdir -p build/wcferry
wget https://github.com/lich0821/WeChatFerry/releases/download/v39.0.14/v39.0.14.zip
unzip -d build/wcferry v39.0.14.zip && rm -f v39.0.14.zip

mv build wrest-chat
zip -r wrest-chat.zip wrest-chat/
