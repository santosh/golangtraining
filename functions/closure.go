package main

import (
    "fmt"
)

func main() {
    x := 42
    {
        fmt.Println(x)
        y := "The credit belongs with the one who is in the ring."
        fmt.Println(y)
    }
}

// maybe closure is all about scopes?
// at least that is what I think after watching the video
