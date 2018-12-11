package main

import (
    "fmt"
    "bufio"
    "strings"
    "os"
)

func main() {
    const input = "The power went off. It just came right now."

    scanner := bufio.NewScanner(strings.NewReader(input))

    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading input: ", err)
    }
}
