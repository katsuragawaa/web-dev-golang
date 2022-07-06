package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy")
}

func cat(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty")
}

func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

	http.ListenAndServe(":8080", nil)
}
