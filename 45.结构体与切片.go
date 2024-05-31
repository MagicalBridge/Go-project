package main

import "fmt"

type Student2 struct {
	id   int
	name string
	age  int
	addr string
}

func main() {
	// 全部初始化
	var s = []Student2{
		{id: 10,
			name: "btc",
			age:  12,
			addr: "china"},
		{
			id:   11,
			name: "ada",
			age:  12,
			addr: "china",
		},
	}
	fmt.Println(s)

	//	循环输出结构体
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for k, v := range s {
		fmt.Println(k)
		fmt.Println(v.name)
	}

	//	可以使用append 函数向切片中添加数据
	s2 := append(s, Student2{
		id:   12,
		name: "UNI",
		age:  12,
		addr: "china",
	})

	fmt.Println(s2)
}
