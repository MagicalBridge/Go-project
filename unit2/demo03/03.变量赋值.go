package main

import "fmt"

/*
*
1、对于变量来说，可以初始化进行赋值，也可以先初始化后面再赋值
2、我们同样的可以将一个变量的值赋值给另一个变量
3、一个小例子：交换两个变量，我们可以使用中间变量进行中转
*/
func main() {
	// 初始化一个变量 age，它的类型是int，先不赋值
	var age int

	// 给age变量赋值10
	age = 10
	fmt.Println(age)

	var a int = 10 // 这里的int可以不写，类型推导就可以确定
	var b int
	// 将a的值赋值给b
	b = a
	fmt.Println(b) // 10
}
