package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<center><font dize="32">hello 3</font></center>\n")
}

func main() {
	http.HandleFunc("/hello2", hello)
	http.ListenAndServe(":8090", nil)
}
