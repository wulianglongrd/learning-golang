package main

import (
	"fmt"
)

func main() {
	a, b := swap("foo", "bar")
	fmt.Println(a, b)

	x, _ := swap("x1", "x2")
	fmt.Println(x)
}

func swap(x, y string) (string, string) {
	return y, x
}
