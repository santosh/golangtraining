package main

import (
    "fmt"
)

type Person struct {
    first   string
    last    string
    age     int
}

func main() {
    p1 := Person{"James", "Bond", 20}
    p2 := Person{"Miss", "Moneypenny", 18}

    fmt.Println(p1.first, p1.last, p1.age)
    fmt.Println(p2.first, p2.last, p2.age)
}
