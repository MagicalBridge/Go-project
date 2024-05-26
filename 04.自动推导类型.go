package main

import "fmt"

func main() {
	// 自动推导类型并不需要写关键字var和具体的变量类型, 会根据后面的赋值来确定类型
	num := 10
	fmt.Println(num)
}
