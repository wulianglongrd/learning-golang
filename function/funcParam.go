package main

import (
	"fmt"
)

func main() {
	echoArgs(1, 2)
	fmt.Println("-----------")
	echoArgs(1, 2, 3)

	fmt.Println("==========")

	a1 := 1
	a2 := add(&a1)
	fmt.Printf("a1 = %d, a1 = %d\n", a1, a2) // a1的值也发生变化
}

// 可变参数
func echoArgs(arg ...int) {
	for _, n := range arg {
		fmt.Println(n)
	}
}

// 指针传递
// 指针使得多个函数有操作同一个对象
// 传指针比较轻量级(8bytes)，只是传内存地址，我们可以用指针传递体积大的结构体。
// 如果使用值传递，在每次copy上面需要花费较多的系统开销（内存和时间）
// 所以，当需要传递大的结构体的时候，用指针是一个明智的选择
// ** slice、map实现机制类似指针，可以直接传递，而不用取地址后传递指针。若函数需要改变slice的长度，则仍需要使用指针传递。
func add(a *int) int {
	*a = *a + 1
	return *a
}
