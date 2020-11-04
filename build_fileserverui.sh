#!/usr/bin/env bash

Title='FileServerUI'
Name=$Title
Desc='A small and simple fileserver for home.'
ShortName='FSUI'

GOOS="linux" GOARCH="amd64" go build -o bin/fileserverui.exe -ldflags "-X main.Title=$Title main.Name=$Name main.Desc=$Desc main.ShortName=$ShortName" cmd/fileserverui/server.go
GOOS="js" GOARCH="wasm" go build -o web/app.wasm app.go

sudo cp -r web/ vaiktorg.github.io/

./bin/fileserverui.exe # 7z a dist.zip src/web/ bin/