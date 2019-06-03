package main

import "fmt"

// assertion is related to casting, with different syntax
// casting: int("5")
// assertion: "5".(int)

func try1() {
	// this will fail
	// name := "Sydney"
	var name interface{} = "Sydney"

	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("Value if not a string\n")
	}
}

func try2() {
	var val interface{} = 7
	fmt.Println(val.(int) * 6)
}

func main() {
	try2()
}
