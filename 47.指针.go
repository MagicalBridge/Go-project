package main

import "fmt"

// 指针也是一个变量，但它是一种特殊的变量，因为它存储的数据不仅仅是一个普通的值，如简单的
// 字符串或者整数，而是另一个变量的内存地址
func main() {
	var a = 10
	var p *int
	p = &a
	// 使用*p 这个方式打印值
	fmt.Println(*p)
	*p = 100       // 给p重新赋值
	fmt.Println(a) // a的结果也会发生改变

	//	空指针：直接声明一个指针，但是没有引用任何地址
	// 对于空指针，不要直接给空指针赋值
	var p1 *int
	*p1 = 900        // 这种写法是不推荐的，虽然没有报错
	fmt.Println(p1)  // <nil>
	fmt.Println(*p1) // invalid memory address or nil pointer dereference

}
