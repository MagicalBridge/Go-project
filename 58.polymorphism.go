package main

import "fmt"

// 先定义接口  在根据接口实现功能
type Humaner58 interface {
	//方法  方法声明
	SayHello()
	//Result(int, int) int
}

type Student58 struct {
	name  string
	age   int
	sex   string
	score int
}

type Teacher58 struct {
	name    string
	age     int
	sex     string
	subject string
}

func (s *Student58) SayHello() {
	fmt.Printf("大家好，我是%s，我今年%d岁，我是%s生，我的成绩是%d分\n",
		s.name, s.age, s.sex, s.score)
}

func (t *Teacher58) SayHello() {
	fmt.Printf("大家好，我是%s，我今年%d岁，我是%s生，我的学科是%s\n",
		t.name, t.age, t.sex, t.subject)
}

// 多态的实现 将接口作为函数参数  实现多态

func SayHi(h Humaner58) {
	h.SayHello()
}

func main() {
	stu := Student58{"小明", 18, "男", 99}
	//调用多态函数
	SayHi(&stu)

	tea := Teacher58{"法师", 31, "男", "go"}
	SayHi(&tea)
}
