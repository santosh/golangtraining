package main

import (
    "fmt"
)

func main() {
    mySlice := []string{"a", "b", "c", "g", "m", "z"}
    fmt.Println(mySlice)
    fmt.Println(mySlice[2:4])  // slicing a slice
    fmt.Println(mySlice[2])   // index access
    fmt.Println("myString"[2])  // index access
}
