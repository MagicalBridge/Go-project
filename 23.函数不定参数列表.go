package main

import "fmt"

// 不定参数的形式更加灵活，相当于一个集合，会给参数添加一个编号，是从0开始算的

func TestSum(args ...int) {
	fmt.Println(args[0])
	fmt.Println(args[1])
	fmt.Println(args[2])
}

func main() {
	TestSum(1, 2, 4)
}
