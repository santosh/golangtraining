package main

import (
	"fmt"
	"log"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	log.Println(z.area())
}

func main() {
	s := Square{10}
	c := Circle{10}
	info(s)
	info(c)
}
