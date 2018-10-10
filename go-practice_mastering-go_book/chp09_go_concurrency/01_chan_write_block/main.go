package main

import (
	"fmt"
	"time"
)

func chanWrite(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	// in this example this line will  not be print
	// as the above channel write blocks until there is a read
	fmt.Println("2", x)
}

func main() {
	c := make(chan int)
	go chanWrite(c, 10)
	time.Sleep(1 * time.Second)
}
