package main

import "fmt"

func main() {
	demo()
}

type Human struct {
	name string
	age  int
}

type Address struct {
	city, state string
}

type Teacher struct {
	Human      // 这种只有类型，没有字段名的字段称为匿名字段
	address    Address
	name       string
	speciality string
}

func demo() {
	a := *new(Address) // new 返回指标
	t := Teacher{Human{name: "innerName", age: 10}, a, "outerName", "sing"}
	// 因为有匿名字段存在，所以 speciality 初始化时不使用写 field，下面这种是错误的
	// t := Teacher{Human{name: "foo", age: 10}, speciality: "sing"}

	fmt.Println(t)            // {{innerName 10} { } outerName sing}
	fmt.Println(t.name)       // outerName
	fmt.Println(t.Human.name) // innerName, 匿名字段中的字段与非匿名字段名字相同，则最外层的优先访问，就近原则

	p2 := Address{
		city:  "bj",
		state: "xxx",
	}
	t.address = p2
	fmt.Println(t)     // {{innerName 10} {bj xxx} outerName sing}
	fmt.Println(t.age) // 访问匿名结构体的字段：提升字段
}
