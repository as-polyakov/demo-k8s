package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<html><body><center><font size=\"32\">hello 3</font></center></body></html>\n")
}

func main() {
	http.HandleFunc("/hello2", hello)
	http.ListenAndServe(":8090", nil)
}
