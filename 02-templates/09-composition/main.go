package main

import (
	"html/template"
	"log"
	"os"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	year := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"CSCI-40", "Intro to Go", "4"},
				{"CSCI-130", "Mobile", "4"},
			}},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"CSCI-50", "Advanced Go", "4"},
				{"CSCI-190", "Mobile 2", "4"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, year)
	if err != nil {
		log.Fatal(err)
	}
}
