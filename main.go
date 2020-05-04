package main

import (
	"fmt"
	"log"
	"net/http"
)

//go:generate ./build-tinygo-wasm.sh
func main() {

	fmt.Println("Open browser: http://localhost:8080/")

	http.Handle("/", http.FileServer(http.Dir("web/")))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
