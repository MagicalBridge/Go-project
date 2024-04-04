package main

import "fmt"

/*
1、对于变量来说，可以初始化进行赋值，也可以先初始化后面再赋值
2、我们同样的可以将一个变量的值赋值给另一个变量
*/
func main() {
	// 初始化一个变量 age，它的类型是int
	var age int
	// 给age变量赋值10
	age = 10
	fmt.Println(age)

	var a int = 10
	var b int
	// 将a的值赋值给b
	b = a
	fmt.Println(b)
}
