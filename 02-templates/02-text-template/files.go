package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Consider Template as a container for your templates
	goTemplate, err := template.ParseFiles("templates/tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	// The template has too ParseFiles method that add more templates
	goTemplate, err = goTemplate.ParseFiles("templates/another-template.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	// Then you can choose which one of then you want to execute
	println("\n-----TPL-----")
	err = goTemplate.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	println("\n-----Another-----")
	err = goTemplate.ExecuteTemplate(os.Stdout, "another-template.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Now Execute will run the first one that it find
	println("\n-----Execute-----")
	err = goTemplate.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}
