package main

import "fmt"

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	// interface{} 表示空接口
	// interface{}(container) 把 container 转为接口类型
	// 以下语句：判断 container 是不是 []string 类型的切片
	if v, ok := interface{}(container).([]string); ok {
		fmt.Println("container is type []string")
		fmt.Printf("%+v", v)
		return
	}

	if v, ok := interface{}(container).(map[int]string); ok {
		fmt.Println("container is type map[int]string")
		fmt.Printf("%v", v)
		return
	}
}
