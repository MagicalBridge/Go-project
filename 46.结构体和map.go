package main

import "fmt"

type Student3 struct {
	id   int
	name string
	age  int
	addr string
}

func main() {
	m := make(map[int]Student3)
	m[1] = Student3{
		id:   10,
		name: "btc",
		age:  12,
		addr: "china",
	}

	m[2] = Student3{
		id:   11,
		name: "eth",
		age:  13,
		addr: "china",
	}

	fmt.Println(m[1])

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
	//结构体作为函数的参数传递的时候，在函数内部修改结构体，是不会影响结构体原来的值的
}
