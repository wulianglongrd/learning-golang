package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main()  {
  ch := make(chan int)

  go func() {
    for {
      time.Sleep(time.Second)
      v := rand.Intn(10)
      ch <- v
      fmt.Println("producer: ", v)
    }
  }()

  for {
    v := <- ch
    fmt.Println("consume: ", v)
  }
}
