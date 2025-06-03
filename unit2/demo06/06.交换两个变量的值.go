package main

import "fmt"

// := 这种形式的赋值 在声明变量的时候，不需要使用var和变量
// 类型，其类型是由所赋值的值来决定的
func main() {
	num1 := 10
	num2 := 20

	// 这样写的核心含义是将num2赋值给num1，将num1的值赋值给num2
	num1, num2 = num2, num1
	// 打印相关的变量
	fmt.Println(num1, num2) // 20 10
}
