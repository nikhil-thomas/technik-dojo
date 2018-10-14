package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(c chan int) {
	sum := 0
	t := time.NewTimer(1 * time.Second)

	for {
		select {
		case input := <-c:
			fmt.Print(input, " ")
			sum += input
		case <-t.C:
			// a nil channel always blocks, which means that trying
			// to read or write from a nil channel will block.
			// This property of channels can be very useful when
			//  you want to disable a branch of a select statement
			// by assigning the nil value to a channel variable.
			c = nil
			fmt.Printf("\n%d\n", sum)
		}
	}
}

func genRand(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	c := make(chan int)

	go add(c)
	go genRand(c)
	time.Sleep(3 * time.Second)
	fmt.Println("end of main")
}
