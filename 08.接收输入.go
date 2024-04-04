package main

import "fmt"

/*
*
scanf 函数将键盘输入的数据赋值给变量 但是变量前一定要加 & 符号
scan 这个函数使用起来和 scanf 很类似，而且要简单一些，在以后的编程过程中使用这个居多
*/
func main() {
	var age int
	fmt.Println("请输入年龄：")
	// 第一个参数代表的是我想输入的是整数类型，赋值给哪个变量的内存地址
	// 前面需要加上&，这点很重要
	fmt.Scanf("%d", &age)
	fmt.Println("age=", age)
	fmt.Println(&age) // 0xc0000b4008 打印的是内存地址

}
