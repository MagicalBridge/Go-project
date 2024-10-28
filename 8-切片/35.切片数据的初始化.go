package main

import "fmt"

func main() {
	// 初始化一个切片，没有赋值
	var s []int
	// 向切片中放入数据,这里使用的是append方法
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s) // [1 2 3 4 5]

	//	第二种方式是自动推导,语法看起来稍微有些怪异
	s1 := []int{1, 2, 3, 4, 5}
	// 通过索引修改对应的值
	s1[3] = 88
	fmt.Println(s1) // [1 2 3 88 5]

	//	可以使用循环的方式来进行初始化
}
