package main

import (
    "fmt"
)

func main() {
    for i := 8000; i < 8200; i++ {
        fmt.Printf("%d \t %b \t %#x \t %q\n", i, i, i, i)
    }
}
