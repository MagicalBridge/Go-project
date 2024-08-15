package main

import (
	"errors"
	"fmt"
)

// 接收两个整型参数并返回它们的除法结果，如果除数为0，则返回异常字符串。
func dive(a int, b int) (value int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	value = a / b
	return
}

func main() {
	a := 10
	b := 2
	value, err := dive(a, b)
	if err != nil {
		fmt.Println("runtime error：", err)
	} else {
		fmt.Println(value)
	}
}
