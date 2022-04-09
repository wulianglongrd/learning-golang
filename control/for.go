package main

import "fmt"

func main() {
	basicFor()
	fmt.Println("------------")
	basicFor2()
	fmt.Println("------------")
	forRange()
}

func basicFor() {
	// for init; condition; post {}
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func basicFor2() {
	// for condition {}
	a := 0
	b := 3
	for a < b {
		fmt.Println(a)
		a++
	}
}

func forRange() {
	numbers := [3]int{1, 2, 3}
	for k, v := range numbers {
		fmt.Printf("key=%d, value=%d\n", k, v)
	}

	for v := range numbers {
		fmt.Printf("value=%d\n", v)
	}
}
