package main

import "fmt"

// 数组：是指同一系列同一类型数据的集合

// 如何定义数组: var 数组名 [元素数量] 类型

func main() {
	// 1、声明了一个int类型的数组，有5个元素，并且做了初始化
	var Nums [5]int = [5]int{1, 2, 3, 4, 5}
	//可以通过索引访问数组的元素
	fmt.Println(Nums[0])

	// 2、可以部分赋值操作
	Nums2 := [5]int{1, 2} // [1 2 0 0 0]
	fmt.Println(Nums2)

	// 3、还可以指定某个元素初始化
	Nums3 := [5]int{1: 3, 2: 5} // [0 3 5 0 0]
	fmt.Println(Nums3)

	// 4、通过初始化确定数组的长度, 使用的是...形式，先不确定长度
	Nums4 := [...]int{7, 8, 5} // [7,8,5]
	fmt.Println(len(Nums4))

	var Num5 [5]int
	//	5、使用循环的形式赋值
	for i := 0; i < 5; i++ {
		Num5[i] = i + 1
	}
	fmt.Println(Num5)
}
