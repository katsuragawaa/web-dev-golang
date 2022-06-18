package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	goTemplate, err := template.ParseFiles("templates/tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file: ", err)
	}

	// File also implements Writer interface
	err = goTemplate.Execute(newFile, nil)
	if err != nil {
		log.Fatal(err)
	}
}
