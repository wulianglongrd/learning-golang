package main

import "fmt"

func main()  {
  r := Rectangle2{1, 2}
  fmt.Println(r)
  r.setWidth()
  fmt.Println(r) // {100 2}
  r.setHeight()
  fmt.Println(r) // {100 2}
}

type Rectangle2 struct {
  width, height float64
}

func (r *Rectangle2) setWidth()  {
  r.width = 100
}

// r只是一个copy，不能修改接收者中的数据
func (r Rectangle2) setHeight()  {
  r.height = 200
}
