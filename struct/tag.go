package main

import (
  "fmt"
  "reflect"
)

type Person struct {
  Name string `json:"name"`
}

func main()  {
  p := Person{"logx"}
  fmt.Println(p) // {logx}

  t := reflect.TypeOf(p)
  fmt.Println(t) // main.Person

  name := t.Field(0)
  fmt.Println(name) // {name main string json:"name" 0 [0] false}

  tag := name.Tag.Get("json")
  fmt.Println(tag) // name
}
