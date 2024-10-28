package main

import (
	"fmt"
	"sync"
)

// 为了解决多个携程冲突的问题，我们给协程加上一个锁的机制
var totalNum2 int
var wg3 = sync.WaitGroup{}
var lock = sync.Mutex{}

func add3() {
	defer wg3.Done()

	for i := 0; i < 100000; i++ {
		lock.Lock()
		totalNum2 = totalNum2 + 1
		lock.Unlock()
	}
}

func sub3() {
	defer wg3.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		totalNum2 = totalNum2 - 1
		lock.Unlock()
	}
}

func main() {
	wg3.Add(2)
	go add3()
	go sub3()
	wg3.Wait()

	// 加上锁之后，无论执行多少次，结果总是0
	fmt.Println(totalNum2)
}
