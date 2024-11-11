package main

import (
	"fmt"
	"sync"
)

func main() {
	// 初始化 sync.Cond 和互斥锁
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	// 控制打印顺序的标志
	printNum := true

	// 使用 WaitGroup 等待 goroutine 完成
	var wg sync.WaitGroup
	// 清楚知道，我需要开起两个go协程
	wg.Add(2)

	// 定义打印数字的 goroutine
	go func() {
		defer wg.Done() // 在 goroutine 完成后通知 WaitGroup
		for i := 1; i <= 28; i++ {
			// 保护共享资源 printNum
			cond.L.Lock()
			for !printNum {
				cond.Wait()
			}
			fmt.Printf("%d", i) // 打印当前数字
			if i < 28 {
				i++
				fmt.Printf("%d", i) // 因为每次都打印两个数字，这是打印的下一个数字
			}
			printNum = false // 设置标识，接下来应该打印字母
			cond.Signal()    // 通知等待的字母打印 goroutine 可以继续执行
			cond.L.Unlock()
		}
	}()

	// 定义打印字母的 goroutine
	go func() {
		defer wg.Done() // 在 goroutine 完成后通知 WaitGroup
		for i := 'A'; i <= 'Z'; i++ {
			cond.L.Lock()
			for printNum {
				cond.Wait()
			}
			fmt.Printf("%c%c", i, i+1) // 打印当前字母和下一个字母（因为每次循环打印两个字母）。
			i++                        // 跳过下一个字母，因为，已经打印过了
			printNum = true
			cond.Signal()
			cond.L.Unlock()
		}
	}()

	// 等待两个 goroutine 执行完毕
	wg.Wait()
}
