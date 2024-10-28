package main

import "fmt"

func main() {
	var Num = [5]int{1, 2, 3, 4, 5}
	getPrint(Num)

	setItemValue(Num)

	// 虽然我们setItemValue中修改了数组的元素，但是数组并没有发生改变
	fmt.Println(Num) // [1 2 3 4 5]
}

// 将数组作为参数传递
func getPrint(n [5]int) {
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}
}

// 如果在函数中修改数组元素，并不会影响原来的数组
func setItemValue(n [5]int) {
	for i := 0; i < len(n); i++ {
		n[i] = i + 1
	}
}
