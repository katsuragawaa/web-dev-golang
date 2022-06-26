package main

import (
	"log"
	"os"
	"text/template"
)

var structMapTpl *template.Template

type sage2 struct {
	Name  string
	Motto string
}

func init() {
	structMapTpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	buddha := sage2{
		"Buddha",
		"The belief of no beliefs",
	}

	gandhi := sage2{
		"Gandhi",
		"Be the change",
	}

	mlk := sage2{
		"MLK",
		"I have a dream",
	}

	sages := []sage2{buddha, gandhi, mlk}

	err := structMapTpl.ExecuteTemplate(os.Stdout, "struct-map.gohtml", sages)
	if err != nil {
		log.Fatal(err)
	}

}
