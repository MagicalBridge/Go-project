package main

//【1】管道的遍历:
// 管道支持for-range的方式进行遍历，请注意两个细节
// 1）如果管道没有关闭，则会出现deadlock的错误
// 2) 在遍历时，如果管道已经关闭，则会正常遍历数据，遍历完后，就会退出遍历

func main() {
	// 	定义管道、声明管道
	var intChan chan int
	//	通过make初始化管道
	intChan = make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan <- i
	}

	//	在遍历前，如果没有关闭管道，就会出现deadblock错误
	close(intChan)
	for v := range intChan {
		println(v)
	}
}
