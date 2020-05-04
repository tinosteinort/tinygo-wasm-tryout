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

The matching `wasm_exec.js` can be used from the `tinygo` docker image. Run `load-wasm_exec.sh`
to get a fresh copy.
