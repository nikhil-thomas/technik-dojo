package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

var times int

func chenGen(cc chan chan int, done chan bool) {
	c := make(chan int)
	cc <- c
	defer close(c)
	chanID := 0
	sum := 0
	for {
		select {
		case x := <-c:
			chanID = x
			for i := 0; i <= x; i++ {
				sum += i
			}
			fmt.Println("sum", x, "calculated")
			c <- sum
			fmt.Println("sum", x, "send")
		case <-done:
			fmt.Println("chan", chanID, "done")
			return
		}
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("specify number of channels")
	}

	times, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	cc := make(chan chan int)

	for i := 1; i <= times; i++ {
		done := make(chan bool)

		go chenGen(cc, done)

		ch := <-cc
		ch <- i

		sum := <-ch
		fmt.Print("Sum(", i, "): ", sum)

		fmt.Print("\n")
		close(done)
		time.Sleep(1 * time.Second)
	}
}
