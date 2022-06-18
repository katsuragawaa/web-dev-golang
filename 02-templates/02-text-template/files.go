package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Consider Template as a container for your templates
	goTemplate, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	// Print to the console the template
	err = goTemplate.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

	// The template has too ParseFiles method that add more templates
	goTemplate, err = goTemplate.ParseFiles("another-template.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	// Then you can choose which one of then you want to execute
	err = goTemplate.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = goTemplate.ExecuteTemplate(os.Stdout, "another-template.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Now Execute will run the first one that it find
	err = goTemplate.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}
