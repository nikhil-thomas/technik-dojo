package main

import (
	"fmt"
	"time"
)

func a(c1, c2 chan struct{}) {
	<-c1
	fmt.Println("executing A")
	close(c2)
}

func b(c1, c2 chan struct{}) {
	<-c1
	fmt.Println("executing B")
	close(c2)
}

func c(c1 chan struct{}) {
	<-c1
	fmt.Println("executing C")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go c(z)
	go a(x, y)
	go a(x, y)
	go c(z)
	go b(y, z)
	go c(z)

	close(x)

	time.Sleep(3 * time.Second)
}

// Calling the C() function multiple times as goroutines will
// work just fine because C() does not close any channels.
// However, if you call A() or B() more than once,
// you will most likely get an error message
// panic: close of closed channel
