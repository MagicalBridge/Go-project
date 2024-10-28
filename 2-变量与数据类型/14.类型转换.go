package main

import "fmt"

// int float64 进行强制类型转换
// 数据转换的时候 将低类型的转换为高类型的 如果高类型转换为低类型会丢失精度
func main() {
	var num float64 = 3.15
	var num1 int = 20

	fmt.Printf("%d\n", int(num))
	fmt.Printf("%f\n", float64(num1))
}
