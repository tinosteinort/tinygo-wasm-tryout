#!/usr/bin/env bash

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./web/wasm_exec.js

GOOS=js GOARCH=wasm go build -o web/main.wasm go-wasm/main.go
