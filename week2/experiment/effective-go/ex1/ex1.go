package main

import (
	"fmt"
	"sort"
)

// TODO: Create a type 'Person' with fields: Name (string) and Age (int)
type Person struct {
	Name string
	Age  int
}

type People []Person

// TODO: Implement the three methods required by sort.Interface:
// - Len() int
// - Less(i, j int) bool  (sort by age)
// - Swap(i, j int)

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	people := People{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	// TODO: Sort the people slice and print it
	sort.Sort(people)
	for _, person := range people {
		fmt.Println(person)
	}

	// sort.IntSlice attaches the sort.Interface methods to the []int slice
	// so that it can be used in sort.Sort
	nums := []int{5, 4, 3, 2, 1, 7, 8}
	sort.Sort(sort.IntSlice(nums))
	fmt.Println(nums)
}
