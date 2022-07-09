package main

import (
	"html/template"
	"net/http"
)

// Take the previous program and change it so that:
// func main uses http.Handle instead of http.HandleFunc
// Constraint: Do not change anything outside func main

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(index))

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	d := req.URL
	tpl.Execute(w, d)
}
