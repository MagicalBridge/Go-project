package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	close(ch) // 关闭 channel

	// 尝试发送消息
	go func() {
		ch <- 42 // 这将导致 panic
	}()

	// 等待片刻以让 goroutine 执行
	fmt.Println("Waiting...")
}

/**
在上面的代码中，close(ch) 关闭了 channel，随后尝试向该 channel 发送消息会导致运行时 panic。

为了避免这种 panic，可以在发送消息之前检查 channel 是否已关闭，
但 Go 的 channel 不支持直接检查是否关闭，因此通常的做法是设计好程序逻辑，
确保在关闭 channel 之后不会再有发送操作。例如，可以使用一个标志来指示 channel 的状态，
或在关闭 channel 之前确保所有发送操作已经完成。

if ch != nil {
    select {
    case ch <- 42:
        // 发送成功
    default:
        // channel 已满或关闭
    }
} else {
    fmt.Println("Channel is nil!")
}
*/
