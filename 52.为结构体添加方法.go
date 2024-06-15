package main

import "fmt"

type Student52 struct {
	id   int
	name string
	age  int
}

// PrintShow 接收者
func (s Student52) PrintShow() {
	fmt.Println(s)
}

// EditInfo 一般都是使用结构体指针的类型
func (s *Student52) EditInfo() {
	s.age = 20
}

func main() {
	stu := Student52{101, "张三", 18}
	stu.PrintShow() //完成对方法的调用，将stu中的值，传递给了方法的s}

	stu.EditInfo()

	stu.PrintShow()
}
