package main

import "fmt"

// 接口：定义对象的行为，表示对象应该做什么。
// 在go中，接口是一组方法签名。`接口`定义了`类型`应该具有的方法，`类型`决定了如何实现这些`方法`
// 当`类型`为`接口`中的所有方法提供定义时，它被称为实现接口

/*
type interface_name interface {
  method_name1 [return_type]
  method_name2 [return_type]
  ...
}
*/

type Phone interface {
	call()
}

// -------------

type NokiaPhone struct {
}

func (n NokiaPhone) call() {
	fmt.Println("i am nokia, i can call you")
}

// -------------

type IPhone struct {
}

func (i IPhone) call() {
	fmt.Println("i am iPhone, i can call you")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
