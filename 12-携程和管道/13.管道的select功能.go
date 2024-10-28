package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
	}()

	select {
	case value := <-ch1:
		fmt.Println("Received from ch1:", value)
	case value := <-ch2:
		fmt.Println("Received from ch2:", value)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	default:
		fmt.Println("No data received")
	}
}
