#!/bin/sh

APP=${1:-wessage}

mkdir -p build

# for linux & mac
for GOOS_TMP in linux darwin; do
  export GOOS=$GOOS_TMP
  go build -v -o build/${APP}-${GOOS_TMP} main.go
done

# for windows
export GOOS=windows
go build -v -o build/${APP}-win.exe main.go

