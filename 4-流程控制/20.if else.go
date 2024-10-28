package main

import "fmt"

func main() {
	var age int = 20

	age += 12
	// 在go中，if 的条件是可以不写括号的
	if age >= 18 {
		fmt.Println("已经成年了")
	} else {
		fmt.Println("还没有成年")
	}
}
