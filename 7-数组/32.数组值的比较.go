package main

import "fmt"

// 完成两个数组中元素的比较，判断其相同下标对应的元素是否完全一致
func main() {
	var num1 = [5]int{1, 2, 3, 4, 5}
	var num2 = [5]int{1, 2, 3, 4, 5}

	// 可以直接使用 == 来进行比较，但是有一定的限制，就是数组的长度和类型必须一致。
	fmt.Println(num1 == num2)

	var strArr = [2]string{"louis", "chu"}
	fmt.Println(strArr)
}
