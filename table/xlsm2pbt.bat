@echo off
set name=%1%
set xlsm_dir=.
set pbt_dir=.\obj
set xls2pbt_exe=..\tools\gotoolchain\bin\xls2pbt.exe
set protoc_exe=..\tools\protoc.exe
set extra=%3%
if "%2%"=="" (
set protoname=%1%
) else set protoname=%2%

@echo on

%xls2pbt_exe% -protoc=%protoc_exe% -proto=..\proto\%protoname%.proto -xls=%name%.xlsm -pbt_out=%pbt_dir%\%name%.pbt -haltonerr=true %extra%