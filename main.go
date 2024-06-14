package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Welcome to time study of GoLang")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World")
		io.WriteString(w, r.Method)
	}

	http.HandleFunc("/", h1)

	http.ListenAndServe(":5000", nil)
}
