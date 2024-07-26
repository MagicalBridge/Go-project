package main

import "fmt"

/**
1 变量声明使用 var 关键字。并且要声明类型，对于int类型的变量来说，默认值是0
2 变量声明 可以使用 “ ，”（逗号）进行分割，一次性可以声明多个变量
3 Print可以进行打印输出，Println函数可以进行打印换行输出
*/

func main() {
	// 声明了一个int类型的变量，取名为age，但是没有给这个变量赋值
	var age int
	// 打印这个变量，可以看到控制台输出的是0，说明对于一个int类型的变量来说
	// 如果我们没有初始化赋值，它的默认值就是0；
	fmt.Println(age)

	// 这里我声明了两个变量，中间用逗号进行分割，这种写法可以一次性声明多个相同数据类型的变量
	var num, sum int
	fmt.Println(num, sum)

	// 这里我声明了一个布尔类型的变量，默认值是false
	var flag bool
	fmt.Println(flag)

}
