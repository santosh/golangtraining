package main

import (
	"fmt"
	"time"
)

// SendValues takes a channel and sends arbitrary data to it.
func SendValues(c chan string) {
	fmt.Println("Executing Goroutine")
	time.Sleep(1 * time.Second)
	c <- "Hello World"
	fmt.Println("Finished Executing Goroutine")
}

func main() {
	fmt.Println("Go Channels Tutorial")

	values := make(chan string, 2)
	defer close(values)

	go SendValues(values)
	go SendValues(values)

	value := <-values
	fmt.Println(value)

	time.Sleep(1 * time.Second)
}
