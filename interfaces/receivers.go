package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// if the reiever is value receiver,
// we can pass it either value or pointer
// But if reveiver is pointer,
// we can only pass pointer
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	info(&c)
	// info(c) // wrong
}
