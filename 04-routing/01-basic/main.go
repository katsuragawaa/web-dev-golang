package main

import (
	"io"
	"net/http"
)

type hotdog int

func (dog hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy")
	case "/cat":
		io.WriteString(w, "kitty")
	}
}

func main() {
	var dog hotdog
	http.ListenAndServe(":8080", dog)
}
