package main

import (
    "fmt"
)

func main() {
    for i := 50; i < 140; i++ {
        fmt.Println(i, "-", string(i), "-", []byte(string(i)))
    }
}
