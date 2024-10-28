package main

import (
	"fmt"
	"sync"
	"time"
)

// 【1】案例需求:
// 请完成协程和管道协同工作的案例，具体要求:
// 	1) 开启一个writeData协程，向管道中写入50个整数。
// 	2) 开启一个readData协程，从管道中读取writeData写入的数据。
// 	3) 注意:writeData和readDate操作的是同一个管道。
// 	4) 主线程需要等待writeData和readDate协程都完成工作才能退出。

func writeData(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		ch <- i
		fmt.Printf("writeData: %d\n", i)
		time.Sleep(time.Second)
	}
	close(ch) // 写完数据后关闭通道
}

func readData(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("readData: %d\n", data)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2) // 需要等待两个协程

	go writeData(ch, &wg) // 开启写数据的协程
	go readData(ch, &wg)  // 开启读数据的协程

	wg.Wait() // 等待所有协程完成
	fmt.Println("All goroutines completed.")
}
