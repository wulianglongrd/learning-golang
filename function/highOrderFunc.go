package main

import (
	"fmt"
)

func main() {
	anonymousFunc()

	fmt.Println("=======================")

	res := callbackFunc(1, 2, sumInt)
	fmt.Println(res)

	fmt.Println("=======================")

	closureFunc()
}

// 匿名函数
func anonymousFunc() {
	func() {
		fmt.Println("1. 只能执行一次的匿名函数")
	}()

	fmt.Println("------------------------")

	func(a, b int) {
		fmt.Println("2. 自带参数匿名函数", a, b)
	}(1, 2)

	fmt.Println("------------------------")

	fun := func() {
		fmt.Println("3. 能执行多次的匿名函数")
	}
	fun()
	fun()
	fmt.Println(fun)
}

// 回调函数
//
func callbackFunc(a, b int, fun func(int, int) int) int {
	return fun(a, b)
}

// 闭包函数
func closureFunc() {
	res1 := incr()
	fmt.Printf("%T\n", res1)
	fmt.Println(res1)

	v1 := res1()
	v2 := res1()
	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println("----")

	res2 := incr()
	fmt.Println(res2())
	fmt.Println(res2())
	fmt.Println(res2())
}

// 函数返回一个 func() int 签名的函数
func incr() func() int {
	i := 0
	fun := func() int {
		i++
		return i
	}
	return fun
}

func sumInt(a, b int) int {
	return a + b
}
