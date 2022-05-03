package main

import (
  "fmt"
  "runtime"
)

func main()  {
  fmt.Println("GOROOT: ", runtime.GOROOT())
  fmt.Println("num cpu: ", runtime.NumCPU())

  fmt.Println("---------------")

  goSched()
}

func goSched()  {
  go func() {
    for i:=0; i < 3; i++ {
      fmt.Println("goroutine...")
    }
  }()

  for i := 0; i < 2; i++ {
    runtime.Gosched()
    fmt.Println("main...")
  }
}
