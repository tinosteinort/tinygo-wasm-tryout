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
}

//go:export update
func update() {
	value := inputValue()
	setOutputValue(modify(value))
}

func modify(value string) string {
	return strings.ToUpper(value)
}

func inputValue() string {
	document := js.Global().Get("document")
	return document.Call("getElementById", "input").Get("value").String()
}

func setOutputValue(value string) {
	document := js.Global().Get("document")
	document.Call("getElementById", "output").Set("innerHTML", value)
}
