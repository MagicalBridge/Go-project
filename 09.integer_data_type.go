package main

import "fmt"

// 在go语言中，整型分为两种类型: 有符号整型 和 无符号整型
// 有符号整型包含正整数、负整数和0
// 无符号整型包含正整数和0，不包含负整数

// 简单来说，无符号整型是不能包含负数的。这点和java有区别
func main() {
	var age uint
	// 下面这个赋值就错了，因为我用一个有符号的整型赋值给了一个无符号整型
	age = -10
	fmt.Println(age) // constant -10 overflows uint
}
