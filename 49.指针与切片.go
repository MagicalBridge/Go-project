package main

import "fmt"

// 切片指针指的是：定义一个指针指向一个切片
// 切片比数组更加灵活，可以理解为动态数组，但是它并不是数组。
func main() {
	// 定义了一个切片 里面有四个值
	s := []int{1, 2, 23, 4}
	// 声明了一个指针变量
	var p *[]int
	// 将切片的内存地址赋值给p
	p = &s
	fmt.Println(*p) // [1 2 23 4]
	// 获取指定切片位置的数据 这里也需要注意：[] 的运算优先级要大于 *
	fmt.Println((*p)[0]) // [1]

	// 我们修改某个位置的切片的值,
	(*p)[0] = 200

	// 发现切片的值已经变化了
	fmt.Println(s) // [200 2 23 4]
	for index, value := range *p {
		fmt.Println("index", index)
		fmt.Println("value", value)
	}
}
