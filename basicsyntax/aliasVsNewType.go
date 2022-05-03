package main

import "fmt"

// 定义新类型
// type 新类型名 Type

// 定义别名
// type 别名 = Type
func main() {
	//alias()
	newType()
}

// 定义新的类型 newIntType
func newType() {
	type newIntType int

	var x int
	var y newIntType

	x = 100
	//y = x // Cannot use 'x' (type int) as the type newIntType
	fmt.Println(x, y) // 100 0
}

// 定义别名
// 别名之间可以相互赋值
func alias() {
	type myInt = int

	var x int
	var y myInt

	x = 100
	y = x
	fmt.Println(x, y) // 100 100
}
