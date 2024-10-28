package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		p, ok := hashTable[target-x]
		if ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func main() {
	// 第一种初始化的方式, 从编辑器的角度来看，前面的类型可以省略，可以自动推导
	var m map[int]string = map[int]string{1: "louis", 2: "btc"}
	fmt.Println(m) // map[1:louis 2:btc]
	fmt.Println(m[2])

	// 自动推导类型的形式
	m1 := map[int]string{1: "louis", 2: "eth"} // map[1:louis 2:eth]
	fmt.Println(m1)

	// 第二个值是个bool值
	value, ok := m[2]

	if ok {
		fmt.Println(value) // btc
	}

	// 通过for range 遍历map
	for key, value := range m1 {
		fmt.Println(key)
		fmt.Println(value)
	}

	// 使用make函数创建
	m2 := make(map[string]int, 10)
	m2["1"] = 10
	m2["2"] = 20

	// 删除map中的某个键
	delete(m2, "2")
	fmt.Println(m2) // map[1:10 2:20]

	res := twoSum([]int{1, 2, 3, 4, 5}, 5)
	fmt.Println(res)
}
