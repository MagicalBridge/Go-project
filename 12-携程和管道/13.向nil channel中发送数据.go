package main

import (
	"fmt"
)

/**
向一个 nil channel 发送消息会导致程序阻塞。具体来说，Go 的 channel 在以下情况下会导致阻塞：

	1、发送到一个 nil channel：当你尝试向一个 nil channel 发送消息时，程序将停在那里，永远不会继续执行，因为没有任何 goroutine 可以接收这个消息。
	nil channel 被认为是“不可用”的，不会有任何操作发生。

	2、接收从一个 nil channel：类似地，从一个 nil channel 接收消息也会导致阻塞，因为没有任何消息可以接收。
*/

func main() {
	var ch chan int // nil channel

	// 尝试发送消息
	go func() {
		ch <- 42 // 这将导致阻塞
	}()

	// 等待片刻以让 goroutine 执行
	fmt.Println("Waiting...")
}

/**
在上面的代码中，ch 是一个 nil channel，尝试向它发送消息将导致程序永远阻塞在那一行。
为了避免这种情况，确保在发送或接收之前检查 channel 是否为 nil，或者确保 channel 已经被初始化：

if ch != nil {
    ch <- 42
} else {
    fmt.Println("Channel is nil!")
}


*/
