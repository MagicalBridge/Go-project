package main

import "fmt"

// 循环条件不加括号，感觉好别扭呀
// go 语言中是没有提供while循环的
func main() {
	var i int
	for i = 1; i <= 10; i++ {
		fmt.Println("为什么循环不加括号啊")
	}

	var j int
	var sum int

	for j = 1; j <= 100; j++ {
		sum = sum + j
	}

	fmt.Println("sum", sum)
}
