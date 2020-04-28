#!/usr/bin/env bash

GOOS=js GOARCH=wasm go build -o web/go/main.wasm go-wasm/main.go
