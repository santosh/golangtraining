package main

import (
    "os"
    "encoding/json"
)

type Person struct {
    First       string
    Last        string
    Age         int
    notExported int
}

func main() {
    p1 := Person{"James", "Bond", 20, 007}
    json.NewEncoder(os.Stdout).Encode(p1)
}
