package main

import "fmt"

func main() {
	//keyPanic()
	nilPanic()
}

func keyPanic() {
	// panic: runtime error: hash of unhashable type []int
	m := map[interface{}]int{
		"1":      1,
		[]int{2}: 2, // []int{2} 切片无法计算hash值，引发panic
		3:        3,
	}
	fmt.Println(m)
}

func nilPanic() {
	var m map[string]int
	fmt.Println(m) // map[]

	v, _ := m["foo"]
	fmt.Println(v) // 0

	delete(m, "foo") // no panic

	// panic: assignment to entry in nil map
	// m 未初始化
	m["foo"] = 111
}
