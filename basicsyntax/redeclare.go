package main

import "fmt"

// 变量重申明：已经申明过的变量再次申明
// 1. 重申明的变量类型应该相同
// 2. 只能在短变量时才能重声明，否则无法通过编译
// 3. "声明并赋值"的变量必须是多个，并且其中至少有一个是新的变量
func main() {
	var name string
	name, age := hello() // name 变量重申明
	fmt.Println(name, age)

	//name := "foo" // 无法通过编译
}

func hello() (string, int) {
	return "logx", 10
}
