package main

import (
  "flag"
  "fmt"
)

//  go run flagString.go -name logx
func main()  {
  // 返回存储命令参数值的地址
  name := flag.String("name", "default value", "usage info")
  flag.Parse()

  fmt.Printf("%T\n", name) // *string
  fmt.Println("hello " + *name)
}
