package main

import "fmt"

func main() {
	arrayDefine()
	arrayLength()
	arrayLoop()
}

func arrayDefine() {
	// 1. define array
	// var name [size]type
	var foo [3]int
	fmt.Println(foo) // [0 0 0]

	foo = [3]int{1, 2, 3}
	fmt.Println(foo) // [1 2 3]

	// 2. define array with init value
	// name := [size]type{...}
	bar := [3]int{1, 2, 3}
	fmt.Println(bar) // [1 2 3]

	// 3. define without size
	// name := [...]type{...}
	names := [...]string{"foo", "bar"}
	fmt.Println(names) // [foo bar]

	// 4. define with key
	f := [...]int{0: 1, 3: 2, 5: 3}
	fmt.Println(f) // [1 0 0 2 0 3]
}

func arrayLength() {
	fmt.Println("---------------")

	var foo [3]int
	fmt.Println(len(foo)) // 3

	bar := [3]int{1, 2, 3}
	fmt.Println(len(bar)) // 3

	names := [...]string{"foo", "bar"}
	fmt.Println(len(names)) // 2

	f := [...]int{0: 1, 3: 2, 5: 3}
	fmt.Println(len(f)) // 6
}

func arrayLoop() {
	fmt.Println("---------------")
	names := [...]string{"foo", "bar"}
	for i := 0; i < len(names); i++ {
		fmt.Printf("%d, %s\n", i, names[i])
	}

	fmt.Println("---------------")
	for k, v := range names {
		fmt.Printf("%d, %s\n", k, v)
	}
}
