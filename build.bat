@echo off
::

SET CGO_ENABLED=0
SET GO111MODULE=on

SET GOOS=windows
SET GOARCH=amd64

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

echo building for %GOOS%/%GOARCH%

SET target=build/wrest.exe
go build -ldflags="-s -w" -o %target% main.go

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

COPY README.md build\
COPY config.yml build\

IF "%1" == "" CMD /K
