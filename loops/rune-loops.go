package main

import (
    "fmt"
)

func main() {
    for i := 40; i < 100; i++ {
        fmt.Println(i, "-", string(i), "-", []byte(string(i)))
    }
    // unlike strings, rune is defined with single quotes
    foo := 'a'
    fmt.Println(foo)
    fmt.Printf("%T\n", foo)
    fmt.Printf("%T\n", rune(foo))
}
