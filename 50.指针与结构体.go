package main

import "fmt"

func main() {
	type Student struct {
		id   int
		name string
	}
	stu := Student{101, "BTC"}
	var p *Student
	p = &stu

	fmt.Println(p) // &{101 BTC}
	fmt.Println((*p).name)
}
