package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/img.png", img)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>foo ran</h1>")
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "error executing template", http.StatusInternalServerError)
	}
}

func img(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "img.png")
}
