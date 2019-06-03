package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

var urls = []string{
	"https://google.com",
	"https://sntsh.com",
	"https://twitter.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {
	for _, url := range urls {
		wg.Add(1)
		// you might have seen self calling (closures) function in JavaScript
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "\n%+v\n", err)
			}
			fmt.Fprintf(w, "\n%+v\n", resp.Status)
			defer wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {
	fmt.Println("Go WaitGroup Tutorial")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
