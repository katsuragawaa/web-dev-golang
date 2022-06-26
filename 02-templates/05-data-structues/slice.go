package main

import (
	"log"
	"os"
	"text/template"
)

var sliceTemplates *template.Template

func init() {
	sliceTemplates = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	names := []string{"Josh", "Will", "Max", "Lucas"}

	println("Iterate on a slice")
	err := sliceTemplates.ExecuteTemplate(os.Stdout, "slice.gohtml", names)
	if err != nil {
		log.Fatal(err)
	}

	println("\n\nSame but getting index and element as a variable")
	err = sliceTemplates.ExecuteTemplate(os.Stdout, "slice-with-variable.gohtml", names)
	if err != nil {
		log.Fatal(err)
	}

}
