package main

import (
	"fmt"
	"net/http"
)

type burger int

func (m burger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d burger
	http.ListenAndServe(":8080", d)
}
