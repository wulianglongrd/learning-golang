package main

import (
	"errors"
	"fmt"
	"math"
)

// 自定义错误可以使用 errors包的New函数 或 fmt包的Errorf函数
// func New(text string) error {}
// func Errorf(format string, a ...interface{}) error {}

func errorsNew(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

func errorsFmt(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -1.0
	_, err := errorsNew(radius)
	if err != nil {
		fmt.Println(err)
	}

	_, err2 := errorsFmt(radius)
	if err2 != nil {
		fmt.Println(err2)
	}
}
