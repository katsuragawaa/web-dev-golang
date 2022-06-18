package main

import "fmt"

type person struct {
	name string
	age  int
}

type agent struct {
	person
	license bool
}

func (p person) sayHi() {
	fmt.Println(p.name, "says hi")
}

func (a agent) sayHi() {
	fmt.Println("hello there", a.name)
}

type human interface {
	sayHi()
}

func speak(h human) {
	h.sayHi()
}

func main() {
	josh := person{
		"Josh",
		27,
	}

	james := agent{
		person{
			"James",
			40,
		},
		true,
	}

	speak(josh)
	speak(james)
}
