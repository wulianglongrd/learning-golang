package main

import (
  "encoding/json"
  "fmt"
)

type Human struct {
  firstName string
  age int
}

func main()  {
  h := Human{firstName: "logx", age: 10}
  fmt.Println(h)
  fmt.Println(toStr(h))
}

func toJson(s string) Human {
  h := Human{}
  err := json.Unmarshal([]byte(s), &h)
  if err != nil {
    println(err)
  }
  return h
}

func toStr(h Human) string {
  res, err := json.Marshal(&h)
  if err != nil {
    println(err)
  }
  return string(res)
}
