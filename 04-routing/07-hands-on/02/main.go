package main

import (
	"html/template"
	"net/http"
)

// Take the previous program in the previous folder and change it so that:
// a template is parsed and served
// you pass data into the template

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	d := req.URL
	tpl.Execute(w, d)
}
