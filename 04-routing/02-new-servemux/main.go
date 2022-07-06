package main

import (
	"io"
	"net/http"
)

type hotdog int

func (dog hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy")
}

type hotcat int

func (cat hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty")
}

func main() {
	var dog hotdog
	var cat hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog", dog)
	mux.Handle("/cat", cat)

	http.ListenAndServe(":8080", mux)
}
