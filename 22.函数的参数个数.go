package main

import "fmt"

// 求1-100之间所有数字之和

func SumAdd() {
	var sum int

	for i:=1; i<= 100 ; i++ {
		sum += i
	}

	fmt.Println(sum)
}


func main()  {
	SumAdd()

}
