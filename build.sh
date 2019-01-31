#!/bin/sh

APP=${1:-wessage}

mkdir -p build

# amd64 only
export GOARCH=amd64

# for linux & mac
for GOOS_TMP in linux darwin; do
  export GOOS=$GOOS_TMP
  go build -v -o build/${APP} main.go
  zip build/${APP}-${GOOS}-${GOARCH}.zip build/${APP}
done

# for windows
export GOOS=windows
go build -v -o build/${APP}.exe main.go
zip build/${APP}-${GOOS}-${GOARCH}.zip build/${APP}.exe
