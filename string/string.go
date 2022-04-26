package main

import (
  "fmt"
  "unicode/utf8"
)

func main() {
  stringBasic()
  strLen()
}

// 长文本可使用 ``, 内含双引号不需要转义
func stringBasic() {
  str := "Hello world!"
  for i := 0; i < len(str); i++ {
    fmt.Printf("%d = %c\n", str[i], str[i])
  }
}

// 字节长度：和编码无关，用 len(str) 获取
// strings.Xxx
func strLen()  {
  s := "你好"
  fmt.Println(len(s)) // 6
  fmt.Println(utf8.RuneCountInString(s)) // 2
  fmt.Println(utf8.RuneCountInString("你好ab")) // 4
}
