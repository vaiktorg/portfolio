#!/usr/bin/env bash

Title='Dashboard'
Desc='Personal porfolio and random tools.'
CSSPath='web/static/style.css'
DSTPath='dst/notesui/"

GOOS="linux" GOARCH="amd64" go build -o bin/dashboard -ldflags="-X 'main.Title=$Title' -X 'main.Desc=$Desc' -X 'main.CSSPath=$CSSPath' -X 'main.DSTPath=$DSTPath'" server.go
#GOOS="linux" GOARCH="arm64" go build -o bin/fileserverui -ldflags "-X 'main.Title=$Title, main.Desc=$Desc'" server.go
# =============================

GOOS="js" GOARCH="wasm" go build -o web/app.wasm cmd/dashboard/app.go
# =============================

sudo cp -r web/ vaiktorg.github.io/
# =============================

./bin/dashboard # 7z a dist.zip src/web/ bin/