package main

import (
    "fmt"
)

// below is declaration
// var b string

// below is initialization
// var b string = "Santosh"

func main() {
    // these are the shorthand
    a := 10
    b := "golang"
    c := 4.17
    d := true

    // %v is default format, or say the actual value
    // See: https://godoc.org/fmt
    fmt.Printf("%v\n", a)
    fmt.Printf("%v\n", b)
    fmt.Printf("%v\n", c)
    fmt.Printf("%v\n", d)
}
