package main

import "fmt"

// 返回值可以返回单个值，也可以返回多个值 返回值类型也需要指定

func addResult(num1 int, num2 int) int {
	var sum int
	sum = num1 + num2
	return sum
}

// 函数还可以返回多个值，写法如下：

func GetResult() (int, int) {
	var num1 = 10
	var num2 = 20
	return num1, num2
}

func main() {
	var s = addResult(1, 3)
	fmt.Println(s)

	var s1, s2 = GetResult()
	fmt.Println(s1)
	fmt.Println(s2)
}
