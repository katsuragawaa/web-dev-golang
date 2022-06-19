package main

import (
	"log"
	"os"
	"text/template"
)

var goTemplate *template.Template

func init() {
	// Must method does the error checking
	goTemplate = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := goTemplate.ExecuteTemplate(os.Stdout, "another-template.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
