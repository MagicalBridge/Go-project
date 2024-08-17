package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 将布尔值转换为字符串 这个方法接收一个布尔值并将其转换为一个字符串。
	str := strconv.FormatBool(true)
	fmt.Println(str) // "true"

	// 将int类型转换成字符串 这个方法接收一个整数并将其转换为一个字符串。
	str2 := strconv.Itoa(10)
	fmt.Println(str2) // "10"

	// 将一个字符串转换为一个整型
	num, err := strconv.Atoi("15")
	if err != nil { // 检查是否出现错误
		fmt.Println("转换错误:", err)
	} else {
		fmt.Println(num) // 15
	}

	// 将一个字符串类型的值转换为bool
	strBool := "true"
	b, err := strconv.ParseBool(strBool)
	if err != nil { // 检查是否出现错误
		fmt.Println("转换错误:", err)
	} else {
		fmt.Println(b) // true
	}
}
