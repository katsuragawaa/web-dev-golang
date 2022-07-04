package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (dog hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Bro-Key", "bro")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintln(w, "<h1>Any response that you want</h1>")
}

func main() {
	var dog hotdog
	http.ListenAndServe(":8080", dog)
}
