# GO / WASM tryout


## Compile code to WASM 

Requirements:
 * `go-lang`
 * docker (used for `tinygo`)

Create `wasm` artifacts:
```
go generate ./...
```

This needs to be executed every time when the `go-wasm/main.go` file has changed.



## Run server

```
go run main.go
```

When open the URL `http://localhost:8080`, there should be something in the browser console.



### Notes

For `go` the `wasm_exec.js` from `$(go env GOROOT)/misc/wasm/wasm_exec.js` is used:
```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/go
```

For `tinygo` the `wasm_exec.js` from the tinygo repository is used:
https://github.com/tinygo-org/tinygo/blob/master/targets/wasm_exec.js
