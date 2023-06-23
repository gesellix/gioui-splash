package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	absPath, err := filepath.Abs("./wasm")
	if err != nil {
		fmt.Println("Wrong path?", err)
		return
	}
	webRoot := filepath.Clean(absPath)
	dir := http.Dir(webRoot)
	fmt.Println("serving from", dir)
	err = http.ListenAndServe(":9090", http.FileServer(dir))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
