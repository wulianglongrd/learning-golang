package main

import "fmt"

// 类型转换
// https://golang.google.cn/ref/spec#Conversions
// expression: T(x)
// T is a type
// x is an expression that can be converted to type T
func main() {
	//numberConversion()
	//channelConversion()
	stringConversion()
}

func numberConversion() {
	var i int = 42
	var f float64 = float64(i)
	fmt.Println(f)

	var u uint = uint(f)
	fmt.Println(u)
}

func channelConversion() {
	ch1 := make(chan int)
	ch2 := (<-chan int)(ch1)
	//ch3 := <-chan int(ch1) // fatal error: all goroutines are asleep - deadlock!

	fmt.Printf("%T\n", ch2) // <-chan int
}

func stringConversion() {
	a := []byte{'h', 'e', 'l', 'l', 'o'}
	s := string(a)
	fmt.Printf("%T, %s\n", s, s) // string, hello

	x := []byte(s)
	fmt.Printf("%T, %v\n", x, x) // []uint8, [104 101 108 108 111]
}
