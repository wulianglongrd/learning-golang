package main

import (
  "fmt"
  "reflect"
)

type Person struct {
  name string
  age int
}

func (p *Person) Hello()  {
  fmt.Println("hello " + p.name)
}


func main()  {
  m := make(map[string]string)
  m["foo"] = "foo_v"
  m["bar"] = "bar_v"

  t := reflect.TypeOf(m)
  fmt.Println("type: ", t)
  fmt.Println("name: ", t.Name())
  fmt.Println("kind: ", t.Kind().String())
  fmt.Println("string: ", t.String())

  fmt.Println("----------")

  v := reflect.ValueOf(m)
  fmt.Println("value: ", v)

  fmt.Println("-----------------")

  p := &Person{name: "logx", age: 10}
  pVal := reflect.ValueOf(p)

  //for i:= 0; i < pVal.NumField(); i++ {
  //  fmt.Printf("field %d: %v\n", i, pVal.Field(i))
  //}

  method := pVal.MethodByName("Hello")
  fmt.Println(method)
  method.Call([]reflect.Value{})
}
