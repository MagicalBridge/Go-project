package main

import "fmt"

// 1、遍历数组 for len
// 2、for range
func main() {
	var Num [5]int = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(Num); i++ {
		fmt.Println(Num[i])
	}

	var Num2 = [5]int{5, 4, 3, 2, 1}
	for i, v := range Num2 {
		fmt.Println(i, v)
	}
}
