package main

import "fmt"

func main() {
	s := []int{1, 2, 23, 4}
	var p *[]int
	p = &s
	fmt.Println(*p) // [1 2 23 4]
}
