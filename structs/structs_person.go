package main

import (
    "fmt"
)

type person struct {
    first   string
    last    string
    age     int
}

func (p person) fullName() string {
    return p.first + " " + p.last
}

func main() {
    p1 := person{"Santosh", "Kumar", 22}
    p2 := person{"Jane", "Doe", 28}
    fmt.Println(p1.fullName())
    fmt.Println(p2.fullName())
}
