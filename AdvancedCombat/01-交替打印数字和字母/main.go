package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建两个无缓冲的channel, 用于goroutine之间的通信
	letter, number := make(chan bool), make(chan bool)
	// 创建一个WaitGroup, 用于等待goroutine 完成
	wait := sync.WaitGroup{}

	// 启动一个goroutine, 用于打印数字
	go func() {
		// 初始化变量i为1, 用于记录当前打印的数字
		i := 1
		// 无限循环,不停地打印数字
		for {
			// 使用select语句, 监听channel的状态
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				// 向letter channel发送信号, 通知打印字母的goroutine可以继续执行
				letter <- true
			}
		}
	}()

	// 将 WaitGroup 的计数器加1, 表示有一个goroutine需要等待
	wait.Add(1)

	// 启动一个goroutine, 用于打印字母
	go func(wait *sync.WaitGroup) {
		// 初始化一个变量i,用于记录当前打印的字母
		i := 'A'
		for {
			select {
			// 从letter channel中接收信号, 表示可以打印字母了
			case <-letter:
				// 如果已经打印完了所有的字母, 则调用WaitGroup的Done方法, 表示当前goroutine已经完成, 然后返回
				if i >= 'Z' {
					wait.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				// 向number channel发送信号, 通知打印数字的goroutine可以继续执行
				number <- true
			}
		}
	}(&wait)

	// 向number channel发送信号, 表示可以开始打印数字了
	number <- true

	// 等待所有的goroutine完成
	wait.Wait()
}
