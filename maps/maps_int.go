package main

import (
    "fmt"
)

func main() {
    greeting := map[int]string{
        0: "Good morning!",
        1: "Bonjour!",
        2: "Buenos dias!",
        3: "Bongiorno!",
    }

    fmt.Println(greeting)

    if val, exists := greeting[2]; exists {
        delete(greeting, 2)
        fmt.Println("val:", val)
        fmt.Println("exists:", exists)
    } else {
        fmt.Println("That value doesn't exist.")
        fmt.Println("val:", val)
        fmt.Println("exists:", exists)
    }

    fmt.Println(greeting)
}
