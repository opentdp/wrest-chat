@ECHO OFF
::

SET CGO_ENABLED=0
SET GO111MODULE=on

set SDIR=%~dp0
set PATH=%PATH%;%SDIR%wcf-bin

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CD /d %~dp0

go mod tidy
go run main.go
