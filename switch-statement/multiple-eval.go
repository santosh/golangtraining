package main

import (
    "fmt"
)

func main() {
    switch "Jenny" {
    case "Tim", "Jenny":
        fmt.Println("Wassap Tim, or, err, Jenny")
    case "Marcus", "Medhi":
        fmt.Println("Both of your names start with M")
    case "Julian", "Sushant":
        fmt.Println("Wassap Julian/Sushant")
    }
}
