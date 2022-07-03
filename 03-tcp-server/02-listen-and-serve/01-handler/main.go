package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (hd hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "anything can be a handler")
}

func main() {
	var dog hotdog

	err := http.ListenAndServe(":8080", dog)
	if err != nil {
		return
	}
}
