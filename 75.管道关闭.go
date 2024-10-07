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
