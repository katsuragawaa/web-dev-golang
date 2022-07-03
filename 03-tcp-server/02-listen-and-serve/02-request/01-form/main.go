package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type hotdog int

var tpl *template.Template

func (hd hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(req.Form)
	fmt.Println(req.PostForm)

	_ = tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var dog hotdog

	_ = http.ListenAndServe(":8080", dog)
}
