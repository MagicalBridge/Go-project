package main

import "fmt"

// 定义：用 "单引号" 引起来的 "单个字符" 为字符类型, 限定处的文字非常重要
// 打印出来的是美国标准协会设定的码表
func main() {
	var ch byte = 'a'
	fmt.Println(ch) // 97
}
