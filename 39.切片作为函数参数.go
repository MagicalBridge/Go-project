package main

import "fmt"

func Init(nums []int) {
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		nums[i] = nums[i] + 1
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	Init(s)
	// 在函数中修改切片的值，会影响到原来的切片 这个和数组有比较大的差异
	fmt.Println(s) // [2 3 4 5 6]
}
