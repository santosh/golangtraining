// In simple words, the thing that points towards a memory location is known as pointer

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
    fmt.Println(*b)
}
