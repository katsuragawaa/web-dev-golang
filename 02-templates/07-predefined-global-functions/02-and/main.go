package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	common := user{
		Name:  "Josh",
		Admin: false,
	}

	admin := user{
		Name:  "Max",
		Admin: true,
	}

	nobody := user{
		Name:  "",
		Admin: true,
	}

	users := []user{common, admin, nobody}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatal(err)
	}
}
