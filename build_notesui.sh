#!/usr/bin/env bash

Title='NotesUI'
Desc='A small and simple note taking tool for home.'
CSSPath='web/static/notes.css'


GOOS="linux" GOARCH="amd64" go build -v -o bin/notes -ldflags="-X 'main.Title=$Title' -X 'main.Desc=$Desc' -X 'main.CSSPath=$CSSPath'" src/server/server.go
#GOOS="linux" GOARCH="arm64" go build -o bin/fileserverui -ldflags "-X 'main.Title=$Title, main.Desc=$Desc'" server.go
# =============================

GOOS="js" GOARCH="wasm" go build -v -o src/web/app.wasm cmd/notes/app.go
# =============================

sudo cp -r web/ vaiktorg.github.io/
# =============================

./bin/notes # 7z a dist.zip src/web/ bin/