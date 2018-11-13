// anonymous self executing function
package main

import (
    "fmt"
)

func main() {
    func() {
        fmt.Println("I'm anonymous.")
    }()  // note the paren? we are executing it on the go
    // and it doesn't needs a name
}
