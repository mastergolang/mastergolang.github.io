package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "8000"
	docsDir := "docs"
	fmt.Printf("Serving Master Golang at http://127.0.0.1:%s\n", port)
	http.ListenAndServe(":"+port, http.FileServer(http.Dir(docsDir)))
}
