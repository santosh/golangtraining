package main

import (
    "fmt"
)

func main() {
    sa := make([]string, 4, 8)
    fmt.Println(len(sa))
    fmt.Println(cap(sa))
}
