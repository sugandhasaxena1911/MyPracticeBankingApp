package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe("localhost:7000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My firstHello")
}
