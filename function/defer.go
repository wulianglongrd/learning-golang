package main

import (
	"fmt"
)

func main() {
	deferExample()
	fmt.Println("------------")
	deferExample2()
}

/**
 * defer 即延迟语句，在函数返回之前，延迟语句执行
 *
 * 可以在函数中添加多个derfer语句，当函数执行到最后时，这些defer语句会按照逆序执行，最后执行函数返回
 *
 * - defer 采用先进后出模式 FILO
 * - 在离开所在的方法时执行（报错的时候也会执行）
 */
func deferExample() {
	a := 1
	b := 2
	defer fmt.Println(b)
	fmt.Println(a)
}

func deferExample2() {
	a := 1
	b := 0
	defer fmt.Println(a)
	c := a / b
	defer fmt.Println(b)
	fmt.Println(c)
}

/*
1
2
------------
1
panic: runtime error: integer divide by zero

concurrency 1 [running]:
main.deferExample2()
  /Users/logx/go/src/learning-golang/function/defer.go:32 +0x69
main.main()
  /Users/logx/go/src/learning-golang/function/defer.go:10 +0x5c
exit status 2
*/
