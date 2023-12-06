@ECHO OFF
::

SET GOARCH=386

SET CGO_ENABLED=0
SET GO111MODULE=on

SET SDIR=%~dp0
SET PATH=%PATH%;%SDIR%wcf-bin

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

CD /d %~dp0

go mod tidy
go run main.go
