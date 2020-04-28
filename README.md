# GO / WASM tryout


## Compile code to WASM 

`go-lang` needs to be installed. `tinygo` is used via docker.

Compile `go` and `tinygo`code to wasm:
```
go generate ./...
```

This needs to be executed every time when the `go-wasm/main.go` file has changed.



## Run server

Run with `go wasm` artifact:
```
go run main.go "go"
```

Run with `tinygo wasm` artifact:
```
go run main.go "tinygo"
```

When open the URL, there should be something in the browser console.



### Notes

For `go` the `wasm_exec.js` from `$(go env GOROOT)/misc/wasm/wasm_exec.js` is used:
```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/
```

For `tinygo` the `wasm_exec.js` from the tinygo repository is used:
https://github.com/tinygo-org/tinygo/blob/master/targets/wasm_exec.js
