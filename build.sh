#!/bin/sh
#

set -e
set -o noglob

###########################################

export CGO_ENABLED=0
export GO111MODULE=on

build() {
    local os=${2:-linux}
    local arch=${3:-amd64}
    local sbin=build/${1:-bin}-$os-$arch
    if [ x"$os" = x"windows" ]; then
        sbin="${sbin}.exe"
    fi
    echo building for $os/$arch
    GOOS=$os GOARCH=$arch go build -ldflags="-s -w" -o $sbin main.go
}

####################################################################
# fix version

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
# build binary

sed -i 's#](./#](https://github.com/opentdp/wrest-chat/blob/master/#g' README.md

if [ -f webview/public/browser/index.html ]; then
    cp -av webview/public/browser/. public/
fi

build wrest linux amd64
build wrest windows amd64

####################################################################
# package for linux

cp -av build linux
cp README.md linux/
cp config.yml linux/

rm -rf linux/starter.bat
rm -rf linux/wrest-windows-amd64.exe
mv linux/wrest-linux-amd64 linux/wrest

sed -i '/SdkLibrary:/d' linux/config.yml
sed -i 's/127.0.0.1:7601/192.168.1.2:7601/g' linux/config.yml
sed -i 's/127.0.0.1:7600/0.0.0.0:7600/g' linux/config.yml

cd linux
zip -r ../wrest-linux-v$version.zip .
cd ..

####################################################################
# package for windows

cp -av build windows
cp README.md windows/
cp config.yml windows/

rm -rf windows/wrest-linux-amd64
mv windows/wrest-windows-amd64.exe windows/wrest.exe

mkdir -p windows/wcferry
wget https://github.com/lich0821/WeChatFerry/releases/download/v39.2.4/v39.2.4.zip
unzip -d windows/wcferry v39.2.4.zip && rm -f v39.2.4.zip

cd windows
zip -r ../wrest-windows-v$version.zip .
cd 
