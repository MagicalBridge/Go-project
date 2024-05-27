package main

import "fmt"

// 递归函数适合处理树形结构，不断的调用自身。
// 使用递归函数计算一个数的阶乘
var s = 1

func TestDemo(n int) {
	if n == 1 {
		return
	}
	s *= n
	TestDemo(n - 1)
}

func main() {
	TestDemo(5)
	fmt.Println(s)
}
