package main

import (
	"fmt"
	"net/http"
)

type burger int

func (m burger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
