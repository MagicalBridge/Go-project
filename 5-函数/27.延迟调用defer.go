package main

import "fmt"

// 在函数中，程序员经常需要创建资源，为了在函数执行完毕之后，及时的释放资源，go的设计者提供了defer 关键字

// 使用defer可以延迟执行相关逻辑
// defer的执行顺序，如果一个函数中有多个defer语句，它们会以LIFO(后进先出) 的顺序执行 这是栈的特点

func addTest(num1 int, num2 int) int {
	// 在go语言中，程序遇到defer关键字，不会立即执行defer后面的语句，而是将defer
	// 后面的语句压入一个栈中，然后继续执行函数后面的语句
	// 在函数执行完毕后，从栈中取出语句开始执行，按照先进后出的规则执行语句
	defer fmt.Println("num1=", num1)
	defer fmt.Println("num2=", num2)

	var sum int = num1 + num2
	fmt.Println("sum=", sum)
	return sum
}

func main() {
	fmt.Println(addTest(30, 60))
}
