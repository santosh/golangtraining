package main

import (
    "fmt"
)

func hello() {
    fmt.Print("Hello ")
}

func world() {
    fmt.Println("world")
}

func main() {
    defer world()
    hello()
    fmt.Println("after")
    defer fmt.Println("final")
}
