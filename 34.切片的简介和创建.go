package main

import "fmt"

// 为什么会有切片：
//
//		1、数组的容量是固定的，不能自动扩展
//		2、值传递，数组作为函数参数的时候，会将整个数组值拷贝一份给形参。
//	 	在go语言中，几乎所有的场景，都可以使用切片来代替数组
//
// 切片长度是不固定的，比数组更加灵活，源码中的实现，它其实是一个结构体。
func main() {
	// 创建一个切片
	var s []int
	fmt.Println(s)
	fmt.Println(len(s)) // 0

	//	使用自动推导的方式创建切片
	s1 := []int{}
	fmt.Println(s1) // []

	// 使用make函数创建切片
	s2 := make([]int, 3, 5)
	fmt.Println(s2) // [0 0 0]
}
