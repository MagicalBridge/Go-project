package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	addr string
}

func main() {
	// 全部初始化
	var arr = [2]Student{
		{id: 10,
			name: "btc",
			age:  12,
			addr: "china"},
		{id: 11,
			name: "ada",
			age:  12,
			addr: "china"},
	}
	fmt.Println(arr)         // [{10 btc 12 china} {11 ada 12 china}]
	fmt.Println(arr[0])      // {10 btc 12 china}
	fmt.Println(arr[0].name) // btc

	//	循环输出结构体
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		fmt.Println(arr[i].name)
	}

	for k, v := range arr {
		fmt.Println(k)
		fmt.Println(v.name)
	}
}
