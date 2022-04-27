package main

import (
  "flag"
  "fmt"
)

var name string

func init()  {
  flag.StringVar(&name, "name", "the default value", "usage info")
}

// @see https://golang.google.cn/pkg/flag/
// go run stringVar.go -h
// go run stringVar.go -name logx
// go run stringVar.go -name=logx
func main()  {
  flag.Parse()
  fmt.Println("hello " + name)
}
