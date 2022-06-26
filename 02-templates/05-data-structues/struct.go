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

	err := structTemplates.ExecuteTemplate(os.Stdout, "struct.gohtml", buddha)
	if err != nil {
		log.Fatal(err)
	}
}
