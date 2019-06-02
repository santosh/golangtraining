package main

import (
	"fmt"
	"sort"
)

/*
 *Following challenges:
 *
 *type people []string
 *studyGroup := people{"Zeno", "John", "Al", "Jenny"}
 *
 *s := []string{"Zeno", "John", "Al", "Jenny"}
 *
 *n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
 *
 *Also sort above in reverse order
 */

// For people, we'll have to implement Len, Swap and Less
// as required by sort.Interface
type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func peopleResult() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}

func stringResult() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
	// now reverse
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}

func intResults() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)
	// now reverse
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}

func main() {
	// peopleResult()
	// stringResult()
	intResults()
}
