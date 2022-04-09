package main

import "fmt"

// len 反应的是切片的长度
// cap 反应的是切版的底层数组从“切片的索引”开始的底层数组的长度
func main() {
	makeSlice()
	arrSplit()
}

func makeSlice() {
	var s = make([]int, 3, 5)
	fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 5

	var s2 []int
	fmt.Println(s2, len(s2), cap(s2)) // [] 0 0
}

func arrSplit() {
	arr := [...]int{1, 2, 3, 4, 5}

	s := arr[1:3] // 1,3 为数组 arr的下标，1含 3不含
	fmt.Println(s, len(s), cap(s)) // [2 3] 2 4
}
