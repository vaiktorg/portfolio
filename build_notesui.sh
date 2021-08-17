#!/usr/bin/env bash

# =============================
#Header Info
AppName="notes"
Title="NotesUI"
Desc='A small and simple fileserver for home.'

# Paths
WEB_DIR='models/web'
CSS_PATH="web/static/style.css"
DST_DIR="../../dst/vaiktorg.github.io/$Title"

# =============================
#GoFiles
GoWASMFilePath='app/notes/app.go'
GoServerFilePath='models/server/server.go'
GoStaticFilePath='models/server/staticwebsite.go'

# Export Paths
BinFilePath="bin/$Title/$AppName"
WASMFilePath="bin/$Title/web/app.wasm"

# =============================
#Build this server to host.
GOOS="linux" GOARCH="amd64" \
  go build -o $BinFilePath \
  -ldflags="-X 'main.Title=$Title' -X 'main.Desc=$Desc' -X 'main.CSSPath=$CSS_PATH' -X 'main.DSTPath=$DST_DIR' -X 'main.WEBPath=web'" \
  $GoServerFilePath

# =============================
#Build WASMApp
GOOS="js" GOARCH="wasm" go build -o "$WASMFilePath" "$GoWASMFilePath"

# =============================
cp -r "$WEB_DIR" "bin/$Title/"

# shellcheck disable=SC2164
cd bin/$Title
./$AppName