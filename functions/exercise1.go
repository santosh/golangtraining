package main

import (
    "fmt"
)

func HalfAndIsBool(num int) (half int, even bool) {
    // there are two returns
    // first is half, which returns param/2
    half = num / 2
    // the second returns a bool which says if param is even
    even = num % 2 == 0

    return 
}

func main() {
    fmt.Println(HalfAndIsBool(4))
}
