package main

import (
    "fmt"
)

func main() {
    mySlice := make([]int, 0, 6)

    fmt.Println("==============")
    fmt.Println(mySlice)
    fmt.Println(len(mySlice))
    fmt.Println(cap(mySlice))
    fmt.Println("==============")

    for i := 0; i < 5000; i++{
        mySlice = append(mySlice, i)
        if i > 3000 {
            fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice[i])

        }
    }
}

// shall I assume
