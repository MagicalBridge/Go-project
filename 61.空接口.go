package main

import "fmt"

func main() {

	var i interface{}
	fmt.Printf("%T\n", i)

	i = 10
	fmt.Printf("%T\n", i)
	fmt.Println(i)

	i = 3.14
	fmt.Printf("%T\n", i)
	fmt.Println(i)

	i = "传智播客"
	fmt.Printf("%T\n", i)
	fmt.Println(i)
}
