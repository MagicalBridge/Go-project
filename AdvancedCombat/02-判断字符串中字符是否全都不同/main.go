package main

import (
	"fmt"
	"strings"
)

func isUniqueString(s string) bool {
	// 先判断长度，如果大于3000，直接返回false
	if strings.Count(s, "") > 3000 {
		return false
	}
	// 判断字符是否在ASCII范围内
	for _, v := range s {
		// 如果不在，直接返回false
		if v > 127 {
			return false
		}
		// 如果在，判断是否重复
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

func isUniqueString2(s string) bool {
	// 先判断长度，如果大于3000，直接返回false
	if strings.Count(s, "") > 3000 {
		return false
	}

	for k, v := range s {
		// 如果不在，直接返回false
		if v > 127 {
			return false
		}
		// 如果在，判断是否重复
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}

func main() {
	//	调用 isUniqueString 函数
	fmt.Println(isUniqueString("abc"))
	fmt.Println(isUniqueString("abca"))
	fmt.Println(isUniqueString("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(isUniqueString("abcdefghijklmnopqrstuvwxyza"))

	//	 调用 isUniqueString2 函数
	fmt.Println(isUniqueString2("abc"))
	fmt.Println(isUniqueString2("abca"))
	fmt.Println(isUniqueString2("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(isUniqueString2("abcdefghijklmnopqrstuvwxyza"))
}
