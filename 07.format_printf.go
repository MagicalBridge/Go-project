package main

import "fmt"

func main() {
	num := 10
	num2 := 20
	// %d 这个占位符是有特殊含义的，意思是输出整数类型的值
	// 相应的，对于小数也有特定的符号对应，后面会进行介绍
	fmt.Printf("num = %d, num2 = %d ", num, num2) // num=10, num2=20
}
