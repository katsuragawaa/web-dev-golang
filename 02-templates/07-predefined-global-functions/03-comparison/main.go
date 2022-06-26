package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	data := struct {
		N1 int
		N2 int
	}{
		7,
		9,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
