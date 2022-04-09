package main

import "fmt"

func main() {
	sliceCopy()
}

func sliceCopy() {
	arr := [...]int{1, 2, 3, 4, 5}
	o := arr[1:3]

	n := make([]int, 2)
	num := copy(n, o)
	
	fmt.Println(o) // [2 3]
	fmt.Println(n) // [2 3]
	fmt.Println(num) // 2
}
