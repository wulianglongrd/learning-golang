package main

import (
	"fmt"
)

func main() {
	structDefine()
	structAccess()
}

/*
结构体定义
type struct_type struct {
	member definition
	...
	member definition
}
结构体声明
var_name := struct_type { value1, value2... }
*/

type Student struct {
	name string
	age  int
}

func structDefine() {
	// 按顺序初史化
	s := Student{"andy", 10}
	fmt.Println(s) // {andy 10}

	// 通过 file: value 初史化
	s2 := Student{name: "lily", age: 18}
	fmt.Println(s2) // {lily 18}

	fmt.Println("--------")

	// new的方式
	s3 := new(Student)
	fmt.Println(s3) // &{ 0}
	s3.name = "sam"
	s3.age = 20
	fmt.Println(s3) // &{sam 20}
}

func structAccess() {
	s := Student{name: "andy", age: 10}
	fmt.Println(s.name, s.age)
}
