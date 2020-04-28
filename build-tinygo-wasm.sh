#!/usr/bin/env bash

docker run --rm -v $(pwd):/src tinygo/tinygo tinygo build -o ./src/web/tinygo/main.wasm -target=wasm ./src/go-wasm/main.go
