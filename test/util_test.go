package test

import (
	"fmt"
	"reflect"
	"testing"
)

// go test -v
// go test -v -run="Add" -cover
func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error(`Add(1, 2) != 3`)
	}
}

func TestSplit(t *testing.T) {
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(Split("a,b,c"), want) {
		t.Error("result not excepted")
	}
}

// go test -v -run="Split" -coverprofile=c.out
// go tool cover -html=c.out
func TestSplit2(t *testing.T) {
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(Split("a，b，c"), want) {
		t.Error("result not excepted")
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n int
		v int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
	}

	for _, item := range tests {
		got := FactorialForLoop(item.n)
		if got != item.v {
			t.Errorf("FactorialForLoop func, n: %d, want: %d, got: %d", item.n, item.v, got)
		}
	}

	for _, item := range tests {
		got := FactorialRecursion(item.n)
		if got != item.v {
			t.Errorf("FactorialRecursion func, n: %d, want: %d, got: %d", item.n, item.v, got)
		}
	}
}

func BenchmarkFactorialForLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialForLoop(100)
	}
}

func BenchmarkFactorialRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialRecursion(100)
	}
}

// Example函数以下的注释文案使用的是固定的格式，则否go test无法匹配测试结果（不符合格式的默认是 PASS）
func ExampleAdd() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(2, 3))
	// output:
	// 3
	// 500
}
