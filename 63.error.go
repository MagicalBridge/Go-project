package main

import (
	"errors"
	"fmt"
)

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
