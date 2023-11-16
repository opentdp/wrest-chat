@echo off
::

SET CGO_ENABLED=0
SET GO111MODULE=on

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CALL :build windows 386
CALL :build windows amd64

GOTO :EOF

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

:build
  SET GOOS=%1
  SET GOARCH=%2
  echo building for %1/%2
  SET target=build/tdp-blackbox-%1-%2
  IF "%1"=="windows" (
    SET target=%target%.exe
  )
  go build -ldflags="-s -w" -o %target% main.go
  GOTO :EOF

IF "%1" == "" CMD /K
