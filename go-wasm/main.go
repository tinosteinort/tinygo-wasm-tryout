package main

import (
	"strings"
	"syscall/js"
)

func main() {
	println("Hello WASM world!")

	document := js.Global().Get("document")
	output := document.Call("getElementById", "output")
	output.Set("innerHTML", "set by go from main function")

	registerCallbacks()

	select {}
}

func registerCallbacks() {
	js.Global().Set("update", js.FuncOf(jsUpdate))
}

func jsUpdate(this js.Value, value []js.Value) interface{} {
	s := value[0].String()
	update(s)
	return nil
}

func update(value string) {
	setOutputValue(modify(value))
}

func modify(value string) string {
	return strings.ToUpper(value)
}

func setOutputValue(value string) {
	document := js.Global().Get("document")
	document.Call("getElementById", "output").Set("innerHTML", value)
}
