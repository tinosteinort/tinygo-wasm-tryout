#!/usr/bin/env bash

docker run --rm -v $(pwd):/src tinygo/tinygo tinygo build -o ./src/web/main.wasm -target=wasm --no-debug ./src/go-wasm/main.go
