package main

import (
	"fmt"
	"time"
)

func chanWrite(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func main() {
	c := make(chan int)
	go chanWrite(c, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("read:", <-c)
	time.Sleep(1 * time.Second)
}
