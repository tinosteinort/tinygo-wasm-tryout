#!/usr/bin/env bash

docker run --rm -v $(pwd):/src tinygo/tinygo /bin/bash -c "\
    cp /usr/local/tinygo/targets/wasm_exec.js /src/web/wasm_exec.js && \
    tinygo build -o ./src/web/main.wasm -target=wasm --no-debug ./src/go-wasm/main.go"
