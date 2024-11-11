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
	wg.Add(2)

	// 定义打印数字的 goroutine
	go func() {
		defer wg.Done() // 在 goroutine 完成后通知 WaitGroup
		for i := 1; i <= 28; i++ {
			cond.L.Lock()
			for !printNum {
				cond.Wait()
			}
			fmt.Printf("%d", i)
			if i < 28 {
				i++
				fmt.Printf("%d", i)
			}
			printNum = false
			cond.Signal()
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
			fmt.Printf("%c%c", i, i+1)
			i++
			printNum = true
			cond.Signal()
			cond.L.Unlock()
		}
	}()

	// 等待两个 goroutine 执行完毕
	wg.Wait()
}
