package main

import "fmt"

func main() {
	// 第一种初始化的方式
	var m map[int]string = map[int]string{1: "louis", 2: "btc"}
	fmt.Println(m) // map[1:louis 2:bit]

	// 自动推导类型的形式
	m1 := map[int]string{1: "louis", 2: "eth"} // map[1:louis 2:eth]

	fmt.Println(m1)

	// 使用make函数创建
	m2 := make(map[string]int, 10)
	m2["1"] = 10
	m2["2"] = 20
	fmt.Println(m2) // map[1:10 2:20]
}
