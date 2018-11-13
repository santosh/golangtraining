package main

import (
    "fmt"
)

// code from previous lesson with func expression


func main() {
    HalfAndIsBool := func(num int) (half int, even bool) {
    // there are two returns
    // first is half, which returns param/2
    half = num / 2
    // the second returns a bool which says if param is even
    even = num % 2 == 0

    return 
}
    fmt.Println(HalfAndIsBool(4))
}
