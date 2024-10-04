package main

import (
	"fmt"
	"time"
)

// 使用匿名函数启动携程
/**
在go语言中也有闭包的概念，这个和JavaScript是一样的，感觉比较神奇,可以给闭包传递参数
5
0
4
2
1
3
*/
func main() {
	for i := 0; i <= 5; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(time.Second * 2)
}
