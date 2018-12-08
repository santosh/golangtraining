package main

import (
    "fmt"
)

// 3 ways to create a map

func main() {
    //var greeting = make(map[string]string)
    var greeting map[string]string
    //greeting["Tim"] = "Good morning"
    //greeting["Jenny"] = "Bonjour"

    fmt.Println(greeting)

    var greeting_v2 := map[string]string{
        "Tim": "Good morning!",
        "Jenny": "Bonjour!"
    }

    delete(greeting_v2, "Jenny")

    fmt.Println(greeting_v2)
}
