#!/usr/bin/env bash

Title='FileServerUI'
Desc='A small and simple fileserver for home.'
CSSPath='web/static/style.css'

GOOS="linux" GOARCH="amd64" go build -o bin/fileserverui -ldflags="-X 'main.Title=$Title' -X 'main.Desc=$Desc' -X 'main.CSSPath=$CSSPath'" server.go
#GOOS="linux" GOARCH="arm64" go build -o bin/fileserverui -ldflags "-X 'main.Title=$Title, main.Desc=$Desc'" server.go
# =============================

GOOS="js" GOARCH="wasm" go build -o web/app.wasm cmd/fileserverui/app.go
#tinygo build -o web/app.wasm -no-debug -target wasm cmd/fileserverui/app.go
# =============================

sudo cp -r web/ vaiktorg.github.io/
# =============================

./bin/fileserverui # 7z a dist.zip src/web/ bin/