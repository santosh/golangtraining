// write a function with one variadic parameter that finds the greatest number ina list of numbers

package main

import (
    "fmt"
)

func Largest(list ...int) int {
    var largest int
    for _, v := range list {
        if v > largest {
            largest = v
        }
    }
    return largest
}

func main() {
    greatest := Largest(1, 23, 4, 5, 6)
    fmt.Println(greatest)
}
