package main

import (
	"fmt"
)

func writeOnly(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Writing: %d\n", i)
		ch <- i
	}
	close(ch) // 写完后关闭通道
}

func readOnly(ch <-chan int) {
	for value := range ch {
		fmt.Printf("Reading: %d\n", value)
	}
}

func main() {
	ch := make(chan int) // 创建一个普通的管道

	// 通过参数将通道限制为只写和只读
	go writeOnly(ch) // 写数据
	go readOnly(ch)  // 读数据

	// 等待一些时间，以便协程完成任务
	fmt.Println("Press enter to exit...")
	fmt.Scanln()
}
