package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	tempChan := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		defer close(tempChan)
		w.Wait()
	}()

	select {
	case <-tempChan:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("time duration not specified")
	}

	t, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("timeout period: %s\n", duration)

	var wg sync.WaitGroup
	wg.Add(1)
	if timeout(&wg, duration) {
		fmt.Println("timed out")
	} else {
		fmt.Println("ok")
	}
	wg.Done()

	if timeout(&wg, duration) {
		fmt.Println("timed out")
	} else {
		fmt.Println("ok")
	}
}
