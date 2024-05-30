package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", arr)
	sorted := bubbleSort(arr)
	fmt.Println("Sorted array:", sorted)
}

/**
解释:

首先定义一个名为bubbleSort的函数,它接受一个整型切片arr作为参数,并返回排序后的切片。
在bubbleSort函数中,使用两个嵌套的for循环来实现冒泡排序。

外层循环控制需要进行多少次遍历, 每次遍历都会将当前剩余未排序元素中最大的元素移动到最右端。
内层循环用于比较相邻元素,如果左边的元素大于右边的元素,则交换它们的位置。


在main函数中,首先定义了一个未排序的整型切片arr。
调用bubbleSort函数,传入arr作为参数,并将返回的排序后的切片赋值给变量sorted。
最后,打印原始的未排序切片和排序后的切片。
*/
