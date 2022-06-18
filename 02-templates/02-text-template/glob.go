package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	goTemplate, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatal(err)
	}

	err = goTemplate.ExecuteTemplate(os.Stdout, "another-template.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
