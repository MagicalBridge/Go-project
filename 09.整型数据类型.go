package main

import "fmt"


// go 语言中 整型分为两种类型: 有符号整型和无符号整型
// 简单来说，无符号整型是不能包含负数的。这点和java有区别
func main() {
	var age uint
	age = -10
	fmt.Println(age) // constant -10 overflows uint
}