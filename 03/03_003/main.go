package main

import (
	"log"
	"net/http"
	"text/template"
)

type burger int

var tpl *template.Template

func (m burger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d burger
	http.ListenAndServe(":8080", d)
}
