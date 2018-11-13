package main

import (
    "fmt"
)

// pass by value and pass by reference

func main() {
    age := 44
    fmt.Println(&age)

    changeMe(&age)

    fmt.Println(&age)
    fmt.Println(age)
}

func changeMe(z *int) {
    fmt.Println(z)
    fmt.Println(*z)

    *z = 24

    fmt.Println(z)
    fmt.Println(*z)
}
