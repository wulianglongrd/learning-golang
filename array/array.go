package main

import "fmt"

func main() {
	arrayDefine()
	arrayLenth()
	arrayLoop()
}

func arrayDefine() {
	// define array
	// var name [size]type
	var foo [3]int
	fmt.Println(foo)

	foo = [3]int{1, 2, 3}
	fmt.Println(foo)

	// define array with init value
	// name := [size]type{...}
	bar := [3]int{1, 2, 3}
	fmt.Println(bar)

	// define without size
	// name := [...]type{...}
	names := [...]string{"foo", "bar"}
	fmt.Println(names)

	// define with key
	f := [...]int{0: 1, 3: 2, 5: 3}
	fmt.Println(f)
}

func arrayLenth() {
	fmt.Println("---------------")

	var foo [3]int
	fmt.Println(len(foo)) // 3

	bar := [3]int{1, 2, 3}
	fmt.Println(len(bar)) // 3

	names := [...]string{"foo", "bar"}
	fmt.Println(len(names)) // 3

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
