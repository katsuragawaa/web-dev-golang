package main

import (
	"html/template"
	"log"
	"os"
)

var mapTemplates *template.Template

func init() {
	mapTemplates = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
	}

	err := mapTemplates.ExecuteTemplate(os.Stdout, "map.gohtml", sages)
	if err != nil {
		log.Fatal(err)
	}
}
