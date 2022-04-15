package main

import "fmt"

func main() {
	canCompare()
	canNotCompare()
}

type name struct {
	first, last string
}

// 结构值是值类型
// 如果每个字段都具有可比性，则是可比较的。如果他们对应的字段都相等，则认为两个结构体变量是相等的。
func canCompare() {
	n1 := name{"long", "andy"}
	n2 := name{"long", "andy"}
	fmt.Println(n1 == n2) // true

	n3 := name{"long", "andy"}
	n4 := name{}
	n4.first = "long"

	fmt.Println(n3 == n4) // false
}

type image struct {
	data map[int]int
}

func canNotCompare() {
	image1 := image{data: map[int]int{
		0: 100,
	}}
	image2 := image{data: map[int]int{
		0: 100,
	}}

	fmt.Println(image1)
	fmt.Println(image2)

	// error : 不能比较
	// fmt.Println(image1 == image2)
}
