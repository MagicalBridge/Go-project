package main

import (
	"fmt"
	"runtime"
	"time"
)

func GoroutineNum() {
	//	1、统计当前存在的goroutine数量
	go func() {
		for {
			fmt.Println("NumGoroutine:", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	//	2、启动大量的goroutine
	for {
		go func() {
			time.Sleep(100 * time.Second)
		}()
	}
}

// 使用匿名函数启动携程
/**
在go语言中也有闭包的概念，这个和JavaScript是一样的，感觉比较神奇,可以给闭包传递参数
5
0
4
2
1
3
*/
func main() {
	//for i := 0; i <= 5; i++ {
	//	go func(n int) {
	//		fmt.Println(n)
	//	}(i)
	//}
	//time.Sleep(time.Second * 2)

	GoroutineNum()
}

/**
这个程序的输出是不确定的，因为

`go func(n int) { fmt.Println(n) }(i)`

是在一个 goroutine 中执行的，而 goroutine 的执行顺序是并发的，不受控制。

### 解释：
1. **并发执行**：`go func(n int) { fmt.Println(n) }(i)` 启动了一个新的 goroutine，这些 goroutine 是并发执行的。
2. **输出顺序不确定**：由于 goroutine 的执行顺序是不确定的，`fmt.Println(n)` 的输出顺序也是不确定的。
3. **延迟**：`time.Sleep(time.Second * 2)` 是为了确保主 goroutine 等待足够长的时间，以便让所有的 goroutine 都有机会执行完毕。

### 可能的输出：
由于并发执行的特性，输出可能是 `0, 1, 2, 3, 4, 5` 的任意排列组合。例如：

```
3
5
2
1
0
4
```

或者

```
4
1
2
0
5
3
```
### 总结：
这个程序的输出是随机的，取决于 goroutine 的调度顺序。
*/
