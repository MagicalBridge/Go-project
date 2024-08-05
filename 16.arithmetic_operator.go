package main

import "fmt"

// 对于算数运算符使用时候需要注意，当使用 "/"和"%" 操作时，
// 除数不能为0，否则会报错，这点和在js中的表现非常不一致
func main() {
	var num1 int = 20
	var num2 int = 10
	fmt.Println("num1 + num2", num1+num2)
	fmt.Println("num1 - num2", num1-num2)
	fmt.Println("num1 * num2", num1*num2)
	fmt.Println("num1 / num2", num1/num2)
}
