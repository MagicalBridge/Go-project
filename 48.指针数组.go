package main

import "fmt"

// 指针数组：数组元素是指针类型，指针数组指的是一个数组中存储的都是指针（也就是地址）也就是一个存储了地址的数组
// var 数组名 [下标]*类型
func main() {
	var p [2]*int
	i := 10
	var j = 20
	p[0] = &i
	p[1] = &j

	fmt.Println(p)     // [0x1400009c018 0x1400009c020]
	fmt.Println(*p[0]) // 不要加括号
}
