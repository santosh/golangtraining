package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 12
	var y = 12.12651

	// following is a narrowing conversion
	// trims out float to int
	fmt.Println(int(y) + x)

	// rune to string
	var z rune = 'a' // rune is an alias for int32
	fmt.Println(z)
	fmt.Println(string(z))

	// slice of bytes to string
	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'}))

	// string to slice of bytes
	fmt.Println([]byte("hello"))

	// ascii to int
	var a = "12"

	b, _ := strconv.Atoi(a)
	fmt.Println(b)

	// int to ascii
	c := 14
	d := "I have this many: " + strconv.Itoa(c)
	fmt.Println(d)
}
