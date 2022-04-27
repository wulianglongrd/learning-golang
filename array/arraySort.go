package main

import (
	"fmt"
	"sort"
)

// https://yourbasic.org/golang/how-to-sort-in-go/
func main() {
	//basicSort()
	customizedSort()
}

func basicSort() {
	s := []int{4, 2, 3, 1}
	sort.Ints(s)
	fmt.Println(s) // [1 2 3 4]
}

func customizedSort() {
	type person struct {
		name string
		age  int
	}

	family := []person{
		{"alice", 23},
		{"david", 2},
		{"eve", 2},
		{"bob", 25},
	}

	sort.SliceStable(family, func(i, j int) bool {
		return family[i].age < family[j].age
	})

	fmt.Println(family) // [{david 2} {eve 2} {alice 23} {bob 25}]
}
