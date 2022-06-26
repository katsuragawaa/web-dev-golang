package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var fm = template.FuncMap{
	"uppercase":  strings.ToUpper,
	"firstThree": firstThree,
}

var tpl *template.Template

func init() {
	// need to attach functions before parsing the file
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(str string) string {
	cleanStr := strings.TrimSpace(str)
	return cleanStr[:3]
}

func main() {
	buddha := sage{
		"buddha",
		"the belief of no beliefs",
	}

	gandhi := sage{
		"gandhi",
		"be the change",
	}

	data := struct {
		Wisdom []sage
	}{[]sage{buddha, gandhi}}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}
