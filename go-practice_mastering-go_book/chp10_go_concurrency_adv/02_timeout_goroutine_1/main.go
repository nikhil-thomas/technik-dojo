package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "c1 ok"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case t := <-time.After(1 * time.Second):
		fmt.Println(t)
		fmt.Println("timeout c1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "c2 ok"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case t := <-time.After(3 * time.Second):
		fmt.Println(t)
		fmt.Println("time out c2")
	}
}
