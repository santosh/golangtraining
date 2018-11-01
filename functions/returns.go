package main

import (
    "fmt"
)

func main() {
    fmt.Println(greet("Jane ", "Doe"))
}

func greet(fname, lname string) (x, y string) {
    x = fmt.Sprint(fname, lname)
    y = fmt.Sprint(lname, fname)
    return
}
