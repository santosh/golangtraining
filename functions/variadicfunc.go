package main

import (
    "fmt"
)

func main() {
    n := average(43, 56, 87, 12, 45, 57)
    fmt.Println(n)
}

func average(sf ...float64) float64 {  // ... is like *sf in Python
    fmt.Println(sf)
    fmt.Printf("%T\n", sf)
    total := 0.0
    for _, v := range sf {  // range will always give index and value
        total += v
    }

    return total / float64(len(sf))
}
