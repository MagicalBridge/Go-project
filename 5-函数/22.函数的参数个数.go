package main

import "fmt"

// 求1-100之间所有数字之和

func SumAdd() {
	var sum int

	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func SumAdd2(num int) {
	var sum int

	for i := 1; i <= num; i++ {
		sum += i
	}

	fmt.Println(sum)
}

// 定义两个形参
// 调用的时候传递两个实参，实参和形参的个数保持一致 类型也要保持一致
func add(num1 int, num2 int) {
	fmt.Println(num1 + num2)
}

func main() {
	SumAdd()
	SumAdd2(200)
	SumAdd2(300)
	add(8, 7)
}
