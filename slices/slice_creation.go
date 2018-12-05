package main

import (
    "fmt"
)

func main() {
    // there are potentially 3 ways to create a slice

    // using make
    // thi is the only one which you can access using index straight away
    student_slice_1 := make([]string, 35)


    // in rest, you have to use append
    // slice is init'ed with nil


    // using shorthand notation
    student_slice_2 := []string{}

    // using var keyword
    var student_slice_3 []string



    // when you pass only 2 parameter to make, it's both len and cap of the slice
    // when three is passed, 2nd one is len, and third cap
    student_slice_4 := make([]string, 35)
}


// slices are reference types
// actually slices, maps and channels all are reference types
// reference types == container data types, they basically hold some other data in themselves


// it's good to create reference type datatypes with make
