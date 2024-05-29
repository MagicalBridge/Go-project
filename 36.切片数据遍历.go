package main

import "fmt"

func main() {
	var s []int
	s = append(s, 1, 2, 3, 4, 5)
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	// 使用自动推导 + range 的遍历方式处理切片
	s1 := []int{2, 3, 4, 5, 6}
	for i, v := range s1 {
		fmt.Println(i, v)
	}
}
