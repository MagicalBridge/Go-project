package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
请编写一个程序，完成如下功能:
	1)、在主线程中，开启一个goroutine，该goroutine每隔1秒输出"hello golang"
	2)、在主线程中也每隔一秒输出"hello msb"，输出10次后，退出程序
	3)、要求主线程和goroutine同时执行
*/
// 开启携程之后，打印是交替进行的
//hello main0
//hello golang 0
//hello golang 1
//hello main1
//hello main2
//hello golang 2
//hello golang 3
//hello main3
//hello main4
//hello golang 4
//hello golang 5
//hello main5
//hello main6
//hello golang 6
//hello main7
//hello golang 7
//hello main8
//hello golang 8
//hello main9
//hello golang 9
//hello main10
//hello golang 10

func testFun() {
	for i := 0; i <= 10; i++ {
		fmt.Println("hello golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go testFun()
	for i := 0; i <= 10; i++ {
		fmt.Println("hello main" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
