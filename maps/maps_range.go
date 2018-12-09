package main

import (
    "fmt"
)

func main() {
    greeting := map[int]string{
        0: "Good morning!",
        1: "Bonjour!",
        2: "Buenos dias!",
        3: "Bongiorno!",
    }

    // uses for range to iterate over the map
    for key, val := range greeting {
        fmt.Println(key, " - ", val)
    }
}
