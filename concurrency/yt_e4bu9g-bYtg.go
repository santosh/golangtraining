package main

import "fmt"

// SendValues takes a channel and sends arbitrary data to it.
func SendValues(c chan int) {
	c <- 8
}

func main() {
	fmt.Println("Go Channels Tutorial")

	values := make(chan int)
	defer close(values)

	go SendValues(values)

	value := <-values
	fmt.Println(value)
}
