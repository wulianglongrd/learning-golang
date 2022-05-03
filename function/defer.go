package main

import (
	"fmt"
)

func main() {
	//deferExample()
	deferExample2()
}

/**
 * defer 即延迟语句，在函数返回之前，延迟语句执行
 *
 * 可以在函数中添加多个defer语句，当函数执行到最后时，这些defer语句会按照逆序执行，最后执行函数返回
 *
 * - defer 采用先进后出模式 FILO
 * - 在离开所在的方法时执行（报错的时候也会执行）
 */
func deferExample() {
	// output:
	// a = 2
	// defer, a = 1

	a := 1
	defer fmt.Println("defer, a =", a) // 1
	a = a + 1
	fmt.Println("a =", a) // 2
}

func deferExample2() {
	// output:
	//  1
	//  panic: runtime error: integer divide by zero
	//
	//  goroutine 1 [running]:
	//  main.deferExample2()
	//  /Users/logx/go/src/github.com/wulianglongrd/learning-golang/function/defer.go:39 +0x88
	//  main.main()
	//  /Users/logx/go/src/github.com/wulianglongrd/learning-golang/function/defer.go:9 +0x17

	a := 1
	b := 0
	defer fmt.Println(a)
	c := a / b
	defer fmt.Println(b)
	fmt.Println(c)
}
