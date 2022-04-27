package main

import "fmt"

func main() {
	valueType()
	sizeType()
}

func valueType() {
	a := [...]string{"foo", "bar"}
	b := a

	b[0] = "xxx" // go中的数组是值类型，而不是引用类型，b修改不会返应到a

	fmt.Println(a)
	fmt.Println(b)
}

// 数组的大小是类型的一部份
func sizeType() {
	a := [3]int{1, 2, 3}
	var b [4]int

	fmt.Printf("a type is %T\n", a) // [3]int
	fmt.Printf("b type is %T\n", b) // [4]int

	// b = a // cannot use a (variable of type [3]int) as type [4]int in assignment
}
