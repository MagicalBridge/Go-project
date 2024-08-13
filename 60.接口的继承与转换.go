package main

import "fmt"

// Animal 定义一个基本的接口Animal
type Animal interface {
	Speak() string
}

// Mammal 定义一个扩展的接口，可以看作是继承了Animal接口
type Mammal interface {
	Animal        // 继承了Animal接口，Mammal接口也必须实现Speak方法
	Walk() string // Mammal接口增加了一个新方法
}

// Dog 定义一个具体的结构体Dog，实现Mammal接口
type Dog struct{}

// Speak Dog结构体实现了Animal接口的Speak方法
func (d Dog) Speak() string {
	return "Woof!"
}

// Walk Dog结构体实现了Mammal接口的Walk方法
func (d Dog) Walk() string {
	return "Dog is walking."
}

func main() {
	// 使用Dog的实例
	var d Dog

	// 由于Dog实现了Mammal接口的所有方法，可以将其赋值给Mammal类型的变量
	var m Mammal = d

	fmt.Println(m.Speak()) // 输出: Woof!
	fmt.Println(m.Walk())  // 输出: Dog is walking.
}
