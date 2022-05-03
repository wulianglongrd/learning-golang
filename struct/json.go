package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	orderId  int    // 首字母未大写，json序列化和反序列化会忽略此字段
	Title    string `json:"title"`
	TotalFee int    `json:"total_fee"`
}

func main() {
	o := Order{orderId: 1, Title: "book", TotalFee: 10}
	fmt.Println(o) // {1 book 10}

	fmt.Println(toStr(o)) // {"title":"book","total_fee":10}

	s := "{\"order_id\":2,\"title\":\"milk\",\"total_fee\":5}"
	fmt.Printf("%+v", toJson(s)) // {orderId:0 Title:milk TotalFee:5}
}

func toJson(s string) Order {
	o := Order{}
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		println(err)
	}
	return o
}

func toStr(o Order) string {
	res, err := json.Marshal(&o)
	if err != nil {
		println(err)
	}
	return string(res)
}
