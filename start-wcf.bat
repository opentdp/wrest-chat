@echo off

set SDIR=%~dp0
set PATH=%PATH%;%SDIR%wcf-bin

cd %TEMP%

wcf start 10080
netstat -an | findstr :10080

cmd /k
