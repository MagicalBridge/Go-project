package main

import (
	"fmt"
	"sync"
)

// 使用 WaitGroup 来控制协程的退出
// wg.Add(1) 如果放在外层，一定要注意，里面的数字一定要和携程的数量保持一致

var wg = sync.WaitGroup{}

func main() {
	// 启动5个协程
	for i := 0; i <= 5; i++ {
		wg.Add(1) // 协程开始的时候加1操作
		go func(n int) {
			fmt.Println(n)
			wg.Done() // 协程执行完成的时候减一
		}(i)
	}
	// 主线程一直阻塞，什么时候wg减少为0了，就停止
	wg.Wait()
}

/**
这段代码使用了 `sync.WaitGroup` 来确保主 goroutine 在所有其他 goroutine 完成执行之前不会退出。让我们详细解析这段代码：

### 代码解析：

1. **定义 `sync.WaitGroup` 变量**：
   ```go
   var wg = sync.WaitGroup{}
   ```
   `sync.WaitGroup` 是一个用于等待一组 goroutine 完成的同步原语。它有三个主要方法：
   - `Add(delta int)`：增加等待的 goroutine 计数。
   - `Done()`：减少等待的 goroutine 计数。
   - `Wait()`：阻塞直到等待的 goroutine 计数减少到 0。

2. **启动 5 个 goroutine**：
   ```go
   for i := 0; i <= 5; i++ {
       wg.Add(1) // 协程开始的时候加1操作
       go func(n int) {
           fmt.Println(n)
           wg.Done() // 协程执行完成的时候减一
       }(i)
   }
   ```
   - `wg.Add(1)`：在每个 goroutine 启动之前，调用 `Add(1)` 来增加等待的 goroutine 计数。
   - `go func(n int) { ... }(i)`：启动一个新的 goroutine，并将当前的循环变量 `i` 作为参数传递给 goroutine。
   - `fmt.Println(n)`：在 goroutine 中打印传入的参数 `n`。
   - `wg.Done()`：在 goroutine 执行完成后，调用 `Done()` 来减少等待的 goroutine 计数。

3. **等待所有 goroutine 完成**：
   ```go
   wg.Wait()
   ```
   - `wg.Wait()`：阻塞主 goroutine，直到所有启动的 goroutine 都调用了 `Done()`，即等待的 goroutine 计数减少到 0。

### 输出：
由于 goroutine 的并发执行，输出顺序仍然是不确定的。例如，输出可能是 `0, 1, 2, 3, 4, 5` 的任意排列组合。

### 可能的输出：
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
这段代码确保了主 goroutine 在所有其他 goroutine 完成执行之前不会退出，但输出的顺序仍然是随机的，因为 goroutine 的执行顺序是并发的。
*/
