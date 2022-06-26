package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func firstThree(str string) string {
	cleanStr := strings.TrimSpace(str)
	return cleanStr[:3]
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"uppercase":  strings.ToUpper,
	"firstThree": firstThree,
	"fdate":      monthDayYear,
}

func init() {
	// need to attach functions before parsing the file
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
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

	currentTime := time.Now()

	data := struct {
		Wisdom []sage
		Time   time.Time
	}{
		[]sage{buddha, gandhi},
		currentTime,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}
