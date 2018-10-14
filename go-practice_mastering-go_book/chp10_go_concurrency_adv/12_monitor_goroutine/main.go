package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	readChan  = make(chan int)
	writeChan = make(chan int)
)

func set(newValue int) {
	writeChan <- newValue
}

func get() int {
	return <-readChan
}

func monitor() {
	var value int

	for {
		select {
		case v := <-writeChan:
			value = v
			fmt.Println("new val:", value)
		case readChan <- value:
		}
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("one integer argument required")
	}
	n, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("starting random numbers, n:", n)
	rand.Seed(time.Now().Unix())
	go monitor()

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(10 * n))
		}()
	}
	wg.Wait()
	fmt.Println("last value:", get())
}
