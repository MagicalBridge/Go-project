package main

import "fmt"


// scanf函数将键盘输入的数据赋值给变量 但是变量前一定要加 & 符号
func main() {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanf("%d", &age)
	fmt.Println("age=",age)
	fmt.Println(&age)  // 0xc0000b4008 打印的是内存地址
}