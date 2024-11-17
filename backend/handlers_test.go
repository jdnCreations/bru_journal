package main

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_get_root(t *testing.T) {
	res, err := http.Get("http://localhost:3000")
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("%v", res.StatusCode)
}