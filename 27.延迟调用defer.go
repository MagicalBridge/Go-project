package main

import "fmt"

// 使用defer可以延迟执行相关逻辑
// defer的执行顺序，如果一个函数中有多个defer语句，它们会以LIFO(后进先出)的顺序执行
func main() {
	defer fmt.Println("首先执行")
	fmt.Println("其次执行")
}
