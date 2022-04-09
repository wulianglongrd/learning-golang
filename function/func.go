package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 + 2 =", sum(1, 2))
	fmt.Println("1 + 2 =", sameParamTypeSum(1, 2))

	ok, val := muliReturn(1, 1)
	fmt.Println(ok, val)
}

/*
  函数基本语法
  func funcName(paramName1 type, paramName2 type) (outputType1, outputType2) {
    return value1, value2
  }
*/

func sum(a int, b int) int {
	return a + b
}

func sameParamTypeSum(a, b int) int {
	return a + b
}

func muliReturn(a, b int) (bool, int) {
	return true, a + b
}
