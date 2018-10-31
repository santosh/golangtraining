package main

import (
    "fmt"
)

func main() {
    a := 43

    fmt.Println(a)
    fmt.Println(&a)

    var b *int = &a
    fmt.Println(b)

    // * fetches the value stored at the memory address
    // this is also called dereferencing
    fmt.Println(*b)
}
