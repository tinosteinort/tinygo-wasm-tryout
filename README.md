# tinygo / go - WASM tryout


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

While `build-(tiny)go-wasm.sh` is executed while `go generate ./...` the matching `load-wasm_exec.sh`
is copied to `web` directory. There are differences in the file for `go` / `tinygo`.

#### Switch between `go` and `tinygo`

Activate the favoured `go generate` line in `main.go`

