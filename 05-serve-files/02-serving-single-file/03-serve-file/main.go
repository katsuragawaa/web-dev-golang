package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/toby", toby)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// this calls localhost:8080/toby that returns the image
	io.WriteString(w, `<img src="toby">`)
}

func toby(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "toby.jpeg")
}
