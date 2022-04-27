package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 4}
	s := a[1:3]

	// 数组属于值类型，同属于值类型的还有：基础数据类型、结构体类型
	fmt.Printf("array type: %T\n", a) // [5]int

	// 切片类型属于引用类型，同属于引用类型的还有: map、channel、func
	fmt.Printf("slice type: %T\n", s) // []int
}
