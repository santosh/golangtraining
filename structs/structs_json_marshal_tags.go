package main

import (
    "fmt"
    "encoding/json"
)

type Person struct {
    First       string
    Last        string `json:"-"`
    Age         int    `json:"wisdom score"`
}

func main() {
    p1 := Person{"James", "Bond", 22}
    bs, _ := json.Marshal(p1)
    fmt.Println(string(bs))
}
