package main

import "fmt"

// slice 体积
func main() {
	noInit()
	makeInit()
	arrInit()
}

func noInit() {
	var s []int
	fmt.Println(s, len(s), cap(s)) // [] 0 0
}

func makeInit() {
	fmt.Println()
	s1 := make([]int, 3)
	fmt.Printf("s1 len: %d\n", len(s1)) // 3
	fmt.Printf("s1 cap: %d\n", cap(s1)) // 3
	fmt.Printf("s1 val: %d\n", s1)      // [0 0 0]

	fmt.Println()

	s2 := make([]int, 1, 3)
	fmt.Printf("s2 len: %d\n", len(s2)) // 1
	fmt.Printf("s2 cap: %d\n", cap(s2)) // 3
	fmt.Printf("s2 val: %d\n", s2)      // [0]
}

func arrInit() {
	fmt.Println()

	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s2 := s1[3:6]

	fmt.Println(s2) // [3 4 5]

	// s1[3:6] 是数学区间表示法，实际表示[3,6) 左闭右开（含3，不含6）
	// len: slice的区间长度，即窗口可见元素个数。len = 6 - 3
	fmt.Printf("s2 len: %d\n", len(s2)) // 3

	// cap: 透过这个窗口，最多可以看到的底层数组中元素的个数。窗口无法向左扩展。
	// cap = 底层数组的长度 - 切片的起始位值。cap = 8 - 3
	// 注意：数组的起始索引是0
	fmt.Printf("s2 cap: %d\n", cap(s2)) // 5
}
