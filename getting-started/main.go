package main

import (
	"fmt"
	"math"
)

// HANDS ON 1

// create a type square
type square struct {
	side float64
}

// create a type circle
type circle struct {
	radius float64
}

// attach a method to each that calculates area and returns it
func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// create a type shape which defines an interface as anything which has the area method
type shape interface {
	area() float64
}

// create a func info which takes type shape and then prints the area
func printArea(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	// create a value of type square
	s := square{10.0}
	// create a value of type circle
	c := circle{5.0}
	// use func info to print the area of square
	printArea(s)
	// use func info to print the area of circle
	printArea(c)
}
