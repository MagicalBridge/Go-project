package main

import (
	"fmt"
	"sync"
)

func main() {
	numChan := make(chan bool)
	letterChan := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 28; i += 2 {
			<-letterChan
			fmt.Printf("%d%d", i, i+1)
			numChan <- true
		}
	}()

	go func() {
		defer wg.Done()
		letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		letterIndex := 0

		for range numChan {
			if letterIndex >= len(letters) {
				break
			}
			fmt.Printf("%c%c", letters[letterIndex], letters[letterIndex+1])
			letterIndex += 2
			letterChan <- true
		}
	}()

	// Start the process
	letterChan <- true

	wg.Wait()
}
