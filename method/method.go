package main

import (
  "fmt"
  "math"
)

func main()  {
  e := Employee{
    name: "andy",
    salary: 100,
    currency: "￥",
  }
  e.echoSalary()

  fmt.Println("\n---------------------")
  r := Rectangle{3, 4}
  c := Circle{2}
  fmt.Println("rectangle area is ", r.area())
  fmt.Println("circle area is ", c.area())
}
/**
// 方法用于给用户自定义类型添加新的行为
// 方法与函数的区别在于：方法有一个接收者，给函数添加一个接收者，它就变成了方法。
// 接收者可以是值接收者，也可以是指针接收者
func (t Type) methodName(param list) (return list) {
  ...
}
*/

type Employee struct {
  name string
  salary int
  currency string
}

// 为什么需要方法？
// go不是纯粹面向对象的语言，它不支持类。因此，类型的方法是一种实现类似于类的行为的方法
// 相同名称的方法可以在不同的类型上定义，而不能定义相同名称的函数。
func (e Employee) echoSalary()  {
  fmt.Printf("salary of %s is %s%d", e.name, e.currency, e.salary)
}

// ----------- 同名方法（同名函数是不允许的） ----------

type Rectangle struct {
  width, height float64
}

type Circle struct {
  radius float64
}

func (r Rectangle) area() float64 {
  return r.width * r.height
}

func (c Circle) area() float64 {
  return c.radius * c.radius * math.Pi
}
