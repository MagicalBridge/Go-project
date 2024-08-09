package main

import "fmt"

// Person 父类 has four fields
type Person struct {
	id   int
	name string
	age  int
	sex  string
}

// Student51 子类
type Student51 struct {
	// 将Person结构体作为student结构体的成员
	Person // 匿名字段实现继承关系
	class  int
	score  int
}

func main() {
	// 如果一个结构体包含其他结构体信息 需要依次分开进行初始化
	var stu = Student51{
		Person: Person{1001, "东方不败", 35, "不祥"},
		score:  100,
		class:  302,
	}

	fmt.Printf("姓名：%s\n", stu.name)
	fmt.Printf("性别：%s\n", stu.sex)
	fmt.Printf("年龄：%d\n", stu.age)
	fmt.Printf("编号：%d\n", stu.id) // 这是一种简化写法，无需隐式添加Person前缀名
	fmt.Printf("班级：%d\n", stu.class)
	fmt.Printf("成绩：%d\n", stu.score)

	// partial_initialization 使用匿名字段进行部分初始化
	var partialInitialization = Student51{
		score: 100,
		class: 404,
	}

	fmt.Printf("部分初始化，score：%d\n", partialInitialization.score)
	fmt.Printf("部分初始化，班级：%d\n", partialInitialization.class)

	// if update class field to 405 and run again
	// you will find partial initialization field can be updated
	partialInitialization.class = 405
	fmt.Printf("编号：%d\n", partialInitialization.class) // 405
}
