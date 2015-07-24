set outputpath=..\server\src\protos
set protofile=%2%.proto
set outidfile=%2%_id.go
set packagename=%1%
set protoc_exe=..\tools\protoc.exe
mkdir %outputpath%\%1%
"..\tools\protoc.exe" %protofile% --plugin=protoc-gen-go=..\tools\protoc-gen-go.exe --go_out %outputpath%\%packagename% --proto_path "."
"..\tools\gotoolchain\bin\msgidgen.exe" -protoc=%protoc_exe% -proto=%protofile% -out=%outputpath%\%packagename%\%outidfile%