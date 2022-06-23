package main

import (
	"log"
	"os"
	"text/template"
)

var goTemplate *template.Template

func init() {
	goTemplate = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := goTemplate.ExecuteTemplate(os.Stdout, "example.gohtml", "Hello from outside")
	if err != nil {
		log.Fatal(err)
	}
}
