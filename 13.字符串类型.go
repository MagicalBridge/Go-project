package main

import "fmt"

// 用双引号引用的
// 在go语言中一个汉字占3个字符 打印中文的时候需要注意
func main() {
	var name string = "louis"
	fmt.Println(name)
	fmt.Println(len(name))
}
