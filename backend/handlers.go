package main

import (
	"fmt"
	"io"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("someone came to da site. omg.\n")
	io.WriteString(w, "This is my website!\n")
}