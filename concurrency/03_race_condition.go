package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// race condition is when two routines/thread try to access same
// memory (variable)

// we use mutual exclusion lock...

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		// after this lock is acquired
		// any other go rountine is not able to write to
		// below variables
		mutex.Lock()
		counter++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		fmt.Println(s, i, "Counter: ", counter)
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final counter: ", counter)
}
