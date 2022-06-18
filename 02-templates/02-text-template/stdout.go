package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	goTemplate, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	err = goTemplate.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}
