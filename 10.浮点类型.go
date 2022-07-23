package main

import "fmt"

// 建议使用float64 来声明浮点类型
// 默认保留的是6位小数，我们可以控制显示小数的位数
// 这里有一点需要注意，自动推导类型创建的变量中，创建浮点类型，默认为float64
func main() {
	var num float64
	var num2 float64
	num = 123.45
	num2 = 123.4567
	fmt.Printf("%f \n", num) // 123.450000

	fmt.Printf("%.2f", num2) // 123.46
}
