package main

import (
    "fmt"
)

func main() {
    switch "Medhi" {
    case "Daniel":
        fmt.Println("Wassap Daniel")
    case "Medhi":
        fmt.Println("Wassap Medhi")
    case "Jenny":
        fmt.Println("Wassap Jenny")
    default:
        fmt.Println("You have no friends?")
    }
}
