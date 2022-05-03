package main

import "fmt"

func main() {
	block1()
	block2()
}

func block1() {
	name := "outer"
	{
		name := "inner"
		fmt.Println("block inner: ", name) // inner
	}
	fmt.Println("block outer: ", name) // outer
}

func block2() {
	foo := "str"
	{
		foo := 123
		fmt.Println(foo) // 123
	}
	fmt.Println(foo) // str
}
