package main

import (
	"fmt"
)

func main() {
	pointerDefine()
	nilVal()

	arr := [...]int{1, 2, 3}
	modifyArray(&arr)
	fmt.Println(arr) // [1 111 3]

	fmt.Println("--------------")

	modifySlice(arr[:])
	fmt.Println(arr) // [1 222 3]
}

// 指针是存储另一个变量的内存地址的变量
// & 取址
// * 取值
func pointerDefine() {
	// var name *type

	n := 1
	var i *int
	i = &n

	fmt.Println(i, *i) // 0xc0000b0008 1
}

func nilVal() {
	var f *float32
	fmt.Println(f) // <nil> 空指针
}

// bad 修改数组，不要传递数组指标，使用切片
func modifyArray(arr *[3]int) {
	arr[1] = 111
}

// good
func modifySlice(s []int) {
	s[1] = 222
}
