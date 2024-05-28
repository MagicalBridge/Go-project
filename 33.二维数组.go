package main

import "fmt"

// 如何定义二维数组
// var 数组名 [x][y] 数据类型
func main() {
	var arr [2][3]int
	arr[0][2] = 123
	arr[1][1] = 456

	fmt.Println(arr) // [[0 0 123] [0 456 0]]
}
