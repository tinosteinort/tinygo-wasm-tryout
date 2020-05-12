package main

import (
	"syscall/js"

	"github.com/siongui/godom"
)

func main() {
	println("Hello WASM world!")

	document := js.Global().Get("document")
	output := document.Call("getElementById", "output")
	output.Set("innerHTML", "set by go from main function")
}

//go:export update
func update() {
	godom.Document.GetElementById("output").SetValue("initial value")
}
