package main

import (
	"fmt"
	"github.com/wulianglongrd/learning-golang/common"
	//  . "fmt"  // 调用包内的函数时，不再需要包名前缀
	// bar "foo" // 取别名，bar是别名
	// _ "github.com/ziutek/mymysql/godrv" // 匿名导致：导入包时只会通过init初始化，不能调用包内的函数和常量
)

func main() {
	sum := util.Sum(1, 2)
	fmt.Println(sum)
}
