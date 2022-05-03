package main

import "fmt"

func main() {
	variableSyntax()
	constantSyntax()
}

func variableSyntax() {
	// var name type
	var a = 1
	b := 2
	x, y := 100, "abc"
	fmt.Println(a, b, x, y)
}

func constantSyntax() {
	// const identifier [type] = value
	const a string = "aaa"
	const b = "bbb"
	fmt.Println(a, b)
}
