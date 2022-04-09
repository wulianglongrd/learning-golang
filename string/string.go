package main

import (
  "fmt"
)

func main() {
  stringBasic()
}

func stringBasic() {
  str := "Hello world!"
  for i := 0; i < len(str); i++ {
    fmt.Printf("%d = %c\n", str[i], str[i])
  }
}
