package main

import "fmt"

// 指针数组：数组元素是指针类型，指针数组指的是一个数组中存储的都是指针（也就是地址）也就是一个存储了地址的数组
// 上面这个定义一定要和 数组指针 区分来看
// var 数组名 [下标]*类型
func main() {
	// 创建了一个长度为2的指针数组
	var p [2]*int
	i := 10
	var j = 20
	// 将两个变量的指针分别存放在指针数组的第一个和第二个位置
	p[0] = &i
	p[1] = &j

	fmt.Println(p) // [0x1400009c018 0x1400009c020]
	// p 存放的本来就是地址了 前面加上* 才是真正的根据对应内存地址取值
	fmt.Println(*(p[0])) // 不要加括号

	//	我们使用range方法来遍历指针数组
	for index, value := range p {
		fmt.Println(index)  // 打印的索引
		fmt.Println(value)  // 打印是地址
		fmt.Println(*value) // 打印的是值
	}
}
