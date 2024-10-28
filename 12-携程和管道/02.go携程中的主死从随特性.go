package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
接着上面的那个案例，我们来演示下，主线程提前结束，看看子协程如何表现：主死从随
协程会随着主线程的死亡会锤死挣扎一下，所以还会打印 10 和 11。
*/

func testFun2() {
	for i := 0; i <= 10000; i++ {
		fmt.Println("hello golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go testFun2()
	// 只遍历10次，结束的比较早
	for i := 0; i <= 10; i++ {
		fmt.Println("hello main" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

/**
hello main0
hello golang 0
hello main1
hello golang 1
hello golang 2
hello main2
hello main3
hello golang 3
hello main4
hello golang 4
hello main5
hello golang 5
hello golang 6
hello main6
hello main7
hello golang 7
hello golang 8
hello main8
hello main9
hello golang 9
hello main10
hello golang 10
hello golang 11
*/
