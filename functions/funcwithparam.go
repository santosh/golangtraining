package main

import (
    "fmt"
)

func main() {
    greet("jane")
    greet("john")
}

func greet(name string) {
    fmt.Println("Good morning", name)
}

