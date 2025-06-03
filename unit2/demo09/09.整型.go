package main

import "fmt"

// 在go语言中，整型分为两种类型: 有符号整型 和 无符号整型
// 有符号整型包含正整数、负整数和0
// 无符号整型包含正整数和0，不包含负整数

// 简单来说，无符号整型是不能包含负数的.
func main() {
	// 在go语言中整型有 int8 int16 int32 int64
	var a int8 = 127 // 8位有符号整型，范围是 -128 到 127
	fmt.Println(a)

	var b int16 = 32767 // 16位有符号整型，范围是 -32768 到 32767
	fmt.Println(b)

	var c int32 = 2147483647 // 32位有符号整型，范围是 -2147483648 到 2147483647
	fmt.Println(c)

	var d int64 = 9223372036854775807 // 64位有符号整型，范围是 -9223372036854775808 到 9223372036854775807
	fmt.Println(d)

	//var age uint
	// 下面这个赋值就错了，因为我用一个有符号的整型赋值给了一个无符号整型
	//age = -10
	//fmt.Println(age) // constant -10 overflows uint
}
