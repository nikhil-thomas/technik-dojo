package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

func contextCancel(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("contextCancel():", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("contextCancel():", r)
	}
	return
}

func contextTimeout(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(4)*time.Second)
	defer cancel()

	select {
	case <-c2.Done():
		fmt.Println("contextTimeout():", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("contextTimeout():", r)
	}
	return
}

func contextDeadline(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(4) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	select {
	case <-c3.Done():
		fmt.Println("contextTimeout():", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("contextTimeout():", r)
	}
	return
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("specify delay")
	}
	delay, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Delay:", delay)
	contextCancel(delay)
	contextTimeout(delay)
	contextDeadline(delay)
}
