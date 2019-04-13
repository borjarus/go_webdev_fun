package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type burger int

var tpl *template.Template

func (m burger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d burger
	http.ListenAndServe(":8080", d)
}