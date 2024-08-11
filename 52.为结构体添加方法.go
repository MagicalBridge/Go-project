package main

import "fmt"

// 相关的语法
// func (对象 结构体类型) 方法名 (参数列表) (返回值列表) { 代码体 }

type Student52 struct {
	id   int
	name string
	age  int
}

type Teacher52 struct {
	id   int
	name string
}

// PrintShow 接收者
func (s Student52) PrintShow() {
	fmt.Println(s)
}

func (t Teacher52) PrintShow() {
	fmt.Println(t)
}

// EditInfo 一般都是使用结构体指针的类型，在函数内部修改结构体的数据
func (s *Student52) EditInfo() {
	// 对于结构体的修改，会作用于stu
	s.age = 20
}

func main() {
	stu := Student52{101, "张三", 18}
	stu.PrintShow() // 完成对方法的调用，将stu中的值，传递给了方法的s parameter，并调用了方法中的PrintShow()方法

	stu.EditInfo()

	stu.PrintShow()

	// 虽然 teacher  和 student 有同名的方法，但是，因为 teacher 的类型和student的类型不同，因此这两个方法不会被编译器认为是一样的方法
	// 也就是说，在调用方法时，类型很重要
	tch := Teacher52{201, "李老师"}
	tch.PrintShow()
}
