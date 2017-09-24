// Use https://godoc.org/sort to sort the following:

// (1)
// type people []string
// studyGroup := people{"Zeno", "John", "Al", "Jenny"}

// (2)
// s := []string{"Zeno", "John", "Al", "Jenny"}

// (3)
// n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

// Also sort the above in reverse order
package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Less(i, j int) bool { return p[i] < p[j] }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {

	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	// (2)
	s := []string{"Zeno", "John", "Al", "Jenny"}
	// (3)
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println("unsorted: ", studyGroup)
	sort.Sort(studyGroup)
	fmt.Println("sorted: ", studyGroup)
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println("reversed: ", studyGroup)

	fmt.Println("===================================")

	fmt.Println("unsorted: ", s)
	sort.Strings(s)
	// alternatives:
	// sort.StringSlice(s).Sort() // method attached to Interface
	// sort.Sort(sort.StringSlice(s)) // conversion to Interface
	fmt.Println("sorted: ", s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("reversed: ", s)

	fmt.Println("===================================")

	fmt.Println("unsorted: ", n)
	sort.Ints(n)
	// alternative:
	// sort.Sort(sort.IntSlice(n)) // conversion
	fmt.Println("sorted: ", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("reversed: ", n)
}
