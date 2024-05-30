package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4}
	var s1 = []int{7, 8}
	copy(s, s1)

	// 上面的操作中，将s1拷贝到s上，这样就覆盖了s上的前面两个位置
	// 同样的，如果反过来，将s拷贝到s1上面，s1前两个位置也会被覆盖
	fmt.Println(s)
}
