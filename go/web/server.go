package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", myHandler1)
	http.ListenAndServe(":8080", nil)
}

func myHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world \n")
}
