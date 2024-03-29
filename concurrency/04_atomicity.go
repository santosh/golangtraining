package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// atomicity works similar to mutexes...
// variables which are defined as to be atomic are only accessed
// by one goroutine at a single point of time.

var wg sync.WaitGroup
var counter int64

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter: ", atomic.LoadInt64(&counter))
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
