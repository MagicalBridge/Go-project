package main

import (
	"fmt"
	"sync"
)

// 下面这个例子演示的是多个协程操作的同一份的数据，输出的结果和我们的预期不太一样
// 最后打印的结果比较随机，思考下为什么会出现这样的问题，
// 如何解决这个问题：确保一个协程在执行逻辑的时候另外的协程不执行
var totalNum int
var wg2 = sync.WaitGroup{}

func add2() {
	defer wg2.Done()
	for i := 0; i < 100000; i++ {
		totalNum = totalNum + 1
	}
}

func sub() {
	defer wg2.Done()
	for i := 0; i < 100000; i++ {
		totalNum = totalNum - 1
	}
}

func main() {
	wg2.Add(2)
	go add2()
	go sub()
	wg2.Wait()

	fmt.Println(totalNum)
}
