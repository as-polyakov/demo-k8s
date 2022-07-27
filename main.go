package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello 3\n")
}

func main() {
	http.HandleFunc("/hello2", hello)
	http.ListenAndServe(":8090", nil)
}
