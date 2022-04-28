package main

import "fmt"

// https://go.dev/ref/spec#Struct_types
// https://stackoverflow.com/questions/45122905/how-do-struct-and-struct-work-in-go
// struct{} and struct{}{}
// can do for https://geektutu.com/post/hpg-empty-struct.html
func main() {
	withField()

	fmt.Println()

	noField()
}

func withField() {
	// 1. 先定义，后初始化
	type person struct {
		name string
	}
	p := person{
		name: "logx",
	}
	fmt.Println(p)

	// 2. 定义同时初始化
	p2 := struct {
		name string
	}{
		name: "logx",
	}
	fmt.Println(p2)
}

// struct{}   是一个没有字段的struct类型
// struct{}{} 是一个没有字段的struct类型变量(composite literal 复合字面量)
//
// struct{} {}
// |  ^    | ^
//   type    empty element list
func noField() {
	// 1. 先定义，后初始化
	type person struct{} // an empty struct
	p := person{}
	fmt.Println(p)

	// 2. 定义同时初始化
	p2 := struct{}{} // struct{} 相当于p声明处的 person
	fmt.Println(p2)
}
