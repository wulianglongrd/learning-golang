package main

import "fmt"

func main() {
	mapDefine()
	fmt.Println("-----------")
	readMap()
	fmt.Println("-----------")
	mapUpdate()
	fmt.Println("-----------")
	mapLength()
}

func mapDefine() {
	// 声明变量：var name map[key_type]value_type
	var m map[string]string
	m = make(map[string]string)

	m["foo"] = "aa"
	m["bar"] = "bb"

	fmt.Println(m) // map[bar:bb foo:aa]

	m2 := map[string]string{"foo": "11", "bar": "22"}
	fmt.Println(m2) // map[bar:22 foo:11]
}

func readMap() {
	m := map[string]string{
		"foo": "11",
		"bar": "22",
	}

	v, ok := m["xxx"]
	fmt.Println(ok, v)

	v2, ok2 := m["foo"]
	fmt.Println(ok2, v2)
}

// map 是引用类型
func mapUpdate() {
	m := map[string]string{
		"foo": "11",
		"bar": "22",
	}
	m2 := m

	m2["foo"] = "xxx"
	fmt.Println(m)  // map[bar:22 foo:xxx]
	fmt.Println(m2) // map[bar:22 foo:xxx]
}

func mapLength() {
	m := map[string]string{
		"foo": "11",
		"bar": "22",
	}
	fmt.Println(len(m)) // 2
}
