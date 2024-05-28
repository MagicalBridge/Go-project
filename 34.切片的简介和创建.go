package main

import "fmt"

// 切片长度是不固定的，比数组更加灵活，可以理解为动态数组，但是它并不是数组，切片的使用时非常广泛的
func main() {
	//	创建一个切片
	var s []int
	fmt.Println(s)
	fmt.Println(len(s)) // 0

	//	使用自动推导的方式创建切片
	s1 := []int{}
	fmt.Println(s1) // []

	// 使用make函数创建切片
	// 使用make()函数创建总结：
	// 1.容量是已经开辟的空闸，包括已经初始
	// 2.长度是已经初始化的空间。化的空间和空闲的空间时，一定要注意切片长度要小于容量
	s2 := make([]int, 3, 5)
	fmt.Println(s2) // [0 0 0]
}
