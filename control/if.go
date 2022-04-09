package main

import "fmt"

func main() {
	basicIf()
	withStatement()
}

func basicIf() {
	// if condition {
	//
	// } else if condition {
	//
	// } else {
	//
	// }

	a := 10
	if a >= 1 {
		fmt.Println("a > 1")
	} else {
		fmt.Println("a < 1")
	}
}

func withStatement() {
	// if statement; condition {
	//
	// } else {
	//
	// }

	if num := 1; num % 2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}