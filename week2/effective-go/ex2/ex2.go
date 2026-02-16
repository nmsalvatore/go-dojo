package main

import (
	"fmt"
	"sort"
)

type IntSequence []int

// TODO: Implement a String() method that:
// 1. Makes a copy of the sequence
// 2. Sorts it using sort.IntSlice
// 3. Returns it as a formatted string like "[1 2 3]"

func (i IntSequence) String() string {
	s := append([]int{}, i...)
	sort.Ints(s)
	return fmt.Sprint(s)
}

func main() {
	seq := IntSequence{5, 2, 8, 1, 9}
	fmt.Println(seq) // Should print sorted sequence
}
