package main

import "fmt"

// 在go语言中，提供了两种浮点数据类型 一个是float32，一个是float64
// 建议使用float64 来声明浮点类型，因为一些函数库中都是使用float64的
// float32 精确到小数点后面7位，float64 精确到小数点后面15位
// 这里有一点需要注意，自动推导类型创建的变量中，创建浮点类型，默认为float64

func main() {
	var num float64
	var num2 float64
	num = 123.45
	num2 = 123.4567
	fmt.Printf("%f \n", num) // 123.450000
	fmt.Printf("%.2f", num2) // 123.46
}
