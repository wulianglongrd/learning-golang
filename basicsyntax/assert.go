package main

import "fmt"

// 断言： x.(T)
// x 代表要被判断类型的值，必须是接口类型
// @see https://golang.google.cn/ref/spec#Type_assertions
func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	// interface{} 表示空接口
	// interface{}(container) 把 container 转为接口类型
	// 以下语句：判断 container 是不是 []string 类型的切片
	value, _ := interface{}(container).([]string)
	fmt.Println(value)
}
