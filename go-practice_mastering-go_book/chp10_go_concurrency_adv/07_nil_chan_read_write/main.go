package main

import (
	"fmt"
	"sync"
)

func main() {
	var c chan int
	// nil channel

	//c = make(chan int)

	fmt.Println(c)
	// output
	// nil

	// close(c)
	// will panic
	// panic: close of nil channel

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		c <- 123
		fmt.Println("write val:", 123)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		val := <-c
		fmt.Println("read val:", val)
	}()

	wg.Wait()
	fmt.Println("end of main")

	// output
	// fatal error: all goroutines are asleep - deadlock!

	// nil channel always blocks
}
