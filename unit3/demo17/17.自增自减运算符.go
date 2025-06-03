package main

import "fmt"

// 注意：go语言中是没有前自增 和 前自减的。
func main() {
	var num1 int = 20
	var num2 int = 10
	num1++
	num2--
	fmt.Println("num1", num1)
	fmt.Println("num2", num2)
}
