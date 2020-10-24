#!/usr/bin/env bash

GOOS="linux" GOARCH="amd64" go build -o bin/AppExample.exe server.go
GOOS="js" GOARCH="wasm" go build -o web/app.wasm app.go

cp -r web/ vaiktorg.github.io/

./bin/AppExample.exe # 7z a dist.zip src/web/ bin/