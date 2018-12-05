package main

import (
    "fmt"
)

func main() {
    mySlice := []int{0, 6}

    fmt.Println("==============")
    fmt.Println(mySlice)
    fmt.Println(len(mySlice))
    fmt.Println(cap(mySlice))
    fmt.Println("==============")

    for i := 0; i < 60; i++{
        mySlice = append(mySlice, i)
        fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice[i])
    }
}

// shall I assume
