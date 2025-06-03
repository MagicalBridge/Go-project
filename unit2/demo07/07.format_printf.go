package main

import "fmt"

func main() {
	num := 10
	num2 := 20
	// %d 这个占位符是有特殊含义的，意思是输出整数类型的值 d 是 decimal 十进制的意思
	// 相应的，对于小数也有特定的符号对应，后面会进行介绍
	fmt.Printf("num = %d, num2 = %d ", num, num2) // num=10, num2=20

	// 我们想打印浮点小数，那么可以使用%f
	num3 := 10.1

	// 打印num3，默认保留6位小数，
	fmt.Printf("num3 = %f", num3) // num3=10.100000

	// 如果我们想保留2位小数，可以使用%.2f
	fmt.Printf("num3 = %.2f", num3) // num3=10.10
}
