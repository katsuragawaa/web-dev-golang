package main

import (
	"log"
	"os"
	"text/template"
)

var smlTemplate *template.Template

type sage3 struct {
	Name  string
	Motto string
}

type car struct {
	Model string
	Doors int
}

type items struct {
	Wisdom    []sage3
	Transport []car
}

func init() {
	smlTemplate = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	buddha := sage3{
		"Buddha",
		"The belief of no beliefs",
	}

	gandhi := sage3{
		"Gandhi",
		"Be the change",
	}

	mlk := sage3{
		"MLK",
		"I have a dream",
	}

	sages := []sage3{buddha, gandhi, mlk}

	toyota := car{
		"Toyota",
		4,
	}

	mobi := car{
		"Mobi",
		2,
	}

	cars := []car{toyota, mobi}

	data := items{sages, cars}

	err := smlTemplate.ExecuteTemplate(os.Stdout, "struct-map-list.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}

}
