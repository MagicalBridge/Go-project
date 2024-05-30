package main

import "fmt"

func main() {
	// 结构体是由一系列具有相同类型或不同类型的数据组成的数据集合
	//结构体可以很好的管理一批有联系的数据，使用结构体可以提高程序的易读性
	// 成员名称前不能添加var
	type Student struct {
		id   int
		name string
		age  int
		addr string
	}

	// 全部初始化
	var s Student = Student{
		id:   10,
		name: "btc",
		age:  12,
		addr: "china",
	}

	fmt.Println(s) // {10 btc 12 china}

	var s1 = Student{
		name: "btc",
		addr: "china",
	}

	s1.id = 23 // 可以使用具体的成员变量名称赋值

	fmt.Println(s1) // {0 btc 0 china}
}
