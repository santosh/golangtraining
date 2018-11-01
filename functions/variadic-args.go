package main

import (
    "fmt"
)

func main() {
    data := []float64{43, 56, 87, 12, 45, 57}
    n := average(data)  // this is how we flatten slices
    fmt.Println(n)
}

func average(sf []float64) float64 {  // ... is like *sf in Python
    fmt.Println(sf)
    fmt.Printf("%T\n", sf)
    total := 0.0
    for _, v := range sf {  // range will always give index and value
        total += v
    }

    return total / float64(len(sf))
}
