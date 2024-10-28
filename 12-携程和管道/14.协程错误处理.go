package main

import (
	"fmt"
	"time"
)

func printnum() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func devide() {

	defer func() {
		//	调用recover函数来捕获错误, defer后面加上匿名函数的调用
		err := recover()
		if err != nil {
			fmt.Println("err", err)
		}
	}()

	num1 := 10
	num2 := 0

	result := num1 / num2
	fmt.Println(result)
}

func main() {
	// 启动两个协程
	go printnum()
	go devide()

	time.Sleep(time.Second)
}
