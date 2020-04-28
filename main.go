package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//go:generate ./build-go-wasm.sh
//go:generate ./build-tinygo-wasm.sh
func main() {

	dir := "go"

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "tinygo" {
		dir = "tinygo"
	}

	fmt.Println("Serve wasm compiled with", dir, "at http://localhost:8080/")

	http.Handle("/", http.FileServer(http.Dir("web/"+dir)))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
