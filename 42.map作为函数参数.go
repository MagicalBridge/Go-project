package main

import "fmt"

// 在函数中修改map的值，会影响到原来的map的值
// map是无序的
func main() {
	var m map[int]string = map[int]string{1: "btc", 2: "eth"}
	DeleteMap(m)

	fmt.Println(m)
}

func DeleteMap(m map[int]string) {
	delete(m, 2)
}
