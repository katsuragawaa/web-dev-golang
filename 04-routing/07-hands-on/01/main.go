package main

import (
	"fmt"
	"net/http"
)

//ListenAndServe on port ":8080" using the default ServeMux.
//Use HandleFunc to add the following routes to the default ServeMux:
//"/" "/dog/" "/me/
//Add a func for each of the routes.
//Have the "/me/" route print out your name.

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "index")
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "doggy")
}

func me(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "katsuragawa")
}
