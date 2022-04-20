package main

import "fmt"

// 类型转换
func main()  {
  var i int = 42
  var f float64 = float64(i)
  fmt.Println(f)
  var u uint = uint(f)
  fmt.Println(u)
}
