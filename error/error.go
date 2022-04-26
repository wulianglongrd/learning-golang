package main

import (
	"fmt"
	"os"
)

// error and panic
// 错误：可能出现问题的地方出现了问题，比如打开文件失败 error
// 异常：不应该出现问题的地方出现了问题，比如引用了空指针、下标越界、除数为0、不该出现的分支 panic
// go中的错误是一种内置类型，使用error表示

func main() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println("-----------")
		fmt.Println(err)
		fmt.Println("-----------")
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
