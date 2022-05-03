package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		// areaError 类型实现了 error 接口，可以作为 error 返回
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -1.0
	_, err := circleArea(radius)
	if err != nil {
		// 使用断言，如果断言成功，则可使用断言的结果获取更多信息
		if err, ok := err.(*areaError); ok {
			fmt.Printf("radius %0.2f is less than zero\n", err.radius) // radius -1.00 is less than zero
			fmt.Println(err.err)                                       // radius is negative
			return
		}

		fmt.Println(err)
	}
}
