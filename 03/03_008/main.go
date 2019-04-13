package main

import (
	"fmt"
	"net/http"
)

type burger int

func (m burger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Ziom-Key", "this is from Ziom")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var d burger
	http.ListenAndServe(":8080", d)
}
