package main

import (
	"fmt"
	"learning-golang/util"
)

// golang两个保留函数： main 和 init
// init 在包导致时执行，init在main函数执行前执行
// 一个go文件内可以有多个init函数，同一个文件内的init函数从上到下依次执行
// 一个go文件可以被其它包多次import，但init只会执行一次
// 多个go文件的init函数，按照go文件名的字符串从小到大依次执行
// 相互依赖的，被依赖的先执行
// @see https://www.qfgolang.com/?special=baoguanli&pid=2038
func main() {
	fmt.Println(util.Sum(1, 2))
}
