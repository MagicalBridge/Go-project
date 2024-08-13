package main

import "fmt"

// 注意：关系运算符只能返回true或者false
func main() {
	var num1 int = 10
	var num2 int = 20
	fmt.Println("num1 == num2", num1 == num2)
	fmt.Println("num1 != num2", num1 != num2)
	fmt.Println("num1 >  num2", num1 > num2)
	fmt.Println("num1 <  num2", num1 < num2)
	fmt.Println("num1 >= num2", num1 >= num2)
	fmt.Println("num1 <= num2", num1 <= num2)
}
