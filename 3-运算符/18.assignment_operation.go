package main

import "fmt"

// 注意：算数运算符的优先级要高于赋值运算符
func main() {
	var num1 int = 10
	num1 %= 2 + 3
	fmt.Println("num1", num1) // 0
}
