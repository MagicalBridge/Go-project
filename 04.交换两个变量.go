package main

import "fmt"

// := 这种形式的赋值 在声明变量的时候，不需要使用var和变量
// 类型，其类型是由所赋值的值来决定的
func main() {
	num := 10
	num2 := 20
	num, num2 = num2, num
	fmt.Println(num, num2)
}
