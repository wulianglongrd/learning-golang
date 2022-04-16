package main

import "fmt"

// 定义新类型
// type 新类型名 Type

// 定义别名
// type 别名 = Type

type myint int    // 定义新的类型
type myint2 = int // 不是定义新的类型，只是别名

func main() {
	var i1 myint
	var i2 = 100
	i1 = 100
	//i1 = i2 // Cannot use 'i2' (type int) as the type myint
	fmt.Println(i1, i2)

	var i3 myint2
	i3 = i2
	fmt.Println(i3)
}
