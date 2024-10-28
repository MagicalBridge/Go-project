package main

import "fmt"

// 在go中 管道本质上是一种数据结构：队列，所以符合队列的数据结构的特点，先进先出
// 自身是线程安全的，多协程访问的时候，不需要加锁。不会发生资源争抢的问题
// 管道是有类型的，一个string的管道只能存放string类型的数据

// 【1】管道的定义
// var 变量名chan
// 数据类型
// PS1: chan管道关键字
// PS2: 数据类型指的是管道的类型，里面放入数据的类型，管道是有类型的，intChan只能写入整数intPS3:管道是引用类型，必须初始化才能写入数据，即make后才能使用
func main() {

	//	定义管道 声明管道 定义一个int类型的管道
	var intChan chan int
	//	通过make进行初始化：管道可以存放3个int类型的数据
	intChan = make(chan int, 3)

	// 证明管道是引用类型
	fmt.Println("intChan的值 %v", intChan) // 0xc00009e000

	// 向管道中存放数据
	intChan <- 10
	intChan <- 20
	intChan <- 30

	// 注意不能存放大于用量的数据
	// intChan <- 80

	fmt.Println("管道的实际长度：%v, 管道的容量是：%v", len(intChan), cap(intChan))

	// 在管道中读取数据
	num := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(num3)

	//	注意：在没有使用协程的情况下，如果管道的数据已经全部取出，那么再取就会报错
	//num4 := <-intChan
	//fmt.Println(num4)

}
