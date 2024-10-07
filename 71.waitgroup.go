package main

import (
	"fmt"
	"sync"
)

// 使用waitgroup来控制协程的退出
// wg.Add(1) 如果放在外层，一定要注意，里面的数字一定要和携程的数量保持一致

var wg = sync.WaitGroup{}

func main() {
	// 启动5个协程
	for i := 0; i <= 5; i++ {
		wg.Add(1) // 协程开始的时候加1操作
		go func(n int) {
			fmt.Println(n)
			wg.Done() // 携程执行完成的时候减一
		}(i)
	}
	// 主线程一直阻塞，什么时候wg减少为0了，就停止
	wg.Wait()
}
