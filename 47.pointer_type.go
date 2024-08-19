package main

import "fmt"

// 指针也是一个变量，但它是一种特殊的变量，因为它存储的数据不仅仅是一个普通的值，如简单的字符串或者整数，而是另一个变量的内存地址

func main() {
	var a = 10

	// 指针的语法 var 变量名 *类型
	var p *int

	// 使用&符号 + 变量名 给指针变量赋值
	p = &a

	// 使用 *p 这个方式打印值，这里打印的是10。
	fmt.Println(*p)

	// 如果指针变量前面加上 &符号，那打印的是p的地址，而不是值，这点需要格外注意
	fmt.Println(&p) // 0xc0000ac018

	*p = 100       // 给p重新赋值 *p这种形式是解引用的语法,需要特别注意
	fmt.Println(a) // a的结果也会发生改变,从10变为100

	// 空指针：直接声明一个指针，但是没有引用任何地址
	// 对于空指针，不要直接给空指针赋值
	//var p1 *int
	// *p1 = 900 // 这种写法是不推荐的，虽然没有报错
	//fmt.Println(p1)  // <nil>
	// fmt.Println(*p1) // invalid memory address or nil pointer dereference

	// 数组指针的定义：var 数组指针的变量 *[下标] 类型
	// 这里初始化一个长度为10的数组,并填充相应的值
	arr_pointer_type := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 创建一个指针变量
	var pArr_pointer_type *[10]int

	// 将数组的指针赋值给指针变量p
	pArr_pointer_type = &arr_pointer_type

	// 打印数组中的所有元素
	fmt.Println(*pArr_pointer_type)

	// 打印数组中的某一个值, 这里需要注意，因为[]的运算符的优先级比 * 要高，所以，(*pArr_pointer_type) 需要加上括号
	fmt.Println((*pArr_pointer_type)[0])
	// 在go语言中，对于这种操作其实是有简单写法的，下面这种写法和上面的写法是等价的
	fmt.Println(pArr_pointer_type[0])

	// 我想利用p循环打印每一项，应该怎么操作呢
	for j := 0; j < len(pArr_pointer_type); j++ {
		fmt.Println(pArr_pointer_type[j])
	}

	// 将数组指针作为参数传递给函数，如果在函数内部修改数组指针的值，那么原来的数组的值也会跟着改变
}
