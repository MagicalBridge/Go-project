package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2) // 创建一个容量为2的缓冲Channel

	go func() {
		fmt.Println("Sending data...")
		ch <- 1
		ch <- 2
		ch <- 3 // 缓冲区已满，阻塞直到有接收者接收数据
		fmt.Println("Data sent")
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Receiving data...")
	value1 := <-ch
	fmt.Println("Received:", value1)
	value2 := <-ch
	fmt.Println("Received:", value2)
	value3 := <-ch
	fmt.Println("Received:", value3)
}
