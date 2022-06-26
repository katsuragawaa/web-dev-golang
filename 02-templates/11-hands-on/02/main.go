package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

type stock struct {
	Date, Open, High, Low, Close, Volume, AdjClose string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	f, err := os.Open("table.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// allocates a slice of capacity equal rows and returns a slice of length 0
	stocks := make([]stock, 0, len(rows))
	for i, row := range rows {
		if i == 0 { // skip header
			continue
		}

		s := stock{
			row[0],
			row[1],
			row[2],
			row[3],
			row[4],
			row[5],
			row[6],
		}
		stocks = append(stocks, s)
	}

	err = tpl.Execute(os.Stdout, stocks)
	if err != nil {
		log.Fatal(err)
	}
}
