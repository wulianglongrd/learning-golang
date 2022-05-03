package main

import "fmt"

// 安全断言
// <目标类型值>, <布尔参数> ：= <表达式>.(目标类型)

// 非安全断言
// <目标类型值> := <表达式>.(目标类型)

type student struct {
	name string
	age  int32
}

func main() {
	//var i1 interface{} = new(student)
	//s := i1.(student) // 不安全，如果断言失败，会直接panic
	//fmt.Println(s)

	var i2 interface{} = new(student)
	s, ok := i2.(student) // 安全，如果断言失败, ok 为false，不会panic
	if ok {
		fmt.Println(s)
	}
}
