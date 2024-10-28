package main

import "fmt"

func main() {
	// 定义管道 声明管道
	var intChan2 chan int

	//	通过 make 初始化 管道可以 存放3个int类型的数据
	intChan2 = make(chan int, 3)

	//	在管道中存放数据：
	intChan2 <- 10
	intChan2 <- 20

	close(intChan2)

	//	再次写入数据会出错
	intChan2 <- 30

	num := intChan2

	fmt.Println(num)
}

/**
### 修正后的代码：

```go
func main() {
	// 定义管道 声明管道
	var intChan2 chan int

	// 通过 make 初始化 管道可以 存放3个int类型的数据
	intChan2 = make(chan int, 3)

	// 在管道中存放数据：
	intChan2 <- 10
	intChan2 <- 20

	// 关闭管道
	close(intChan2)

	// 从管道中读取数据
	num1 := <-intChan2
	num2 := <-intChan2

	// 输出读取的数据
	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)

	// 尝试再次读取数据（管道已关闭，读取到的值为默认值0）
	num3 := <-intChan2
	fmt.Println("num3:", num3)
}
```

### 解释：

1. **关闭管道**：
   ```go
   close(intChan2)
   ```
   关闭管道后，不能再向管道中写入数据，但可以从管道中读取数据，直到管道为空。

2. **从管道中读取数据**：
   ```go
   num1 := <-intChan2
   num2 := <-intChan2
   ```
   从管道中读取数据，`num1` 和 `num2` 分别读取到 `10` 和 `20`。

3. **尝试再次读取数据**：
   ```go
   num3 := <-intChan2
   ```
   由于管道已经关闭且为空，`num3` 读取到的值为默认值 `0`。

### 输出：
```
num1: 10
num2: 20
num3: 0
```

### 总结：
修正后的代码展示了如何正确地使用管道进行读写操作，并在管道关闭后如何处理读取操作。
*/
