package main

import "fmt"

func main() {
	sliceDefine()
	sliceFromArray()
  sliceMakeAndNew()
}

func sliceDefine() {
	// var name []type = make([]type, len)
	var s1 []int = make([]int, 3)
	fmt.Println(s1)

	// simple with make create slice
	// name := make([]type, len)
	s2 := make([]int, 3)
	fmt.Println(s2)

	// var identifier []type
	foo := []int{1, 2, 3}
	fmt.Println(foo)
}

func sliceFromArray() {
	arr := [5]int{1, 2, 3, 4, 5}

	s1 := arr[1:3]
	s2 := arr[2:4]
	fmt.Println(s1) // [2 3]
	fmt.Println(s2) // [3 4]

	// slice和数组共享数据，slice和数组修改都会相互反应
	s1[0] = 100
	arr[2] = 200
	fmt.Println(s1) // [100 200]
	fmt.Println(arr) // [1 100 200 4 5]
}

// New 返回指针地址 more @see func description
// make 返回第一个元素，可预设内存空间，避免未来的内存拷贝 more @see func description
func sliceMakeAndNew()  {
  s1 := new([]int)
  fmt.Println(s1) // &[]

  s2 := make([]int, 3)
  fmt.Println(s2) // [0 0 0]
}
