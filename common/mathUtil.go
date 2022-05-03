package util

import "fmt"

func init() {
	fmt.Println("mathUtil init 1")
}

func init() {
	fmt.Println("mathUtil init 2")
}

func Sum(a, b int) int {
	return a + b
}

func SameParamTypeSum(a, b int) int {
	return a + b
}

func MultiReturnSum(a, b int) (int, bool) {
	return a + b, true
}
