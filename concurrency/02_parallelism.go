package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// concurrency is threading, parallelism is multiprocessing
// concurrency can share data, parallelism can't

// concurrency is context switching, parallelism runs in parallel
// a process can have multiple threads

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
