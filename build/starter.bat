@echo off

REG ADD HKEY_CURRENT_USER\Software\Tencent\WeChat /v NeedUpdateType /t REG_DWORD /d 0 /f >NUL 2>&1

:app
  start "" /b /wait wrest.exe

:wait
  echo retry after 3 seconds
  ping -n 3 localhost > nul
  goto app
