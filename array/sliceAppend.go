package main

import "fmt"

func main() {
	appendNotFull()
	fmt.Println("---------")
	appendFull()
}

func appendNotFull() {
	arr := [...]int{1, 2, 3, 4, 5}
	s := arr[1:3]

	fmt.Println(s) // [2 3]

	// append函数会改变slice所引用的数组的内容
	s = append(s, 400)
	fmt.Println(s) // [2 3 400]
	fmt.Println(arr) // [1 2 3 400 5]
}

func appendFull() {
	arr := [...]int{1, 2, 3}
	s := arr[1:3]

	fmt.Println(s)

	// 当slice中没有剩余空间时，将动态分配新的数组空间，append返回的数组指向新分配的空间，原数组的内容保持不变
	s = append(s, 400)
	fmt.Println(s) // [2 3 400]
	fmt.Println(arr) // [1 2 3]
}


