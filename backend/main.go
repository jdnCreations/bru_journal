package main

import (
	"flag"
	"net/http"
)


func main() {
	listenAddr := flag.String("listenaddr", ":3000", "todo")
	flag.Parse()
	
	http.HandleFunc("/", handleRoot)

	http.ListenAndServe(*listenAddr, nil)	
}