package main

import (
	"log"
	"os"
	"text/template"
)

var structTemplates *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	structTemplates = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	buddha := sage{
		"buddha",
		"the belief of no beliefs",
	}

	println("struct")
	err := structTemplates.ExecuteTemplate(os.Stdout, "struct.gohtml", buddha)
	if err != nil {
		log.Fatal(err)
	}

	println("\n\nstruct with variables")
	err = structTemplates.ExecuteTemplate(os.Stdout, "struct-with-variable.gohtml", buddha)
	if err != nil {
		log.Fatal(err)
	}
}
