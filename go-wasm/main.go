package main

import (
	"syscall/js"
)

func main() {
	println("Hello WASM world!")
}

//go:export printSomething
func printSomething(value js.Value) {
	println(value.String())
	//document := js.Global().Get("document")
	//output := document.Call("getElementById", "output")
	//output.Set("value", value)
}
