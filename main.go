package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello2-harness\n")
}
func liveness(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Satus OK\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", liveness)
	fmt.Println("I am a new versioni 7")
	http.ListenAndServe(":80", nil)
}
