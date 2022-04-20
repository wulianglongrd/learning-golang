package main

import "fmt"

// 有可能只会输出 hello
// go不等待goroutine执行结果即开始执行后面的代码
// goroutine的任何返回值将被忽略
// 如果main goroutine终止了，则程序退出，不会等待其它goroutine执行结束
func main() {
	go hello()
	fmt.Println("main")
}

func hello() {
	fmt.Println("hello")
}
