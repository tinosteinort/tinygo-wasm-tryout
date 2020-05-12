#!/usr/bin/env bash

#-v /home/tino/go/src:/usr/local/go/src
docker run --rm -v $(pwd):/src -v /home/tino/go/src:/root/go/src tinygo/tinygo /bin/bash -c '\
    cp /usr/local/tinygo/targets/wasm_exec.js /src/web/wasm_exec.js && \
    pwd && \
    ls -laF /root/go && \
    echo $GOROOT && \
    tinygo env && \
    export GOOS=js && \
    export GOARCH=wasm && \
    tinygo env && \
    tinygo build -o ./src/web/main.wasm -target=wasm --no-debug ./src/go-wasm/main.go'
