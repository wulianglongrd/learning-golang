package main

import (
	"fmt"
	util "github.com/wulianglongrd/learning-golang/common"
)

/*
  函数基本语法
  func funcName(paramName1 type, paramName2 type) (outputType1, outputType2) {
    return value1, value2
  }
*/
func main() {
	fmt.Println("1 + 2 =", util.Sum(1, 2))
	fmt.Println("1 + 2 =", util.SameParamTypeSum(1, 2))

	ok, val := util.MultiReturnSum(1, 1)
	fmt.Println(ok, val)
}
