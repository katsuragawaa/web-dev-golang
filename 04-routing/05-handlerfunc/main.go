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
	// conversion
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/cat", http.HandlerFunc(cat))

	http.ListenAndServe(":8080", nil)
}

type batata int

func similarConvertion() {
	var b batata
	b = 7

	var y int
	//y = b <- wont work, different types
	y = int(b)

	println(y)
}
