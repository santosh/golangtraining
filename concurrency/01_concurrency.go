package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Concurrency is the composition of independently executing processes.
Parallelism is the simultaneous execution of (possibily related) computation.
Concurrency is about dealing with lots of things at once.
Parallelism is about doing lots of things at once.
*/

// We can make any routine (function/method) a goroutine by adding 'go' keyword.
// But main function will not wait for those goroutines.
// 1. For main func to wait for goroutines, we add a WaitGroup.
// https://golang.org/pkg/sync/#WaitGroup

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // <- this tells goroutine is done,
}

func main() {
	wg.Add(2) // <- this tells we are going to add 2 goroutines
	go say("Hey")
	go say("There")
	wg.Wait() // <-  will wait for 2 routines to complete here

	wg.Add(1) // <- another waitgroup which sets counter for 1 routine
	go say("Hi")
	wg.Wait()
}
