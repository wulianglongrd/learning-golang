package test

import (
	"strings"
)

func Add(a, b int) int {
	return a + b
}

func Split(s string) []string {
	s = strings.ReplaceAll(s, "，", ",")
	a := strings.Split(s, ",")
	return a
}

// FactorialForLoop n! = n * (n-1) * (n-2) .. * 1
func FactorialForLoop(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	val := 1
	for i := n; i > 1; i-- {
		val *= i
	}
	return val
}

func FactorialRecursion(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * FactorialRecursion(n-1)
}
