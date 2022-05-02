package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 类型断言
// v, ok = x.(T)
// x 是一个接口类型的表达式
// T 是一个类型（断言类型）
// ok 返回断言是否成功
// v 如果断言成功，返回x的值（其类型为T）；如果断言失败，返回T类型的零值
// @see https://golang.google.cn/ref/spec#Type_assertions
func main() {
	//demo1()
	demo2()
}

func demo1() {
	var w io.Writer
	w = os.Stdout

	f := w.(*os.File)
	c := w.(*bytes.Buffer) // 如果只接收一个返回则，则在断言失败时会 panic

	fmt.Println(f)
	fmt.Println(c)
}

func demo2() {
	var w io.Writer
	w = os.Stdout

	//rw, ok := w.(io.ReadWriter)
	//fmt.Println(rw, ok)

	// 断言之后的结果再赋值给w，原有的值被新的值覆盖
	if w, ok := w.(io.ReadWriter); ok {
		fmt.Println(w)
	}
}
