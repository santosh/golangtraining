package main

import (
    "fmt"
)

type Contact struct {
    greeting    string
    name        string
}

func SwitchOnType(x interface{}) {
    switch x.(type) {
    case int:
        fmt.Println("int")
    case string:
        fmt.Println("string")
    case Contact:
        fmt.Println("Contact")
    default:
        fmt.Println("unknown")
    }
}

func main() {
    SwitchOnType(7)
    SwitchOnType("McLeod")
    var t = Contact{"Good to see you,", "Tim"}
    SwitchOnType(t)
    SwitchOnType(t.greeting)
    SwitchOnType(t.name)
}
